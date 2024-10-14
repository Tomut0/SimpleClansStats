<?php

namespace App\Enums;

use App\Traits\EnumPackable;
use Exception;

enum SortTypes
{
    use EnumPackable;

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
                    "translation" => "general.clan.kdr"
                ],
            ],
            SortTypes::Balance => [
                "balance" => [
                    "icon" => "CircleStackIcon",
                    "translation" => "general.clan.balance"
                ],
            ],
            SortTypes::Members => [
                "members" => [
                    "icon" => "UserGroupIcon",
                    "translation" => "general.clan.members"
                ],
            ],
            default => throw new Exception('Unknown sort type'),
        };
    }
}
