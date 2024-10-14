<script setup>

import {__} from "@/trans.js";
import {ChevronUpIcon, ChevronDownIcon, MinusIcon} from "@heroicons/vue/24/solid/index.js";
import {addColors} from "../colors.js";
import {formatNumber, visitViaQuery} from "@/helpers.js";
import {usePage} from "@inertiajs/vue3";

const {locale} = usePage().props;

const props = defineProps({
    clans: {
        type: Array,
        required: true
    },
    currentSort: {
        type: Object,
        required: true
    },
    addToIndex: {
        type: Number,
        default: 1
    }
});

</script>

<template>
    <table class="table-auto rounded">
        <thead class="text-gray-300 bg-darkside-900 h-12">
        <tr>
            <th></th>
            <th>#</th>
            <th>{{ __('general.clan.tag') }}</th>
            <th>{{ __('general.clan.name') }}</th>
            <th>{{ __(currentSort.value.translation) }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(clan, indx) in clans" :key="clan.tag"
            @click="visitViaQuery(route('dashboard'), [{name: 'clanTag', value: clan.tag}], ['queryClan', 'selectors'])"
            class="h-12 text-gray-200 text-center cursor-pointer hover:bg-darkside-500 odd:bg-darkside-600 even:bg-darkside-700">
            <td v-if="clan.position_status === 'raised'">
                <ChevronUpIcon class="w-4 h-4 text-green-400 ml-4"/>
            </td>
            <td v-else-if="clan.position_status === 'lowered'">
                <ChevronDownIcon class="w-4 h-4 text-red-400 ml-4"/>
            </td>
            <td v-else>
                <MinusIcon class="w-4 h-4 text-gray-400 ml-4"></MinusIcon>
            </td>
            <td>{{ indx + addToIndex }}</td>
            <td v-html="addColors(clan.color_tag)"></td>
            <td>{{ clan.name }}</td>
            <td v-if="props.currentSort.key === 'balance'">
                {{ formatNumber(clan.balance, locale, "currency") }}
            </td>
            <td v-else>
                {{ formatNumber(clan[props.currentSort.key], locale, "decimal", 0, 2) }}
            </td>
        </tr>
        </tbody>
    </table>
</template>

<style scoped>
th:first-of-type {
    border-top-left-radius: 12px;
}

th:last-of-type {
    border-top-right-radius: 12px;
}

tr:last-of-type td:first-of-type {
    border-bottom-left-radius: 12px;
}

tr:last-of-type td:last-of-type {
    border-bottom-right-radius: 12px;
}
</style>
