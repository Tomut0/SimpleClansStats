<?php

namespace App;

use Exception;

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
}
