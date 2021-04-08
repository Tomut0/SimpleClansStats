<?php

namespace App\Http\Controllers;

use App\Clans;
use Illuminate\Http\Request;
use Illuminate\Support\Collection;

class ApiController extends Controller
{
    private $clans;

    public function __construct()
    {
        $this->clans = new Clans();
    }

    public function getClans(Request $request): Collection
    {
        $length = $request->input("length", 10);
        $sortBy = $request->input("sortBy", "KDR");
        $order = $request->input("order", "asc") == "asc";

        return $this->clans->getTopClans($length, $sortBy, $order);
    }

    public function getClan(string $tag): array
    {
        return (array) $this->clans->getClan($tag);
    }
}
