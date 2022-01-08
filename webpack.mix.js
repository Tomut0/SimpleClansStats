const mix = require('laravel-mix');
/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

// TODO: Replace all SCSS to PostCSS with plugins
mix.js('resources/js/app.js', 'public/js').sourceMaps();
mix.sass('resources/sass/style.scss', 'public/css', {
    sassOptions: {
        outputStyle: 'compressed'
    }
}, [require("tailwindcss")]);

if (mix.inProduction()) {
    mix.version();
}
