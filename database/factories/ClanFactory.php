<?php

/** @var Factory $factory */

use App\Clan;
use Faker\Generator as Faker;
use Illuminate\Database\Eloquent\Factory;

/*
 * TODO: Add all column to factory
 * - color_tag | doesn't have colors now
 * - packed_allies
 * - packed_rivals
 * - cape_url
 * - banner
 */
$factory->define(Clan::class, function (Faker $faker) {
    $tag = $faker->unique()->word();

    return [
        'verified' => mt_rand(0, 1),
        'tag' => $tag,
        'color_tag' => $tag,
        'name' => $faker->words(mt_rand(1, 3), true),
        'description' => $faker->optional()->text(255),
        'founded' => $faker->unixTime(),
        'last_used' => $faker->unixTime(),
        'packed_bb' => function () use ($faker) {
            $bbMessages = $faker->sentences(mt_rand(1, 3), false);
            foreach ($bbMessages as &$message) {
                $message = time() . '_' . $message;
            }

            return implode('|', $bbMessages);
        },
        'balance' => mt_rand(0, 160000),
        'fee_enabled' => mt_rand(0, 1),
        'fee_balance' => mt_rand(0, 1000000),
        'ranks' => $faker->words(3, true),
    ];
});
