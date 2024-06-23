<?php

namespace App\Console;

use App\Enums\Period;
use App\Utils;
use Illuminate\Console\Scheduling\Schedule;
use Illuminate\Foundation\Console\Kernel as ConsoleKernel;

class Kernel extends ConsoleKernel
{
    /**
     * Define the application's command schedule.
     */
    protected function schedule(Schedule $schedule): void
    {
        $schedule->command('scstats:collect-monthly')->weekly();
        $schedule->command('scstats:collect-weekly')->daily();
        $schedule->command('scstats:collect-daily')->hourly();
    }

    /**
     * Register the commands for the application.
     */
    protected function commands(): void
    {
        $this->load(__DIR__.'/Commands');

        require base_path('routes/console.php');
    }
}
