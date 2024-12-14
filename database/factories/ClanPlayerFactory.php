<?php

namespace Database\Factories;

use App\Models\Clan;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\ClanPlayer>
 */
class ClanPlayerFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'name' => $this->faker->userName,
            'leader' => $this->faker->boolean,
            'tag' => fn() => Clan::factory()->create()->tag,
            'friendly_fire' => $this->faker->boolean,
            'neutral_kills' => $this->faker->numberBetween(0, 100),
            'rival_kills' => $this->faker->numberBetween(0, 100),
            'civilian_kills' => $this->faker->numberBetween(0, 100),
            'deaths' => $this->faker->numberBetween(0, 100),
            'last_seen' => $this->faker->unixTime,
            'join_date' => $this->faker->unixTime,
            'trusted' => $this->faker->boolean,
            'flags' => json_encode([
                'bb-enabled' => true,
                'channel-state' => [true, true, true],
                'cape-enabled' => true,
                'channel' => 'NONE',
                'chat-shortcut' => false,
                'rank' => '',
                'hide-tag' => true,
            ]),
            'packed_past_clans' => '',
            'resign_times' => '{}',
            'uuid' => $this->faker->uuid,
            'locale' => fake()->locale, // Set empty or customize
            'ally_kills' => $this->faker->numberBetween(0, 100),
        ];
    }
}
