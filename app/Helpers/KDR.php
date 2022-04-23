<?php

namespace App\Helpers;

use App\Clan;
use App\Player;
use Illuminate\Support\Env;

class KDR
{
    private int $neutrals;
    private int $rivals;
    private int $civilians;
    private int $allies;
    private int $deaths;

    /**
     * Creates a new KDR (Kill death ratio) instance from basic parameters.
     *
     * @param $neutrals int The count of killed neutral players
     * @param $rivals int The count of killed rivals
     * @param $civilians int The count of killed civilians
     * @param $allies int The count of killed allies
     * @param $deaths int The count of deaths
     */
    public function __construct(int $neutrals, int $rivals, int $civilians, int $allies, int $deaths)
    {
        $this->neutrals = $neutrals;
        $this->rivals = $rivals;
        $this->civilians = $civilians;
        $this->allies = $allies;
        $this->deaths = $deaths;
    }

    /**
     * Instantiates KDR from @param Player $player
     *
     * @return KDR
     * @link Player
     */
    public static function ofPlayer(Player $player): KDR
    {
        $neutral = $player->neutral_kills * doubleval(config("app.kdr.neutral", 1.0));
        $rival = $player->rival_kills * doubleval(config("app.kdr.rival", 2.0));
        $civilian = $player->civilian_kills * doubleval(config("app.kdr.civilian", 0.0));
        $ally = $player->ally_kills * doubleval(config("app.kdr.ally", -1.0));

        return new KDR($neutral, $rival, $civilian, $ally, $player->deaths);
    }

    /**
     * Instantiates KDR from @param Clan $clan
     *
     * @return KDR
     * @link Clan
     */
    public static function ofClan(Clan $clan): KDR
    {
        $totalKDR = $clan->players()->getResults()->map(function ($player) {
            return $player->getKDR()->toArray();
        });

        $result = $totalKDR->pipe(function ($collection) {
            return collect([
                'rivals' => $collection->sum('rivals'),
                'allies' => $collection->sum('allies'),
                'civilians' => $collection->sum('civilians'),
                'neutrals' => $collection->sum('neutrals'),
                'deaths' => $collection->sum('deaths'),
            ]);
        });

        return self::fromArray($result->toArray());
    }

    /**
     * Instantiates KDR from array
     *
     * @param array $array
     * @return KDR
     */
    public static function fromArray(array $array): KDR
    {
        return new self($array['neutrals'], $array['rivals'], $array['civilians'], $array['allies'], $array['deaths']);
    }

    /**
     * Retrieves KDR as string value
     *
     * @return string
     */
    public function asString(): string
    {
        return number_format($this->asFloat(), 2);
    }

    /**
     * Retrieves KDR as float value
     *
     * @return float
     */
    public function asFloat(): float
    {
        $kills = ($this->civilians + $this->rivals + $this->neutrals + $this->allies);
        $deaths = $this->deaths == 0 ? 1 : $this->deaths;

        return $kills / $deaths;
    }

    /**
     * Transforms KDR to array representation
     *
     * @return array
     */
    public function toArray(): array
    {
        return [
            'neutrals' => $this->neutrals,
            'allies' => $this->allies,
            'civilians' => $this->civilians,
            'rivals' => $this->rivals,
            'deaths' => $this->deaths
        ];
    }
}
