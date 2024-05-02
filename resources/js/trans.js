import {usePage} from "@inertiajs/vue3";


function isFallback(locale) {
    return locale === usePage().props.fallback_locale;
}

export function __(key, replacement = {}, locale = usePage().props.locale) {
    if (import.meta.env.MODE === 'development') {
        console.log('Loaded translation key:', key);
    }

    if (!key) {
        throw new Error("Translation key is undefined or null");
    }

    const translations = JSON.parse(usePage().props.translations);

    let message = key.split('.').reduce((t, i) => (t && t[i]) ?
        t[i] : isFallback(locale) ? false : __(key, replacement, usePage().props.fallback_locale), translations[locale]);

    if (!message) {
        console.warn(`Translation key [${key}] wasn't found`);
        return key;
    }

    Object.keys(replacement).forEach(key => {
        message = message.replace(`:${key}`, replacement[key]);
    });

    return message;
}
