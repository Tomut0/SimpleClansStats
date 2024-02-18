<?php

namespace App\Http\Controllers;

use App\Enums\Period;
use App\Enums\SortTypes;
use App\Models\Clan;
use Error;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;
use Inertia\Response;
use Inertia\ResponseFactory;

// todo:
// 1. TB support
// 2. Overall leaderboard (by all sort types at the same time)

class LeaderboardController extends Controller
{
    public function index(Request $request): Response|ResponseFactory
    {
        $validatedData = $request->validate([
            'sortBy' => ['bail', 'nullable', 'string'],
            'period' => ['nullable', 'string'],
            'limit' => ['nullable', 'integer', 'min:1', 'max:100'],
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

        $limit = $request->query('limit', 10);

        // Sort data by specified sorting criteria
        $data = Clan::data();
        $sortedData = $this->sortTopAll($data, $sortKeys);

        // Initial cache over all periods
        foreach (Period::cases() as $p) {
            Cache::add($p->name, $sortedData, now()->addDays($p->value));
        }

        // Get past positions from cache
        $pastPositions = Cache::get($period->name);

        // Calculate new positions based on the specified sorting criteria
        $newPositions = $data->sortByDesc($sortBy)->values()->take($limit)->toArray();

        // Compare positions between past and new positions
        $comparedPositions = $this->comparePositions($sortBy, $pastPositions, $newPositions);

        return inertia('Dashboard', ['clans' => $comparedPositions, 'selectors' => $selectors]);
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
