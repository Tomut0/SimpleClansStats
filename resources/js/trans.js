import {usePage} from "@inertiajs/vue3";

export function __(key, replacement = {}) {
    const translations = JSON.parse(usePage().props.translations);

    let message = key.split('.').reduce((t, i) => t[i] || null, translations[usePage().props.locale]);

    Object.keys(replacement).forEach(key => {
        message = message.replace(`:${key}`, replacement[key]);
    });

    return message;
}
