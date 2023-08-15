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

function omit(obj, keys) {
    return keys.reduce((acc, key) => {
        if (!acc[key]) {
            acc[key] = obj.value[key];
        }
        return acc;
    });
}

export {debounce, omit};
