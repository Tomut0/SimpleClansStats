<script setup>
import {CheckIcon, XMarkIcon, CalendarDaysIcon} from "@heroicons/vue/24/solid/index.js";
import ShieldIcon from "@/Components/icons/ShieldIcon.vue";
import {bakeIcon, formatNumber, queryValue} from "@/helpers.js";
import {usePage} from "@inertiajs/vue3";
import {__} from "../trans.js";
import {extractFirstColor, getColorByBgColor} from "@/colors.js";

const {locale} = usePage().props;

const props = defineProps({
    clan: {
        type: Object,
        required: true
    },
    positionIcon: {
        required: true,
        type: String,
        default: 'ShieldIcon'
    },
    sortBy: {
        type: String,
        required: true,
    },
    positionColors: {
        type: String,
        default: 'text-white'
    }
});

let value;
if (props.sortBy === 'balance') {
    value = formatNumber(props.clan[props.sortBy], locale.value, "currency");
} else {
    value = formatNumber(props.clan[props.sortBy], locale.value, "decimal", 0, 2);
}

const textColor = getColorByBgColor(extractFirstColor(props.clan.color_tag));
</script>

<template>
    <div class="flex flex-col items-center relative h-full">
        <div class="position-text flex flex-col items-center justify-evenly h-2/3">
            <div class="flex flex-col space-x-4 items-center xl:flex-row">
                <span v-if="clan.verified" class="flex items-center">
                    <CheckIcon class="w-4 h-4"/>
                        {{ __('general.clan.verified') }}
                </span>
                <span v-else class="flex items-center">
                    <XMarkIcon class="w-4 h-4"/>
                    {{ __('general.clan.unverified') }}
                </span>

                <span class="flex items-center">
                    <CalendarDaysIcon class="w-4 h-4"/>
                    {{ clan.formatted_founded }}
                </span>
            </div>

            <div class="flex flex-col items-center">
                <span class="font-bold uppercase lg:text-xl xl:text-3xl">{{ clan.tag }}</span>
                <span class="lg:text-lg xl:text-xl">{{ clan.name }}</span>
            </div>

            <div class="flex justify-center items-center font-semibold text-base xl:text-3xl">
                <component :is="bakeIcon(positionIcon)" class="w-4 h-4 mr-2 xl:w-8 xl:h-8"/>
                {{ value }}
            </div>
        </div>

        <ShieldIcon class="h-full w-2/3" :class="positionColors" :fill="extractFirstColor(props.clan.color_tag)"/>
    </div>
</template>

<style scoped>
.position-text {
    position: absolute;
    top: 75%;
    left: 50%;
    transform: translate(-50%, -75%);
    color: v-bind(textColor);
}
</style>
