<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Env;
use Illuminate\Support\Facades\DB;
use phpDocumentor\Reflection\Types\Collection;

class Players extends Model
{
    public function getTable()
    {
        return DB::table('sc_players');
    }

    public function getPlayers()
    {
        return $this->getTable()->get();
    }

    public function getTopPlayersByKDR()
    {
        $topPlayers = $this->getPlayers()->map(function ($player) {
            $neutral = $player->neutral_kills * doubleval(Env::get('KILL_NEUTRAL'));
            $rival = $player->rival_kills * doubleval(Env::get('KILL_RIVAL'));
            $civilian = $player->civilian_kills * doubleval(Env::get('KILL_CIVILIAN'));

            $kills = ($civilian + $rival + $neutral);

            if ($player->deaths != 0) {
                return ["Name" => $player->name, "KDR" => number_format($kills / $player->deaths, 2)];
            }

            if ($kills != 0.0) {
                return ["Name" => $player->name, "KDR" => number_format($kills, 2)];
            }

            return null;
        });

        return $topPlayers->filter()->sortBy('KDR')->reverse()->splice(0, 10)->all();
    }

    public function getLastPlayersKills()
    {
        $clans = new Clans();
        $lastPlayers = collect([]);

        DB::table('sc_kills')->select('attacker','attacker_tag', 'victim', 'victim_tag')->get()->reverse()->splice(0, 10)->each(function ($player, $indx) use ($lastPlayers, $clans) {
            $player->attacker_colored_tag = $clans->getHTMLColorTagByTag($player->attacker_tag);
            $player->victim_colored_tag = $clans->getHTMLColorTagByTag($player->victim_tag);
            $lastPlayers->put($indx, $player);
        });
        return $lastPlayers;
    }
}
