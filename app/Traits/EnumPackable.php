<?php

namespace App\Traits;

trait EnumPackable
{
    public static function packed(): array
    {
        return array_reduce(static::cases(), function ($carry, $case) {
            return array_merge($carry, $case->value());
        }, []);
    }
}
