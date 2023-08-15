<script setup>
import {MagnifyingGlassIcon} from "@heroicons/vue/24/solid/index.js";
import {debounce} from "@/helpers.js";
import ExpandedButton from "@/Components/ExpandedButton.vue";

const props = defineProps({
    placeholder: String,
    searchStyles: String,
    iconStyles: String,
    expanded: {
        type: Boolean,
        default: false
    },
    disabled: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['search']);

const debouncedEmit = debounce((input) => {
    emit('search', input.target.value)
}, 250);

</script>

<template>
    <ExpandedButton is-absolute-icon :expanded="expanded" :disabled="disabled">
        <template #icon>
            <MagnifyingGlassIcon class="w-6 h-6 text-white" :class="iconStyles"/>
        </template>

        <input type="text" :placeholder="placeholder" :class="searchStyles" @input="debouncedEmit"
               class="border-gray-400 rounded-md pl-7">
    </ExpandedButton>
</template>

<style scoped>

</style>
