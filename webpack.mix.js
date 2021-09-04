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

mix.sass('resources/sass/style.scss', 'public/assets/css', {
    sassOptions: {
        outputStyle: 'compressed'
    }
}).js(['resources/js/detail_modal.js', 'resources/js/filter.js', 'resources/js/trim_names.js'], 'public/assets/js/app.js').sourceMaps();

if (mix.inProduction()) {
    mix.version();
}
