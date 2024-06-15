import {defineAsyncComponent} from "vue";

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

function formatNumber(num, locale = 'en-US', style = 'currency', minimumFractionDigits = 0, maximumFractionDigits = 1) {
    const formatter = new Intl.NumberFormat(locale, {
        style: style,
        notation: 'compact',
        currency: 'USD', // this is the fallback value
        compactDisplay: 'short',
        minimumFractionDigits: minimumFractionDigits,
        maximumFractionDigits: maximumFractionDigits
    });

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

function omit(obj, keys) {
    return keys.reduce((acc, key) => {
        if (!acc[key]) {
            acc[key] = obj.value[key];
        }
        return acc;
    });
}

export {debounce, queryValue, formatNumber, bakeIcon};
