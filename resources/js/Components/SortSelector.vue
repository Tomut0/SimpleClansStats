<script setup>
import ExpandedButton from "@/Components/ExpandedButton.vue";
import {CircleStackIcon, FireIcon, UserGroupIcon} from "@heroicons/vue/24/solid/index.js";
import {__} from "../trans.js";
import {router, usePage} from "@inertiajs/vue3";
import {reactive} from "vue";

const sortTypes = usePage().props.sortTypes;
const icons = [FireIcon, CircleStackIcon, UserGroupIcon];

/**
 * Map sort types to icon components
 * @type {Object}
 */
const sortIcons = sortTypes.reduce((acc, key, i) => {
    acc[key] = icons[i];
    return acc;
}, {});

let defaults = reactive({
    sortType: sortTypes[0]
});

/**
 * Handle selecting a sort type
 *
 * @param {Object} shown - Reactive toggle state
 * @param {string} sortType - Selected sort type
 */
function selectType(shown, sortType) {
    defaults.sortType = sortType;
    shown.value = false;

    router.visit("/", {
        data: {
            sortBy: sortType
        }
    });
}
</script>

<template>
    <div>
        <ExpandedButton v-for="(value, key) in sortIcons"
                        :after-expand="(shown) => {selectType(shown, key);}"
                        :is-locked="defaults.sortType === key"
                        direction="right"
                        class="px-4 py-2 text-white inline-flex bg-darkside-900 sortType"
                        :class="{'bg-indigo-500': defaults.sortType === key}">
            <template #icon>
                <component :is="value" class="w-6 h-6 transition hover:text-darkside-500"></component>
            </template>

            <span class="ml-2">{{ __(`components.sortselector.${key}`) }}</span>
        </ExpandedButton>
    </div>
</template>

<style scoped>
.sortType:first-child {
    border-radius: 0.5rem 0 0 0.5rem;
}

.sortType:last-child {
    border-radius: 0 0.5rem 0.5rem 0;
}

.active {
    @apply bg-indigo-500
}
</style>
