<?php

/** @var Factory $factory */

use App\Kill;
use Faker\Generator as Faker;
use Illuminate\Database\Eloquent\Factory;

$factory->define(Kill::class, function (Faker $faker) {

    $randomClan = Clan::all()->random();
    $randomPlayers = $faker->randomElements($randomClan->players()->getResults()->toArray(), 2);

    return [
        'kill_type' => $faker->randomElement(['n', 'c', 'r', 'a']),
        'attacker' => $randomPlayers[0]["name"],
        'attacker_tag' => $randomPlayers[0]["tag"],
        'attacker_uuid' => $randomPlayers[0]["uuid"],
        'victim' => $randomPlayers[1]["name"],
        'victim_tag' => $randomPlayers[0]["tag"],
        'victim_uuid' => $randomPlayers[1]["uuid"],
    ];
});
