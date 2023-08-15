<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

/**
 * @property string $tag
 */
class Clan extends Model
{
    use HasFactory;

    protected $connection = 'simpleclans';
    protected $table = 'sc_clans';
    public $timestamps = false;

    public function members(): Collection
    {
        return ClanPlayer::byClan($this->tag)->get();
    }

    public function countMembers(): int
    {
        return sizeof($this->members());
    }

    public function kdr(): float|int
    {
        return $this->members()->map(function (ClanPlayer $player) {
            return $player->kdr();
        })->sum();
    }

    public static function data(): Collection
    {
        return Clan::all(['id', 'tag', 'balance'])->map(function (Clan $clan) {
            $clan->members = $clan->countMembers();
            $clan->kdr = $clan->kdr();

            return $clan;
        });
    }
}
