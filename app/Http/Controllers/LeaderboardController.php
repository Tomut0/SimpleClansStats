<?php

namespace App\Http\Controllers;

use App\Enums\Period;
use App\Enums\SortTypes;
use App\Models\Clan;
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
        $sortBy = $request->input('sortBy', 'balance');
        $period = Period::from($request->input('period', 7));
        $limit = $request->input('limit', 10);
        $data = Clan::data();

        $sortTypes = array_column(SortTypes::cases(), 'name');

        // Sort and cache the data for the specified period
        $sortedData = $this->sortTopAll($data, $sortTypes);
        $cachedPastPositions = Cache::remember($period->name, now()->addDays($period->value), function () use ($sortedData, $limit, $sortBy) {
            return collect($sortedData)->toJson();
        });

        // Retrieve past positions from cache
        $pastPositions = json_decode($cachedPastPositions, true);

        // Calculate new positions based on the specified sorting criteria
        $newPositions = $data->sortByDesc($sortBy)->values()->take($limit)->toArray();

        // Compare positions between past and new positions
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

        return inertia('Dashboard', ['clans' => $newPositions, 'sortTypes' => $sortTypes]);
    }

    function sortTopAll(Collection $clans, array $sortBy, int $limit = 10): array
    {
        $sortedData = [];
        foreach ($sortBy as $sortType) {
            $this->sortTop($sortedData, $clans, $sortType, $limit);
        }

        return $sortedData;
    }

    function sortTop(array &$accumulator, Collection $clans, string $sortBy, int $limit = 10): void
    {
        $sorted = $clans->sortByDesc($sortBy)->values()->take($limit)->all();

        foreach ($sorted as $position => $clan) {
            $accumulator[$position][$sortBy] = $clan;
        }
    }
}
