<?php

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

use Illuminate\Support\Env;
use Illuminate\Support\Facades\Route;

//Route::get('/detail/clan/{tag}', 'ViewController@detail')->name('clan');
//Route::get('/detail/player/{name}', 'ViewController@detail')->name('player');

$LOCALE_REGEX = '^[A-Za-z]{2,4}([_-][A-Za-z]{4})?([_-]([A-Za-z]{2}|[0-9]{3}))?$';

Route::get('/', 'ViewController@show');
Route::get('/clans', 'ViewController@show')->name('clans');

if (Env::get('LANGUAGE_SELECTOR')) {
    Route::group([
        'prefix' => '{locale}',
        'where' => ['locale' => $LOCALE_REGEX]
    ], function () {
        Route::get('/', 'ViewController@show');
        Route::get('/clans', 'ViewController@show')->name('clans');
    });
}
