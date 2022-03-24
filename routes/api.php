<?php

use Illuminate\Http\Resources\Json\JsonResource;
use Illuminate\Http\Resources\Json\ResourceCollection;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::get('/player/{id}', function (Player $player) {
    return new JsonResource($player);
});

Route::get('/clan/{id}', function (Clan $clan) {
    return new JsonResource($clan);
});

Route::get('/clans', function () {
    return new ResourceCollection(Clan::all());
});
