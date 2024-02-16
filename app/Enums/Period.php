<?php

namespace App\Enums;

enum Period: int
{
    case Daily = 1;
    case Weekly = 7;
    case Monthly = 30;
}
