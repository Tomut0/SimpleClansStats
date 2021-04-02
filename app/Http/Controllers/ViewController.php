<?php

namespace App\Http\Controllers;

use App\Clans;
use App\Players;
use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Route;

class ViewController extends Controller
{
    public function show($locale = 'en')
    {
        $view_name = Route::currentRouteName();
        if (!isset($view_name)) {
            $view_name = "index";
        }
        $clans = new Clans();
        $players = new Players();
        App::setLocale($locale);

        return view($view_name, ['clans' => $clans, 'players' => $players, 'locale' => $locale]);
    }
}
