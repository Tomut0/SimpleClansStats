<?php

namespace App\Enums;

enum SortTypes
{
    case Kdr;
    case Balance;
    case Members;

    /**
     * @throws Exception if sort type is undefined
     */
    public function value(): array|string
    {
        return match ($this) {
            SortTypes::Kdr => [
                "kdr" => [
                    "icon" => "FireIcon",
                    "translation" => "components.sortselector.kdr"
                ],
            ],
            SortTypes::Balance => [
                "balance" => [
                    "icon" => "CircleStackIcon",
                    "translation" => "components.sortselector.balance"
                ],
            ],
            SortTypes::Members => [
                "members" => [
                    "icon" => "UserGroupIcon",
                    "translation" => "components.sortselector.members"
                ],
            ],
            default => throw new Exception('Unknown sort type'),
        };
    }

    public static function packed(): array
    {
        try {
            return array_merge(SortTypes::Kdr->value(), SortTypes::Balance->value(), SortTypes::Members->value());
        } catch (Exception $e) {
            error_log($e->getMessage());
        }

        return [];
    }
}
