<?php

namespace App;

use App\Helpers\KDR;
use Illuminate\Database\Eloquent\Model;

/**
 * @property int neutral_kills
 * @property int rival_kills
 * @property int ally_kills
 * @property int civilian_kills
 * @property int deaths
 */
class Player extends Model
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
    protected $table = 'sc_players';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'tag', 'uuid', 'leader', 'name', 'last_seen', 'join_date', 'locale', 'neutral_kills', 'civilian_kills',
        'rival_kills', 'ally_kills', 'deaths'
    ];

    /**
     * The model's default values for attributes.
     *
     * @var array
     */
    protected $attributes = [
        'friendly_fire' => 0,
        'trusted' => 0,
    ];

    public function getKDR(): KDR
    {
        return KDR::of($this);
    }
}
