<?php

namespace Database\Seeders;

use App\Models\Clan;
use App\Models\ClanPlayer;
use Illuminate\Database\Seeder;

class SCSSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Clan::factory(rand(5, 10))->create()->each(function (Clan $clan) {
            ClanPlayer::factory()->count(rand(1, 5))->create([
                'tag' => $clan->tag,
            ]);
        });
    }
}
