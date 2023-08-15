<script setup>

import {ref} from "vue";

const props = defineProps({
    direction: {
        type: String,
        default: "right"
    },
    expanded: {
        type: Boolean,
        default: true
    },
    disabled: {
        type: Boolean,
        default: false,
    },
    isAbsoluteIcon: Boolean
});

const shown = ref(false);

function toggle() {
    if (props.disabled) {
        return;
    }

    shown.value = !shown.value;
}
</script>

<template>
    <button>
        <span class="flex items-center" v-if="direction === 'left'">
            <span v-show="shown || expanded"><slot></slot></span>
            <span @click.prevent.stop="toggle" :class="{'absolute':isAbsoluteIcon && shown}"><slot
                name="icon"></slot></span>
        </span>
        <span class="flex items-center" v-else-if="direction === 'right'">
            <span @click.prevent.stop="toggle" :class="{'absolute':isAbsoluteIcon && shown}"><slot
                name="icon"></slot></span>
            <span v-show="shown || expanded"><slot></slot></span>
        </span>
    </button>
</template>

<style scoped>

</style>
