<?php

namespace App;

use App\Helpers\KDR;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property string $tag the clan tag prop
 */
class Clan extends Model
{
    /**
     * Indicates if the model should be timestamped.
     *
     * @var bool
     */
    public $timestamps = false;

    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'sc_clans';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'verified', 'tag', 'color_tag', 'name', 'members', 'description', 'founded', 'last_used', 'packed_allies',
        'packed_rivals', 'packet_bb', 'ranks'
    ];

    /**
     * The model's default values for attributes.
     *
     * @var array
     */
    protected $attributes = [
        'friendly_fire' => 0,
        'balance' => 0,
        'fee_enabled' => 0,
        'fee_balance' => 0,
    ];

    /**
     * Get the color tag of clan
     *
     * @param $key
     * @return string the HTML representation of color tag
     */
    public function getColorTagAttribute($key): string
    {
        return Utils::addColors($key);
    }

    /**
     * Retrieves KDR of all players in clan
     * @return KDR
     */
    public function getKDR(): KDR
    {
        return KDR::ofClan($this);
    }

    /**
     * Retrieves all leaders from clan
     * @return HasMany
     */
    public function leaders(): HasMany
    {
        return $this->players()->where('leader', '=', 1);
    }

    /**
     * Retrieves all players from clan
     * @return HasMany
     */
    public function players(): HasMany
    {
        return $this->hasMany('App\Player', 'tag', 'tag');
    }

    /**
     * @param Clan $clan The other clan to equals
     * @return bool true if clans are the same
     */
    public function equals(Clan $clan): bool
    {
        return $clan->tag === $this->tag;
    }
}
