<?php

use App\Clan;
use App\Player;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        $clans = factory(Clan::class, 10)->create();
        $clans->each(function ($clan) {
            $clan->players()->createMany(factory(Player::class, 10)->create(['clan_id' => $clan->id])->toArray());
        });
    }
}
