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

Route::get('/player/{id}', fn(Player $player) => new JsonResource($player));
Route::get('/clans/{id}', fn(Clan $clan) => new JsonResource($clan));
Route::get('/clans', fn() => new ResourceCollection(Clan::all()));


Route::get('/clan/{id}', function (Clan $clan) {
    return new JsonResource($clan);
});

Route::get('/clans', function () {
    return new ResourceCollection(Clan::all());
});
