<?php

namespace App;

use Illuminate\Support\Env;

class Utils
{

    public static function getAllKills($player): int
    {
        $player = (array)$player;
        return $player['neutral_kills'] + $player['rival_kills'] + $player['civilian_kills'];
    }

    public static function getKDR($player): float
    {
        $player = (array)$player;
        $neutral = $player['neutral_kills'] * doubleval(Env::get('KILL_NEUTRAL'));
        $rival = $player['rival_kills'] * doubleval(Env::get('KILL_RIVAL'));
        $civilian = $player['civilian_kills'] * doubleval(Env::get('KILL_CIVILIAN'));

        $kills = ($civilian + $rival + $neutral);

        $deaths = $player['deaths'] == 0 ? 1 : $player['deaths'];

        return $kills / $deaths;
    }

    public static function getWarringFromFlags($flags): string
    {
        $flags = json_decode($flags, true);
        return implode(", ", $flags['warring']);
    }

    public static function getFormattedKDR($player): string
    {
        return number_format(self::getKDR($player), 2);
    }

    public static function unpackClans($clans): string
    {
        return implode(", ", explode("|", $clans));
    }

    public static function formatPastClans($packed_clans): string
    {
        $packed_clans = explode("|", $packed_clans);
        foreach ($packed_clans as &$clan) {
            $clan = self::addColors($clan);
        }
        return implode(", ", $packed_clans);
    }

    public static function formatDate($founded)
    {
        // TODO get format from env
        return date("d/m/Y", $founded / 1000);
    }

    public static function formatDateTime($last_seen)
    {
        // TODO get format from env
        return date("d/m/Y - H:i", $last_seen / 1000);
    }

    /**
     * @param $string - without colors.
     * @return string - with colors.
     * @author lpostiglione
     * @link https://github.com/lpostiglione/SimpleClansStats2/blob/master/includes/functions.inc.php
     */
    public static function addColors($string): string
    {
        $colors = array(
            "g" => "<span class=\"c0-font\">",
            "1" => "<span class=\"c1-font\">",
            "2" => "<span class=\"c2-font\">",
            "3" => "<span class=\"c3-font\">",
            "4" => "<span class=\"c4-font\">",
            "5" => "<span class=\"c5-font\">",
            "6" => "<span class=\"c6-font\">",
            "7" => "<span class=\"c7-font\">",
            "8" => "<span class=\"c8-font\">",
            "9" => "<span class=\"c9-font\">",
            "a" => "<span class=\"ca-font\">",
            "b" => "<span class=\"cb-font\">",
            "c" => "<span class=\"cc-font\">",
            "d" => "<span class=\"cd-font\">",
            "e" => "<span class=\"ce-font\">",
            "f" => "<span class=\"cf-font\">",
            "l" => "<span class=\"cl-font\">",
            "m" => "<span class=\"cm-font\">",
            "n" => "<span class=\"cn-font\">",
            "o" => "<span class=\"co-font\">",
            "k" => "<span class=\"ck-font\">"
        );

        $string = str_replace("§0", "§g", $string);
        $motdarr = explode("§", $string);
        $spans = 0;
        $colored = '<span class="colored">';

        foreach ($motdarr as $row) {
            if (!isset($isset)) {
                $isset = "isset";
                $colored .= $row;
            } elseif (empty($row)) {
                continue;
            } elseif (strtolower(substr($row, 0, 1)) == "r") {
                while ($spans > 0) {
                    $colored .= "</span>";
                    $spans--;
                }
            } else {
                $color_code = strtolower(substr($row, 0, 1));
                if (preg_match("/^[0-9a-g]$/i", $color_code)) {
                    while ($spans > 0) {
                        $colored .= "</span>";
                        $spans--;
                    }
                }
                $colored .= $colors[$color_code];
                $colored .= substr($row, 1);
                $spans++;
            }
        }
        while ($spans > 0) {
            $colored .= "</span>";
            $spans--;
        }
        return $colored . "</span>";
    }

}
