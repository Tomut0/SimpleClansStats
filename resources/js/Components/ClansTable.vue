<script setup>

import {bakeIcon, createRow, formatNumber, visitViaQuery} from "@/helpers.js";
import {usePage} from "@inertiajs/vue3";
import {h} from "vue";
import DataTable from "@/Components/DataTable.vue";
import {addColors} from "@/colors.js";

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

const positionArrow = (clan) => {
    if (clan.position_status === 'raised') {
        return h(bakeIcon('ChevronUpIcon'), {class: 'w-4 h-4 text-green-400 ml-4'});
    } else if (clan.position_status === 'lowered') {
        return h(bakeIcon('ChevronDownIcon'), {class: 'w-4 h-4 text-red-400 ml-4'});
    } else {
        return h(bakeIcon('MinusIcon'), {class: "w-4 h-4 text-grey-400 ml-4"})
    }
};

const sortValue = (clan) => {
    if (props.currentSort.key === 'balance') {
        return formatNumber(clan.balance, locale, "currency");
    } else {
        return formatNumber(clan[props.currentSort.key], locale, "decimal", 0, 2);
    }
};

const rows = [...props.clans].map((clan, index) => {
    const row = [positionArrow(clan), index + props.addToIndex, addColors(clan.color_tag), clan.name, sortValue(clan)];
    return createRow(row, () => visitViaQuery(route('leaderboard.index'), [{
        name: 'clanTag',
        value: clan.tag
    }], ['queryClan', 'selectors']))
});

const columns = ['#', 'general.clan.tag', 'general.clan.name', props.currentSort.value.translation];
</script>

<template>
    <DataTable :rows="rows" :columns="columns"/>
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
