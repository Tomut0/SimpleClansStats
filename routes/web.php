<?php

use App\Http\Controllers\LeaderboardController;
use App\Http\Controllers\TranslationController;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', [LeaderboardController::class, 'index'])->name('leaderboard.index');
Route::post('/locale', [TranslationController::class, "update"])->name("locale.update");
