<?php

namespace App\Http\Controllers;

use App\Enums\CacheKeys;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\File;
use Locale;

class TranslationController extends Controller
{

    public function update(Request $request): RedirectResponse
    {
        $validated = $request->validate([
            'locale' => ['regex:~^[a-z]{2}_[A-Z]{2}|en$~']
        ]);

        Session()->put('locale', $validated['locale']);

        return back();
    }

    public static function jsonMessages()
    {
        return Cache::rememberForever(CacheKeys::LanguageMap->name, function () {
            $files = File::allFiles(lang_path());

            $locales = [];

            foreach ($files as $file) {
                $directory = $file->getRelativePath();

                $language = Locale::getDisplayLanguage($directory, $directory);
                $region = Locale::getDisplayRegion($directory, $directory);
                $displayName = empty($region) ? "$language" : "$language ($region)";

                $displayName = mb_convert_case($displayName, MB_CASE_TITLE, 'UTF-8');
                $locales[$directory][$file->getBasename('.php')] = trans($file->getBasename('.php'), [], $directory);
                $locales[$directory]['displayName'] = $displayName;
            }

            return json_encode($locales);
        });
    }

    public function localizationMap()
    {
        return Cache::rememberForever(CacheKeys::LocalizationMap->name, function () {
            // get default locale (English 'en' by default)
            $default = config("app.locale");

            $languageMap = json_decode($this->jsonMessages(), true);

            // get default strings in array-like view
            // ex.: array [ ["key1"] => ["value1"], ["key2"] => ["value2"] ]
            $defaultLocale = $languageMap[$default];

            $percentMap = [];
            foreach ($languageMap as $key => $strings) {
                // Compare locales strings with default locale
                $comparison = $this->arrayPercentageDifference($defaultLocale, $strings);

                // get localization percent between default and current locale
                if ($key === $default) {
                    $percentDifference = 100;
                } else {
                    $percentDifference = 100 - ($comparison['matched'] / $comparison['total']) * 100;
                }

                // add to map, formats percent double to string
                $percentMap[$key] = number_format($percentDifference) . '%';
            }

            return $percentMap;
        });
    }

    // Function to recursively compare two arrays and calculate the percentage difference
    private function arrayPercentageDifference($array1, $array2): array
    {
        $totalKeys = 0;
        $matchedKeys = 0;

        foreach ($array1 as $key => $value) {
            if (is_array($value)) {
                if (isset($array2[$key]) && is_array($array2[$key])) {
                    $subArrayResult = $this->arrayPercentageDifference($value, $array2[$key]);
                    $matchedKeys += $subArrayResult['matched'];
                    $totalKeys += $subArrayResult['total'];
                }
            } else {
                $totalKeys++;
                if (isset($array2[$key]) && $array2[$key] === $value) {
                    $matchedKeys++;
                }
            }
        }

        return ['matched' => $matchedKeys, 'total' => $totalKeys];
    }
}
