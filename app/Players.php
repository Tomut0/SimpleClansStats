<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\DB;

class Players extends Model
{
    public function getTable()
    {
        return DB::table('sc_players');
    }

    public function getPlayers(): Collection
    {
        return $this->getTable()->get();
    }

    public function getPlayer(string $name): ?array
    {
        $player = $this->getPlayers()->where('name', '=', $name)->first();
        if (!isset($player)) {
            return null;
        }
        $player = (array) $player;
        $last_deaths = DB::table("sc_kills")->where('victim', '=', $name)
            ->select('attacker', 'attacker_tag', 'kill_id')->get();
        if (isset($last_deaths)) {
            $last_deaths =  $last_deaths->sortByDesc('kill_id')->splice(0, 5)->all();
        } else {
            $last_deaths = [];
        }
        $last_kills = DB::table("sc_kills")->where('attacker', '=', $name)
            ->select('victim', 'victim_tag', 'kill_id')->get();
        if (isset($last_kills)) {
            $last_kills = $last_kills->sortByDesc('kill_id')->splice(0, 5)->all();
        } else {
            $last_kills = [];
        }
        $player['last_kills'] = $last_kills;
        $player['last_deaths'] = $last_deaths;
        return (array)$player;
    }

    public function getMembers($clan_tag): Collection
    {
        return $this->getPlayers()->where('tag', '=', $clan_tag);
    }

    /** @noinspection PhpUnused */
    public function getTopPlayersByKDR(): array
    {
        $players = $this->getPlayers();
        if (!isset($players)) {
            return [];
        }
        $topPlayers = $players->map(function ($player) {
            $kdr = Utils::getKDR($player);

            if ($kdr > 0) {
                return ["Name" => $player->name, "KDR" => $kdr];
            }

            return null;
        });

        return $topPlayers->filter()->sortBy('KDR')->reverse()->splice(0, 10)->all();
    }

    /** @noinspection PhpUnused */
    public function getLastPlayersKills(): Collection
    {
        $clans = new Clans();
        $lastPlayers = collect([]);

        DB::table('sc_kills')->select('attacker', 'attacker_tag', 'victim', 'victim_tag')->get()
            ->reverse()->splice(0, 10)->each(function ($player, $index) use ($lastPlayers, $clans) {
                $player->attacker_colored_tag = $clans->getHTMLColorTagByTag($player->attacker_tag);
                $player->victim_colored_tag = $clans->getHTMLColorTagByTag($player->victim_tag);
                $lastPlayers->put($index, $player);
            });
        return $lastPlayers;
    }
}
