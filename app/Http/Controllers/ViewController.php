<?php

namespace App\Http\Controllers;

use App\Clans;
use App\Players;
use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Route;

class ViewController extends Controller
{
    private $clans;
    private $players;

    public function __construct()
    {
        $this->clans = new Clans();
        $this->players = new Players();
    }

    public function show($locale = 'en')
    {
        $view_name = Route::currentRouteName();
        if (!isset($view_name)) {
            $view_name = "index";
        }
        App::setLocale($locale);

        return view($view_name, ['clans' => $this->clans, 'players' => $this->players, 'locale' => $locale]);
    }

    //TODO locale
    public function detail(string $id)
    {
        $detail_name = Route::currentRouteName();
        $clan = null;
        $player = null;
        switch ($detail_name) {
            case "clan":
                $clan = $this->clans->getClan($id);
                break;
            case "player":
                $player = $this->players->getPlayer($id);
        }

        return view('detail.' . $detail_name, ['clan' => $clan, 'player' => $player]);
    }
}
