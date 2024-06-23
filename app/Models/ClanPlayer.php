<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * App\Models\ClanPlayer
 *
 * @property int $rival_kills
 * @property int $neutral_kills
 * @property int $civilian_kills
 * @property int $deaths
 * @property int|null $id
 * @property string $name
 * @property int|null $leader
 * @property string $tag
 * @property int|null $friendly_fire
 * @property int $last_seen
 * @property int $join_date
 * @property int|null $trusted
 * @property string $flags
 * @property string|null $packed_past_clans
 * @property string|null $resign_times
 * @property string|null $uuid
 * @property-read \App\Models\Clan|null $clan
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer newQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer query()
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereCivilianKills($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereDeaths($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereFlags($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereFriendlyFire($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereJoinDate($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereLastSeen($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereLeader($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereName($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereNeutralKills($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer wherePackedPastClans($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereResignTimes($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereRivalKills($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereTag($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereTrusted($value)
 * @method static \Illuminate\Database\Eloquent\Builder|ClanPlayer whereUuid($value)
 * @mixin \Eloquent
 */
class ClanPlayer extends Model
{
    protected $connection = 'simpleclans';
    protected $table = 'sc_players';
    public $timestamps = 'false';

    public function clan(): BelongsTo
    {
        return $this->belongsTo(Clan::class, 'tag', 'tag');
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
