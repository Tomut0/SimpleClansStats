<?php

namespace App\Http\Controllers;

use App\Clans;
use App\Players;
use Illuminate\Support\Facades\App;

class ClansController extends Controller
{
    public function view($locale = 'en')
    {
        $clans = new Clans();
        $players = new Players();
        App::setLocale($locale);

        return view('clans', ['clans' => $clans, 'players' => $players, 'locale' => $locale]);
    }
}
