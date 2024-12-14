<?php

return [
    'db_prefix' => env('SC_DATABASE_PREFIX', 'sc_'),
    'killWeight' => [
        'rival' => env('SC_KILL_WEIGHT_RIVAL', 2.0),
        'neutral' => env('SC_KILL_WEIGHT_NEUTRAL', 1.0),
        'civilian' => env('SC_KILL_WEIGHT_CIVILIAN', 0),
        'ally' => env('SC_KILL_WEIGHT_ALLY', 0),
    ],
    'currency' => env('SC_CURRENCY', 'USD'),
    'currencySymbol' => env('SC_CURRENCY_SYMBOL', '$'),
];
