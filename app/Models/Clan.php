<?php

namespace App\Models;

use Carbon\Carbon;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;
use Illuminate\Support\Collection;

/**
 * App\Models\Clan
 *
 * @property string $tag
 * @property int|null $id
 * @property int|null $verified
 * @property string $color_tag
 * @property string $name
 * @property string|null $description
 * @property int|null $friendly_fire
 * @property int $founded
 * @property int $last_used
 * @property string $packed_allies
 * @property string $packed_rivals
 * @property string $packed_bb
 * @property string $cape_url
 * @property string $flags
 * @property float|null $balance
 * @property int|null $fee_enabled
 * @property float|null $fee_value
 * @property string $ranks
 * @property-read Collection<int, \App\Models\ClanPlayer> $members
 * @method static \Database\Factories\ClanFactory factory($count = null, $state = [])
 * @method static \Illuminate\Database\Eloquent\Builder|Clan newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Clan newQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|Clan query()
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereBalance($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereCapeUrl($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereColorTag($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereDescription($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereFeeEnabled($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereFeeValue($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereFlags($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereFounded($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereFriendlyFire($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereLastUsed($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereName($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan wherePackedAllies($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan wherePackedBb($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan wherePackedRivals($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereRanks($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereTag($value)
 * @method static \Illuminate\Database\Eloquent\Builder|Clan whereVerified($value)
 * @mixin \Eloquent
 */
class Clan extends Model
{
    use HasFactory;

    protected $connection = 'simpleclans';

    public $timestamps = false;

    protected $fillable = [
        'balance'
    ];

    public function members(): HasMany
    {
        return $this->hasMany(ClanPlayer::class, 'tag', 'tag');
    }

    public function allies(): Collection
    {
        return $this->whereIn('tag', explode('|', $this->packed_allies))->get();
    }

    public function rivals(): Collection
    {
        return $this->whereIn('tag', explode('|', $this->packed_rivals))->get();
    }

    public function countMembers(): int
    {
        return $this->members()->count();
    }

    public function kdr(): float|int
    {
        return $this->members()->get()->sum(function (ClanPlayer $player) {
            return $player->kdr();
        });
    }

    public static function data(): Collection
    {
        return Clan::all(['id', 'tag', 'color_tag', 'name', 'description', 'balance', 'founded', 'verified', 'packed_allies', 'packed_rivals', 'banner'])
            ->map(function (Clan $clan) {
                $clan->members_count = $clan->countMembers();
                $clan->kdr = $clan->kdr();
                $clan->formatted_founded = Carbon::parse($clan->founded / 1000)->format('Y-m-d');

                return $clan;
            });
    }

    function getTable(): string
    {
        return config('scstats.db_prefix', 'sc_') . 'clans';
    }
}
