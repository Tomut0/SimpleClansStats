<?php

namespace App;

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
        'verified', 'tag', 'name', 'members', 'description', 'founded', 'last_used', 'packed_allies',
        'packed_rivals', 'packet_bb', 'ranks', 'KDR'
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
     * @param Clan $clan The other clan to equals
     * @return bool true if clans are the same
     */
    public function equals(Clan $clan): bool
    {
        return $clan->tag === $this->tag;
    }

    /**
     * @return HasMany All players from clan
     */
    public function players(): HasMany
    {
        return $this->hasMany('App\Player', 'tag', 'tag');
    }
}
