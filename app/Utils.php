<?php

namespace App;

class Utils
{

    public static function getAllKills($player): int
    {
        $player = (array)$player;
        return $player['neutral_kills'] + $player['rival_kills'] + $player['civilian_kills'];
    }

    public static function getWarringFromFlags($flags): string
    {
        $flags = json_decode($flags, true);
        return implode(", ", $flags['warring']);
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

    /**
     * @param $string - without colors.
     * @return string - with colors.
     * @author lpostiglione
     * @link https://github.com/lpostiglione/SimpleClansStats2/blob/master/includes/functions.inc.php
     */
    public static function addColors($string): string
    {
        $colors = array(
            "g" => "<span class=\"color-black\">",
            "1" => "<span class=\"color-dark_blue\">",
            "2" => "<span class=\"color-dark_green\">",
            "3" => "<span class=\"color-dark_aqua\">",
            "4" => "<span class=\"color-dark_red\">",
            "5" => "<span class=\"color-dark_purple\">",
            "6" => "<span class=\"color-gold\">",
            "7" => "<span class=\"color-gray\">",
            "8" => "<span class=\"color-dark_gray\">",
            "9" => "<span class=\"color-blue\">",
            "a" => "<span class=\"color-green\">",
            "b" => "<span class=\"color-aqua\">",
            "c" => "<span class=\"color-red\">",
            "d" => "<span class=\"color-light_purple\">",
            "e" => "<span class=\"color-yellow\">",
            "f" => "<span class=\"color-white\">",
            "l" => "<span class=\"bold\">",
            "m" => "<span class=\"strikethrough\">",
            "n" => "<span class=\"underline\">",
            "o" => "<span class=\"italic\">",
            "k" => "<span class=\"obfuscated\">"
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

}
