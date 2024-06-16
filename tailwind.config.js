import defaultTheme from 'tailwindcss/defaultTheme';
import forms from '@tailwindcss/forms';

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        './vendor/laravel/framework/src/Illuminate/Pagination/resources/views/*.blade.php',
        './storage/framework/views/*.php',
        './resources/views/**/*.blade.php',
        './resources/js/**/*.vue',
    ],

    theme: {
        extend: {
            fontFamily: {
                sans: ['Ubuntu', ...defaultTheme.fontFamily.sans],
            },
            colors: {
                darkside: {
                    400: '#8E96EB',
                    500: '#6E63F1',
                    600: '#33316B',
                    700: '#2C2B5E',
                    900: '#121434'
                }
            }
        },
    },

    plugins: [forms],
};
