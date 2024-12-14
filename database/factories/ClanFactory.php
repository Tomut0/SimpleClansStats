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

        $banner = fake()->randomElement([
            "cs:
  ==: org.bukkit.inventory.ItemStack
  v: 3700
  type: LIGHT_BLUE_BANNER
  meta:
    ==: ItemMeta
    meta-type: BANNER
    ItemFlags:
    - HIDE_POTION_EFFECTS
    patterns:
    - ==: Pattern
      color: GREEN
      pattern: cr
    - ==: Pattern
      color: LIGHT_GRAY
      pattern: cre
    - ==: Pattern
      color: RED
      pattern: dls
", "cs:
  ==: org.bukkit.inventory.ItemStack
  v: 3700
  type: LIGHT_BLUE_BANNER
  meta:
    ==: ItemMeta
    meta-type: BANNER
    ItemFlags:
    - HIDE_POTION_EFFECTS
", "cs:
  ==: org.bukkit.inventory.ItemStack
  v: 3700
  type: CYAN_BANNER
  meta:
    ==: ItemMeta
    meta-type: BANNER
    ItemFlags:
    - HIDE_POTION_EFFECTS
    patterns:
    - ==: Pattern
      color: BLUE
      pattern: cre
    - ==: Pattern
      color: BLACK
      pattern: sku
", ""
        ]);

        $colors = ['§4', '§2', '§3', '§f', '§9', '§6', '§c', '§8'];
        $color = fake()->randomElement($colors);
        $color_tag = "$color$tag";

        $lastUsed = fake()->unixTime() * 1000;
        $founded = fake()->unixTime($lastUsed);

        $packed_bb = "${founded}_Clan ${tag} created";

        return [
            'verified' => fake()->boolean(),
            'tag' => $tag,
            'color_tag' => $color_tag,
            'name' => join(' ', fake()->unique()->words(rand(1, 2))),
            'description' => fake()->realText(),
            'friendly_fire' => fake()->boolean(),
            'founded' => $founded,
            'last_used' => $lastUsed,
            'balance' => fake()->numberBetween(0, 1000),
            'fee_enabled' => fake()->boolean(),
            'flags' => "{\"warring\":[]}",
            'packed_bb' => $packed_bb,
            'packed_allies' => '',
            'packed_rivals' => '',
            'cape_url' => '',
            'ranks' => '',
            'banner' => $banner,
        ];
    }
}
