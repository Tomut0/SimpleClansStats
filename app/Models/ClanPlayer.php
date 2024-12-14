<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
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
    use HasFactory;

    protected $connection = 'simpleclans';

    public $timestamps = false;

    public function getTable(): string
    {
        return config('scstats.db_prefix', 'sc_') . 'players';
    }

    public function clan(): BelongsTo
    {
        return $this->belongsTo(Clan::class, 'tag', 'tag');
    }

    public function kdr(): float|int
    {
        $rival = $this->rival_kills * config('scstats.killWeight.rival');
        $neutral = $this->neutral_kills * config('scstats.killWeight.neutral');
        $civilian = $this->civilian_kills * config('scstats.killWeight.civilian');

        $kills = ($civilian + $rival + $neutral);

        // after SC 2.15.1
        if (isset($this->ally_kills)) {
            $ally = $this->ally_kills * config('scstats.killWeight.ally');
            $kills += $ally;
        }

        $deaths = $this->deaths ?: 1;

        return $kills / $deaths;
    }
}
