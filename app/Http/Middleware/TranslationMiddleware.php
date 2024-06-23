<?php

namespace App\Http\Middleware;

use App\Http\Controllers\TranslationController;
use Closure;
use Illuminate\Http\Request;
use Inertia\Inertia;
use Symfony\Component\HttpFoundation\Response;

class TranslationMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param Request $request
     * @param \Closure(Request): (Response) $next
     * @return Response
     */
    public function handle(Request $request, Closure $next): Response
    {
        $this->setCurrency();
        app()->setLocale(config('app.locale'));
        if (session()->has('locale')) {
            app()->setLocale(session('locale'));
            return $next($request);
        }

        if (empty($_SERVER['HTTP_ACCEPT_LANGUAGE'])) {
            return $next($request);
        }

        // set language based on browser preferences
        $language = $_SERVER['HTTP_ACCEPT_LANGUAGE'];
        $languagePreferences = explode(',', $language)[0];

        if (str_contains($languagePreferences, "-")) {
            $languagePreferences = explode("-", $languagePreferences)[0];
        }

        $filteredArray = array_filter(json_decode(TranslationController::jsonMessages(), true), function ($value) use ($languagePreferences) {
            // Return true if the key contains the substring
            return str_contains($value, $languagePreferences);
        }, ARRAY_FILTER_USE_KEY);

        if (array_key_first($filteredArray) !== null) {
            app()->setLocale(array_key_first($filteredArray));
        }

        return $next($request);
    }

    private function setCurrency(): void
    {
        Inertia::share('currency', [
            'code' => config("scstats.currency", "USD"),
            'symbol' => config("scstats.currencySymbol", "$"),
        ]);
    }
}
