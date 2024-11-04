import {defineAsyncComponent} from "vue";
import {router, usePage} from "@inertiajs/vue3";

function debounce(func, delay) {
    let timerId;

    return function (...args) {
        const context = this;

        clearTimeout(timerId);

        timerId = setTimeout(() => {
            func.apply(context, args);
        }, delay);
    };
}

/**
 * Get the value of a query parameter from a URL
 *
 * @param url
 * @param key
 * @returns {string}
 */
function queryValue(url, key) {
    const queryString = url.split('?')[1];
    const params = new URLSearchParams(queryString);
    return params.get(key);
}

function visitViaQuery(href, queries = [], only = []) {
    const currentParams = new URLSearchParams(window.location.search);

    for (const query of queries) {
        if (query && query.name) {
            if (!query.value) {
                currentParams.delete(query.name);
                continue;
            }

            currentParams.set(query.name, query.value);
        }
    }

    router.visit(href, {
        data: currentParams,
        preserveScroll: true,
        only: only
    });
}

function formatNumber(num, locale = 'en-US', style = 'currency', minimumFractionDigits = 0, maximumFractionDigits = 1) {
    locale = locale.replace('_', '-');

    const currency = usePage().props.currency.code || 'USD';
    const symbol = usePage().props.currency.symbol || '$';

    const formatter = new Intl.NumberFormat([locale, "en-US"], {
        style: style,
        notation: 'compact',
        currency: currency,
        currencyDisplay: "narrowSymbol",
        compactDisplay: 'short',
        minimumFractionDigits: minimumFractionDigits,
        maximumFractionDigits: maximumFractionDigits
    });

    // slice and replace with currency symbol
    if (style === 'currency') {
        return formatter.formatToParts(num).map(part => {
            if (part.type === 'currency') {
                return symbol;
            }
            return part.value;
        }).join('');
    }

    return formatter.format(num);
}

/**
 * Bake an icon component from the Heroicons Vue library
 *
 * @see https://heroicons.com/
 * @param iconName
 */
function bakeIcon(iconName) {
    return defineAsyncComponent(() => {
        return import(`../../node_modules/@heroicons/vue/24/solid/esm/${iconName}.js`);
    })
}

function chartUnit(unit) {
    switch (unit) {
        case 'daily':
            return 'hour';
        case 'weekly':
            return 'day';
        case 'monthly':
            return 'day';
        case 'yearly':
            return 'month';
        default:
            return 'hour';
    }
}

function omit(obj, keys) {
    return keys.reduce((acc, key) => {
        if (!acc[key]) {
            acc[key] = obj.value[key];
        }
        return acc;
    });
}

function getKeyAndValue(obj) {
    const key = Object.keys(obj)[0];
    return {key: key, value: obj[key]};
}

function createRow(elements, onclick) {
    return {
        elements: elements || [],   // default to an empty array if no value is provided
        onclick: onclick || function() {} // default to a no-op function if none is provided
    };
}

export {debounce, queryValue, formatNumber, bakeIcon, getKeyAndValue, chartUnit, visitViaQuery, createRow};
