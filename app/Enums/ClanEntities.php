<?php

namespace App\Enums;

use App\Models\Clan;
use Carbon\Carbon;
use Exception;

use App\Traits\EnumPackable;
use Illuminate\Support\Collection;

enum ClanEntities
{
    use EnumPackable;

    case Members;
    case Allies;
    case Rivals;

    /**
     * @throws Exception if sort type is undefined
     */
    public function value(): array|string
    {
        return match ($this) {
            ClanEntities::Members => [
                "members" => [
                    "translation" => "general.clan.members",
                ],
            ],
            ClanEntities::Allies => [
                "allies" => [
                    "translation" => "general.clan.allies"
                ],
            ],
            ClanEntities::Rivals => [
                "rivals" => [
                    "translation" => "general.clan.rivals"
                ],
            ],
            default => throw new Exception("Unknown clan entity type"),
        };
    }

    public function getEntities(Clan $clan): Collection
    {
        return match ($this) {
            ClanEntities::Members => $clan->members()->get()->sortByDesc('leader')->values()->map(function ($member) {
                return [
                    'image_src' => "https://mc-heads.net/avatar/$member->name/32",
                    'name' => $member->name,
                    'value' => $member->leader ? __('general.clanplayer.leader') : __('general.clanplayer.member'),
                ];
            }),
            ClanEntities::Allies => $clan->allies()->map(function ($ally) {
                return [
                    'name' => $ally->name,
                    'value' => Carbon::parse($ally->founded / 1000)->format('Y-m-d'),
                ];
            }),
            ClanEntities::Rivals => $clan->rivals()->map(function ($rival) {
                return [
                    'name' => $rival->name,
                    'value' => Carbon::parse($rival->founded / 1000)->format('Y-m-d')
                ];
            }),
        };
    }
}
