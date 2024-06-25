<?php

namespace App;

use App\Enums\Period;
use App\Models\Clan;
use Illuminate\Support\Facades\Cache;

class Utils
{
    /**
     * This function returns lowered array of enum values.
     *
     * @param string $enum
     * @param string $column_key
     * @return array
     */
    public static function lowerEnum(string $enum, string $column_key = 'name'): array
    {
        return array_map('strtolower', array_column($enum::cases(), $column_key));
    }

    public static function saveStatistics(Period $period): void
    {
        if (!Cache::has($period->name)) {
            return;
        }
        $old = Cache::get($period->name);

        // Collect data from the database
        $old['statistics']['byTime'][now()->toDayDateTimeString()] = [
            "clans" => Clan::count(),
            "balance" => Clan::avg('balance'),
        ];

        Cache::put($period->name, $old, $old['ttl']);
    }
}
