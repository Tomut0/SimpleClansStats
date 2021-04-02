<?php

namespace App;

use Illuminate\Support\Env;

class Utils
{

    public static function getKDR($player): float
    {
        $neutral = $player->neutral_kills * doubleval(Env::get('KILL_NEUTRAL'));
        $rival = $player->rival_kills * doubleval(Env::get('KILL_RIVAL'));
        $civilian = $player->civilian_kills * doubleval(Env::get('KILL_CIVILIAN'));

        $kills = ($civilian + $rival + $neutral);

        $deaths = $player->deaths == 0 ? 1 : $player->deaths;

        return $kills / $deaths;
    }

}
