<?php

namespace App\Http\Controllers;

use App\Enums\Period;
use App\Enums\SortTypes;
use App\Models\Clan;
use App\Models\Kill;
use Carbon\Carbon;
use Error;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;
use Inertia\Response;
use Inertia\ResponseFactory;

// todo:
// 1. TB support
// 2. Overall leaderboard (by all sort types at the same time)
// 3. Dashboard: Most active clans (by total time)
// 4. Clan Page: Clan activity (now, 5 min ago, 15 min ago, 1 hour ago, 1 day ago)
// 5. Clan Page: Clan rivalry analysis - use clans.packed_rivals
// 6. Clan Page: Clan's founded (clans.founded)
// 7. Player page: "Time in a clan: {current time - sc_players.join_date}"
// 8. Player page: Player's loyalty
// 9. Kill frequency - count kills per day/week/month to analyze activity patterns
// 10. Death frequency - count deaths per player to see riskiest players
class LeaderboardController extends Controller
{
    public function index(Request $request): Response|ResponseFactory
    {
        $validatedData = $request->validate([
            'sortBy' => ['bail', 'string'],
            'period' => ['string'],
            'limit' => ['integer', 'min:1', 'max:100'],
        ]);

        $selectors = [
            'sortSelector' => SortTypes::packed(),
            'intervalSelector' => Period::packed(),
        ];

        $sortKeys = array_keys(SortTypes::packed());
        $sortBy = $validatedData['sortBy'] ?? $sortKeys[0];

        try {
            $period = constant("App\Enums\Period::" . ucfirst($validatedData['period'] ?? ""));
        } catch (Error $ignored) {
            $period = Period::from(Period::cases()[0]->value);
        }

        $currentSelectors = [
            'sortSelector' => [
                $sortBy => $selectors['sortSelector'][$sortBy]
            ],
            'intervalSelector' => [
                strtolower($period->name) => $selectors['intervalSelector'][strtolower($period->name)]
            ],
        ];

        $limit = $request->query('limit', 10);

        // Sort data by specified sorting criteria
        $data = Clan::data();
        $sortedData = $this->sortTopAll($data, $sortKeys);

        // Initial cache over all periods
        foreach (Period::cases() as $p) {
            $ttl = now()->addDays($p->value);
            Cache::add($p->name, ["positions" => $sortedData, "ttl" => $ttl], $ttl);
        }

        // Get past positions from cache
        $pastPositions = Cache::get($period->name)["positions"];

        // Calculate new positions based on the specified sorting criteria
        $newPositions = $data->sortByDesc($sortBy)->values()->take($limit)->toArray();

        // Compare positions between past and new positions
        $comparedPositions = $this->comparePositions($sortBy, $pastPositions, $newPositions);

        // Last 10 kills
        $lastKills = Kill::with('victimClan', 'attackerClan')->get(['attacker', 'attacker_tag', 'victim', 'victim_tag'])
            ->where('attacker_tag', '!=', '')->where('victim_tag', '!=', '')
            ->reverse()->take($limit)->map(function (Kill $kill) {
                if (!$kill->attackerClan || !$kill->victimClan) return null;

                return [
                    'attacker' => [
                        'name' => $kill->attacker,
                        'tag' => $kill->attackerClan->color_tag,
                    ],
                    'victim' => [
                        'name' => $kill->victim,
                        'tag' => $kill->victimClan->color_tag,
                    ]
                ];
            })->filter(fn($kill) => !is_null($kill))->values()->toArray();

        // Retrieve kills by type for the specified period
        $killsByType = Kill::whereNotNull('created_at')->get()
            ->filter(function (Kill $kill) use ($period) {
                return Carbon::make($kill->created_at)->greaterThanOrEqualTo(now()->subDays($period->value));
            })->map(function (Kill $kill) {
                return $kill->displayType();
            });


        // Retrieve statistics for charts
        $statistics = Cache::get($period->name)["statistics"] ?? [];
        $statistics['kills'] = $killsByType->countBy()->toArray();

        return inertia('Dashboard', ['clans' => $comparedPositions, 'lastKills' => $lastKills,
            'selectors' => ['current' => $currentSelectors, 'all' => $selectors], 'statistics' => $statistics
        ]);
    }

    /**
     * Compare positions between past and new positions.
     *
     * <p>
     * If the position is the same, set the position_status to "equal".
     * If the position is lowered, set the position_status to "lowered".
     * If the position is raised, set the position_status to "raised".
     * </p>
     *
     * @param string $sortBy
     * @param array $pastPositions
     * @param array $newPositions
     * @return array
     */
    private function comparePositions(string $sortBy, array $pastPositions, array $newPositions): array
    {
        foreach ($pastPositions as $pastIndx => $pastPosition) {
            $newIndex = array_search($pastPosition[$sortBy]['tag'], array_column($newPositions, 'tag'));

            if ($newIndex == $pastIndx) {
                $newPositions[$newIndex]['position_status'] = "equal";
            } else if ($pastIndx > $newIndex) {
                $newPositions[$newIndex]['position_status'] = "raised";
            } else {
                $newPositions[$newIndex]['position_status'] = "lowered";
            }
        }

        return $newPositions;
    }

    private function sortTopAll(Collection $clans, array $sortBy, int $limit = 10): array
    {
        $sortedData = [];
        foreach ($sortBy as $sortType) {
            $this->sortTop($sortedData, $clans, $sortType, $limit);
        }

        return $sortedData;
    }

    private function sortTop(array &$accumulator, Collection $clans, string $sortBy, int $limit = 10): void
    {
        $sorted = $clans->sortByDesc($sortBy)->values()->take($limit)->all();

        foreach ($sorted as $position => $clan) {
            $accumulator[$position][$sortBy] = $clan->toArray();
        }
    }
}
