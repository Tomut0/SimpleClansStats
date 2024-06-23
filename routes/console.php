<?php

use App\Enums\Period;
use App\Utils;
use Illuminate\Foundation\Inspiring;
use Illuminate\Support\Facades\Artisan;

/*
|--------------------------------------------------------------------------
| Console Routes
|--------------------------------------------------------------------------
|
| This file is where you may define all of your Closure based console
| commands. Each Closure is bound to a command instance allowing a
| simple approach to interacting with each command's IO methods.
|
*/

Artisan::command('inspire', function () {
    $this->comment(Inspiring::quote());
})->purpose('Display an inspiring quote');

Artisan::command('scstats:collect-daily', function () {
    Utils::saveStatistics(Period::Daily);
    $this->comment('Daily statistics collected.');
})->purpose('Makes a snapshot of data from the database and stores it in the cache.');

Artisan::command('scstats:collect-weekly', function () {
    Utils::saveStatistics(Period::Weekly);
    $this->comment('Weekly statistics collected.');
})->purpose('Makes a snapshot of data from the database and stores it in the cache.');

Artisan::command('scstats:collect-monthly', function () {
    Utils::saveStatistics(Period::Monthly);
    $this->comment('Monthly statistics collected.');
})->purpose('Makes a snapshot of data from the database and stores it in the cache.');
