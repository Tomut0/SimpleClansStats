<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * App\Models\Kill
 *
 * @property int|null $kill_id
 * @property string $attacker
 * @property string $attacker_tag
 * @property string $victim
 * @property string $victim_tag
 * @property string $kill_type
 * @property string|null $attacker_uuid
 * @property string|null $victim_uuid
 * @property-read \App\Models\Clan|null $attackerClan
 * @property-read \App\Models\Clan|null $victimClan
 * @method static \Illuminate\Database\Eloquent\Builder|Kill newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Kill newQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Kill query()
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereAttacker($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereAttackerTag($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereAttackerUuid($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereKillId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereKillType($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereVictim($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereVictimTag($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Kill whereVictimUuid($value)
 * @mixin \Eloquent
 */
class Kill extends Model
{
    use HasFactory;

    protected $connection = 'simpleclans';
    public $timestamps = false;

    function getTable(): string
    {
        return config('scstats.db_prefix', 'sc_') . 'kills';
    }

    public function scopeCivilian($query)
    {
        return $query->where('kill_type', 'c');
    }

    public function scopeNeutral($query)
    {
        return $query->where('kill_type', 'n');
    }

    public function scopeRival($query)
    {
        return $query->where('kill_type', 'r');
    }

    public function scopeAlly($query)
    {
        return $query->where('kill_type', 'a');
    }


    public function victimClan(): BelongsTo
    {
        return $this->belongsTo(Clan::class, 'victim_tag', 'tag');
    }

    public function attackerClan(): BelongsTo
    {
        return $this->belongsTo(Clan::class, 'attacker_tag', 'tag');
    }

    public function displayType(): string
    {
        return match ($this->kill_type) {
            'c' => __('general.kills.types.civilian'),
            'n' => __('general.kills.types.neutral'),
            'r' => __('general.kills.types.rival'),
            'a' =>  __('general.kills.types.ally'),
        };
    }
}
