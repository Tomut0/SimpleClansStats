<?php

namespace App\Http\Controllers;

use App\Models\Clan;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;

class LeaderboardController extends Controller
{
    public function index(Request $request)
    {
        $startTime = hrtime(true);
        $startMemory = memory_get_peak_usage();

        $sortBy = $request->input('sortBy', 'balance');
        $period = Period::from($request->input('period', 7));
        $limit = $request->input('limit', 10);

        // get past data
        // past data -> get positions
        // new data -> get positions
        // compare positions
        // if new data position > past data position - better
        // otherwise, it's worst
        $cachedPastPositions = Cache::remember($period->name, now()->addDays($period->value), function () use ($limit, $sortBy) {
            return Clan::data()->sortByDesc($sortBy)->values()->take($limit)->toJson();
        });

        $pastPositions = json_decode($cachedPastPositions, true);

        // Calculate new positions
        $newPositions = Clan::data()->sortByDesc($sortBy)->values()->take($limit)->toArray();

        // Compare positions
        foreach ($newPositions as $newIndex => $newPosition) {
            // Get a new position by tag if any
            $pastIndex = array_search($newPosition['tag'], array_column($pastPositions, "tag"));

            if ($pastIndex == $newIndex) {
                $newPositions[$newIndex]['position_status'] = "equal";
            } else if ($pastIndex > $newIndex || $pastIndex === false) {
                $newPositions[$newIndex]['position_status'] = "raised";
            } else {
                $newPositions[$newIndex]['position_status'] = "lowered";
            }
        }

        $endTime = hrtime(true);
        $endMemory = memory_get_peak_usage();

        $executionTime = ($endTime - $startTime) / 1e9; // Convert to seconds
        $memoryUsage = $endMemory - $startMemory;

        echo "Execution time: " . number_format($executionTime, 6) . " seconds<br>";
        echo "Memory usage: " . number_format($memoryUsage / 1024, 2) . " KB";

        dd($newPositions);
    }
}

enum Period: int
{
    case DAILY = 1;
    case WEEKLY = 7;
    case MONTHLY = 30;
}
