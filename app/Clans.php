<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\DB;

class Clans extends Model
{

    private static $players;

    function __construct(array $attributes = [])
    {
        parent::__construct($attributes);
        self::$players = new Players();
    }

    public function getTable()
    {
        return DB::table('sc_clans');
    }

    public function getClans(): Collection
    {
        return $this->getTable()->get()->sortBy('name');
    }

    /** @noinspection PhpUnused */
    public function getTopClans(int $length = 10, string $sortBy = "KDR"): Collection
    {
        $clans = collect([]);
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
        $clans = $clans->sortByDesc($sortBy)->splice(0, $length);
        $rank = 1;
        foreach ($clans as $tag => $clan) {
            $clan_row = $this->getClans()->where('tag', '=', $tag)->first();
            $clan += (array) $clan_row;
            $clan['color_tag'] = !array_key_exists('color_tag', $clan) ? '' : $this->addColors($clan['color_tag']);
            $clan['founded'] = !array_key_exists('founded', $clan) ? 0 : $clan['founded'];
            $clan['last_used'] = !array_key_exists('last_used', $clan) ? 0 : $clan['last_used'];
            if (isset($clan_row)) {
                $clan['name'] = $clan_row->name;
            } else {
                $clan['name'] = "Error: no clan";
            }
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
            return $this->addColors($row->color_tag);
        }
        return null;
    }

    /**
     * @param $string - without colors.
     * @return string - with colors.
     * @author lpostiglione
     * @link https://github.com/lpostiglione/SimpleClansStats2/blob/master/includes/functions.inc.php
     */
    private function addColors($string): string
    {
        $colors = array(
            "g" => "<span class=\"c0-font\">",
            "1" => "<span class=\"c1-font\">",
            "2" => "<span class=\"c2-font\">",
            "3" => "<span class=\"c3-font\">",
            "4" => "<span class=\"c4-font\">",
            "5" => "<span class=\"c5-font\">",
            "6" => "<span class=\"c6-font\">",
            "7" => "<span class=\"c7-font\">",
            "8" => "<span class=\"c8-font\">",
            "9" => "<span class=\"c9-font\">",
            "a" => "<span class=\"ca-font\">",
            "b" => "<span class=\"cb-font\">",
            "c" => "<span class=\"cc-font\">",
            "d" => "<span class=\"cd-font\">",
            "e" => "<span class=\"ce-font\">",
            "f" => "<span class=\"cf-font\">",
            "l" => "<span class=\"cl-font\">",
            "m" => "<span class=\"cm-font\">",
            "n" => "<span class=\"cn-font\">",
            "o" => "<span class=\"co-font\">",
            "k" => "<span class=\"ck-font\">"
        );

        $string = str_replace("§0", "§g", $string);
        $motdarr = explode("§", $string);
        $spans = 0;
        $colored = '<span class="colored">';

        foreach ($motdarr as $row) {
            if (!isset($isset)) {
                $isset = "isset";
                $colored .= $row;
            } elseif (empty($row)) {
                continue;
            } elseif (strtolower(substr($row, 0, 1)) == "r") {
                while ($spans > 0) {
                    $colored .= "</span>";
                    $spans--;
                }
            } else {
                $color_code = strtolower(substr($row, 0, 1));
                if (preg_match("/^[0-9a-g]$/i", $color_code)) {
                    while ($spans > 0) {
                        $colored .= "</span>";
                        $spans--;
                    }
                }
                $colored .= $colors[$color_code];
                $colored .= substr($row, 1);
                $spans++;
            }
        }
        while ($spans > 0) {
            $colored .= "</span>";
            $spans--;
        }
        return $colored . "</span>";
    }
}
