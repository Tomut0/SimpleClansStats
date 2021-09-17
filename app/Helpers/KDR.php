<?php

namespace App\Helpers;

use App\Player;
use Illuminate\Support\Env;

class KDR
{
    private int $neutrals;
    private int $rivals;
    private int $civilians;
    private int $allies;
    private int $deaths;

    /**
     * @param $neutral
     * @param $rival
     * @param $civilian
     * @param $ally
     * @param $deaths
     */
    public function __construct($neutral, $rival, $civilian, $ally, $deaths)
    {
        $this->neutrals = $neutral;
        $this->rivals = $rival;
        $this->civilians = $civilian;
        $this->allies = $ally;
        $this->deaths = $deaths;
    }

    public static function of(Player $player): KDR
    {
        $neutral = $player->neutral_kills * doubleval(Env::get('KILL_NEUTRAL'));
        $rival = $player->rival_kills * doubleval(Env::get('KILL_RIVAL'));
        $civilian = $player->civilian_kills * doubleval(Env::get('KILL_CIVILIAN'));
        $ally = $player->ally_kills * doubleval(Env::get('KILL_ALLY'));

        $deaths = $player->deaths == 0 ? 1 : $player->deaths;

        return new KDR($neutral, $rival, $civilian, $ally, $deaths);
    }

    public function asString(): string
    {
        return number_format($this->asFloat(), 2);
    }

    public function asFloat(): float
    {
        $kills = ($this->civilians + $this->rivals + $this->neutrals + $this->allies);

        return $kills / $this->deaths;
    }
}
