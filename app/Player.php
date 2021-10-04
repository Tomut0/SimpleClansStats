<?php

namespace App;

use App\Helpers\KDR;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

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
        return KDR::ofPlayer($this);
    }

    /**
     * @return int the count of all player' kills
     */
    public function getAllKills(): int
    {
        return ($this->ally_kills + $this->civilian_kills + $this->rival_kills + $this->neutral_kills);
    }

    /**
     * Retrieves the clan of player
     * @return BelongsTo
     */
    public function clan(): BelongsTo
    {
        return $this->belongsTo('App\Clan', 'tag', 'tag');
    }

    /**
     * Retrieves all kills of player
     * @return HasMany
     */
    public function kills(): HasMany
    {
        return $this->hasMany('App\Kill', 'attacker', 'kill_id');
    }
}
