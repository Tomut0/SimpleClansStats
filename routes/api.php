<?php

use App\Http\Controllers\LeaderboardController;
use App\Http\Controllers\SearchController;
use App\Http\Controllers\TranslationController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::get('/search/language/{query}', [SearchController::class, 'language'])->name('search.language');

Route::get('/locales/status', [TranslationController::class, 'localizationMap'])->name('locales.status');

Route::get('/clans', [LeaderboardController::class, 'index']);
