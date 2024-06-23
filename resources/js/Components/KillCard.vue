<script setup>

import {__} from "../trans.js";
import SwordIcon from "@/Components/icons/SwordIcon.vue";
import {addColors} from "../colors.js";

const props = defineProps({
    attacker: {
        type: Object,
        required: true
    },
    victim: {
        type: Object,
        required: true
    },

});
</script>

<template>
    <div class="px-4 py-3 bg-darkside-600 rounded-xl">
        <div class="flex justify-between mb-2">
            <span class="text-red-400">{{ __('general.kills.attacker') }}</span>
            <span class="text-blue-400">{{ __('general.kills.victim') }}</span>
        </div>
        <hr>
        <div class="flex justify-between items-center relative mt-4">
            <div>
                <div class="text-left text-sm" v-html="addColors(props.attacker.tag)"></div>
                <div>{{ props.attacker.name }}</div>
            </div>
            <SwordIcon class="w-6 h-6 text-darkside-400 icon-position"/>
            <div>
                <div class="text-right text-sm" v-html="addColors(props.victim.tag)"></div>
                <div>{{ props.victim.name }}</div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.icon-position {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}
</style>
