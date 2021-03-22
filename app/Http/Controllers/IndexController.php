<?php

namespace App\Http\Controllers;

use App\Clans;
use App\Players;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\App;

class IndexController extends Controller
{
    public function view($locale = 'en')
    {;
        $clans = new Clans();
        $players = new Players();
        App::setLocale($locale);

        return view('index', ['clans' => $clans, 'players' => $players]);
    }
}
