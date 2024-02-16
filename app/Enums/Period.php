<?php

namespace App\Enums;

use Exception;

enum Period: int
{
    case Daily = 1;
    case Weekly = 7;
    case Monthly = 30;

    /**
     * @throws Exception if sort type is undefined
     */
    public function value(): array|string
    {
        return match ($this) {
            Period::Daily => [
                "daily" => [
                    "translation" => "components.intervalselector.daily"
                ],
            ],
            Period::Weekly => [
                "weekly" => [
                    "translation" => "components.intervalselector.weekly"
                ],
            ],
            Period::Monthly => [
                "monthly" => [
                    "translation" => "components.intervalselector.monthly"
                ],
            ],
            default => throw new Exception('Unknown sort type'),
        };
    }


    public static function packed(): array
    {
        try {
            return array_merge(Period::Daily->value(), Period::Weekly->value(), Period::Monthly->value());
        } catch (Exception $e) {
            error_log($e->getMessage());
        }

        return [];
    }
}
