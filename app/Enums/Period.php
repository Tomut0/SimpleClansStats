<?php

namespace App\Enums;

enum Period: int
{
    case DAILY = 1;
    case WEEKLY = 7;
    case MONTHLY = 30;
}
