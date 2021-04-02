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

use Illuminate\Support\Facades\Route;

Route::get('/', 'ViewController@show');
Route::get('/{locale}', 'ViewController@show');
Route::get('/clans', 'ViewController@show')->name("clans");
Route::get('/{locale}/clans', 'ViewController@show')->name("clans");
// FIXME This matches URLs with /
