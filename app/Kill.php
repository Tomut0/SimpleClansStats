<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Kill extends Model
{
    /**
     * Indicates if the model should be timestamped.
     *
     * @var bool
     */
    public $timestamps = false;

    /**
     * The primary key for the model.
     *
     * @var string
     */
    protected $primaryKey = 'kills_id';


    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'sc_kills';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'attacker', 'attacker_tag', 'victim', 'victim_tag', 'kill_type', 'attacker_uuid', 'victim_uuid'
    ];

    /*
     * TODO: Define a specification for relationships of models
     * I don't really like that some models have different function names
     * (ex. $player->clan and $kills->mAttacker)
     */

    /**
     * Retrieves the attacker model
     *
     * @return BelongsTo
     */
    public function mAttacker(): BelongsTo
    {
        return $this->belongsTo('App\Player', 'attacker', 'name');
    }

    /**
     * Retrieves the victim model
     *
     * @return BelongsTo
     */
    public function mVictim(): BelongsTo
    {
        return $this->belongsTo('App\Player', 'victim', 'name');
    }
}
