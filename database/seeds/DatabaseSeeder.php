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
        factory(Clan::class, 10)->create()->each(function ($clan) {
            factory(Player::class, mt_rand(0, 5))->create(['tag' => $clan->tag]);
        });
    }
}
