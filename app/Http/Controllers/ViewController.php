<?php

namespace App\Http\Controllers;

use Illuminate\Contracts\Foundation\Application;
use Illuminate\Contracts\View\Factory;
use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Route;
use Illuminate\View\View;

class ViewController extends Controller
{
    public function show($locale = 'en')
    {
        $view_name = Route::currentRouteName();
        if (!isset($view_name)) {
            $view_name = "index";
        }

        App::setLocale($locale);

        return view($view_name, ['locale' => $locale]);
    }

    /*
     * TODO locale
     */

    /**
     * @param string $id
     * @return Application|Factory|View
     */
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
