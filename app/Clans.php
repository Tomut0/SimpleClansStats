<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\DB;

class Clans extends Model
{

    private static Players $players;

    function __construct(array $attributes = [])
    {
        parent::__construct($attributes);
        self::$players = new Players();
    }

    public function getTable()
    {
        return DB::table('sc_clans');
    }

    public function getClan(string $tag): ?array
    {
        $clan = $this->getTable()->where('tag', '=', $tag)->first();
        if (!isset($clan)) {
            return null;
        }
        $clan = ((array)$clan) + ["KDR" => 0.0, "leaders" => [], "members" => []];

        foreach (self::$players->getMembers($tag) as $member) {
            $clan["KDR"] += Utils::getKDR($member);
            array_push($clan["members"], $member->name);
            if ($member->leader) {
                array_push($clan["leaders"], $member->name);
            }
        }
        return $clan;
    }

    public function getClans(string $sortBy = "name"): Collection
    {
        return $this->getTable()->get()->sortBy($sortBy);
    }

    /** @noinspection PhpUnused */
    public function getTopClans(int $length = 10, string $sortBy = "KDR", bool $asc = false): Collection
    {
        $clans = collect([]);

        //looping through sc_players table to get KDR, members and leaders which are not available in sc_clans
        foreach (self::$players->getPlayers() as $player) {
            $tag = $player->tag;
            if (empty($tag)) {
                continue;
            }
            $clan = $clans->get($tag);
            $KDR = Utils::getKDR($player);
            $player_name = $player->name;
            if (!isset($clan)) {
                $clan = ["members" => [], "leaders" => [], "KDR" => 0.0];
            }
            array_push($clan["members"], $player_name);
            if ($player->leader) {
                array_push($clan["leaders"], $player_name);
            }
            $clan["KDR"] += $KDR;
            $clans->put($tag, $clan);
        }

        //looping through sc_clans to complete the clan data fetched in sc_players
        foreach ($this->getClans() as $clan_row) {
            $tag = $clan_row->tag;
            $clan = $clans->get($tag);
            if (!isset($clan)) {
                continue;
            }
            $clan = array_merge($clan, (array) $clan_row);
            $clan['color_tag'] = !array_key_exists('color_tag', $clan) ? '' : Utils::addColors($clan['color_tag']);
            $clan['founded'] = !array_key_exists('founded', $clan) ? 0 : $clan['founded'];
            $clan['last_used'] = !array_key_exists('last_used', $clan) ? 0 : $clan['last_used'];
            if (isset($clan_row)) {
                $clan['name'] = $clan_row->name;
            } else {
                $clan['name'] = "Error: no clan";
            }
            $clans->put($tag, $clan);
        }
        $clans = $clans->sortBy($sortBy, SORT_REGULAR, !$asc)->splice(0, $length);
        //setting ranks
        $rank = 1;
        foreach ($clans as $tag => $clan) {
            $clan['rank'] = $rank;
            $clans->put($tag, $clan);
            $rank++;
        }
        return $clans;
    }

    /** @noinspection PhpUnused */
    public function getTotalKDR(Collection $members)
    {
        return $members->sum(function ($player) {
            return Utils::getKDR($player);
        });
    }

    /** @noinspection PhpUnused */
    public function getColorTagByTag($tag = null)
    {
        return isset($tag) ? $this->getTable()->where('tag', '=', $tag)->get('color_tag')->first() : null;
    }

    public function getHTMLColorTagByTag($tag = null): ?string
    {
        if (!isset($tag)) {
            return null;
        }
        $row = $this->getTable()->where('tag', '=', $tag)->get('color_tag')->first();
        if (isset($row)) {
            return Utils::addColors($row->color_tag);
        }
        return null;
    }
}
