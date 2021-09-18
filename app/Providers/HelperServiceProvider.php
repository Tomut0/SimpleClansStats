<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;

class HelperServiceProvider extends ServiceProvider
{
    /**
     * Register services.
     *
     * @return void
     */
    public function register()
    {
        $helpers_path = app_path('Helpers');
        $helpers = array_slice(scandir($helpers_path), 2);

        foreach ($helpers as $helper) {
            require_once ($helpers_path . '/' . $helper);
        }

    }

    /**
     * Bootstrap services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }
}
