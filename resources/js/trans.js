import {usePage} from "@inertiajs/vue3";

export function __(key, replacement = {}, locale = usePage().props.locale) {
    const translations = JSON.parse(usePage().props.translations);

    let message = key.split('.').reduce((t, i) => (t && t[i]) ?
        t[i] : __(key, replacement, usePage().props.fallback_locale), translations[locale]);


    Object.keys(replacement).forEach(key => {
        message = message.replace(`:${key}`, replacement[key]);
    });

    return message;
}
