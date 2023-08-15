<?php

namespace Database\Factories;

use App\Models\Clan;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends Factory<Clan>
 */
class ClanFactory extends Factory
{


    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        $tag = fake()->unique()->regexify("[a-zA-Z]{3,7}");
        return [
            'verified' => fake()->boolean(),
            'tag' => $tag,
            'color_tag' => $tag,
            'name' => fake()->unique()->company(),
            'friendly_fire' => fake()->boolean(),
            'founded' => fake()->unixTime(),
            'last_used' => fake()->unixTime(),
            'balance' => fake()->numberBetween(0, 1000),
            'fee_enabled' => fake()->boolean(),
            'flags' => "{}",
        ];
    }
}
