<?php

namespace App\Models;

use Carbon\Carbon;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property string $tag
 */
class Clan extends Model
{
    use HasFactory;

    protected $connection = 'simpleclans';
    protected $table = 'sc_clans';
    public $timestamps = false;

    public function members(): HasMany
    {
        return $this->hasMany(ClanPlayer::class, 'tag', 'tag');
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
        return Clan::all(['id', 'tag', 'color_tag', 'name', 'balance', 'founded', 'verified'])->map(function (Clan $clan) {
            $clan->members = $clan->countMembers();
            $clan->kdr = $clan->kdr();
            $clan->formatted_founded = Carbon::parse($clan->founded / 1000)->format('Y-m-d');

            return $clan;
        });
    }
}
