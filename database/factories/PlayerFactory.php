<?php

/** @var Factory $factory */

use App\Player;
use Faker\Generator as Faker;
use Illuminate\Database\Eloquent\Factory;

$factory->define(Player::class, function (Faker $faker) {
    return [
        'uuid' => $faker->uuid(),
        'name' => $faker->userName(),
        'last_seen' => $faker->unixTime(),
        'join_date' => $faker->unixTime(),
        'locale' => $faker->locale(),
        'neutral_kills' => mt_rand(0, 100),
        'rival_kills' => mt_rand(0, 100),
        'ally_kills' => mt_rand(0, 100),
        'deaths' => mt_rand(0, 100),
    ];
});
