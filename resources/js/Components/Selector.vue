<script setup>

import ExpandedButton from "@/Components/ExpandedButton.vue";
import {router, usePage} from "@inertiajs/vue3";
import {computed, ref} from "vue";
import {__} from "@/trans.js";
import {bakeIcon, queryValue} from "@/helpers.js";

const props = defineProps({
    items: {
        type: Object,
        required: true
    },
    href: {
        type: String,
        default: '/'
    },
    query: {
        type: String,
    },
    canBakeIcons: {
        type: Boolean,
        default: false
    }
});

const selected = ref(queryValue(usePage().url, props.query) ?? Object.keys(props.items)[0]);

const bakeIcons = () => computed(() => {
    return Object.entries(props.items).reduce((acc, value) => {
        acc[value[0]] = bakeIcon(value[1].icon);
        return acc
    }, {})
});

let icons = null;
if (props.canBakeIcons) {
    icons = bakeIcons();
}

const select = (shown, selectedItem) => {
    selected.value = selectedItem;
    if (shown) {
        shown.value = false;
    }

    const currentParams = new URLSearchParams(window.location.search);
    currentParams.set(props.query, selectedItem);

    router.visit(props.href, {
        data: currentParams,
    });
};

</script>

<template>
    <div>
        <ExpandedButton v-if="canBakeIcons" is="button"
                        v-for="(item, key) in items"
                        :after-expand="(shown) => select(shown, key)"
                        direction="right"
                        class="selector-button"
                        :is-locked="selected === key"
                        :class="{'!bg-indigo-500': selected === key}">
            <template #icon>
                <component :is="icons[key]" class="w-6 h-6 transition mr-2 hover:text-darkside-500"/>
            </template>

            {{ __(item.translation) }}
        </ExpandedButton>

        <button type="submit" v-else
                v-for="(item, key) in items"
                @click="select(false, key)"
                class="selector-button"
                :class="{'!bg-indigo-500': selected === key}">

            {{ __(item.translation) }}
        </button>
    </div>
</template>

<style scoped>
.selector-button {
    @apply px-4 py-2 text-white inline-flex bg-darkside-900
}

.selector-button:first-child {
    border-radius: 0.5rem 0 0 0.5rem;
}

.selector-button:last-child {
    border-radius: 0 0.5rem 0.5rem 0;
}
</style>
