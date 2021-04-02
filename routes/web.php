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

Route::get('/', 'IndexController@view');
Route::get('/{locale}', 'IndexController@view');
Route::get('/clans', 'ClansController@view');
Route::get('/{locale}/clans', 'ClansController@view');

