<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;

/**
 * @property int $rival_kills
 * @property int $neutral_kills
 * @property int $civilian_kills
 * @property int $deaths
 */
class ClanPlayer extends Model
{
    protected $connection = 'simpleclans';
    protected $table = 'sc_players';
    public $timestamps = 'false';

    public function scopeByClan(Builder $query, string $clan)
    {
        $query->where('tag', $clan);
    }

    public function kdr(): float|int
    {
        $rival = $this->rival_kills * doubleval(env('KILL_RIVAL', 2.0));
        $neutral = $this->neutral_kills * doubleval(env('KILL_NEUTRAL', 1.0));
        $civilian = $this->civilian_kills * doubleval(env('KILL_CIVILIAN', 0.0));

        $kills = ($civilian + $rival + $neutral);

        $deaths = $this->deaths == 0 ? 1 : $this->deaths;

        return $kills / $deaths;
    }
}
