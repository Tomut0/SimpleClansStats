<script setup>
import {__} from "@/trans.js";

const props = defineProps({
    columns: {
        type: Array,
        default: []
    },
    rows: {
        required: true,
        type: Array,
        default: []
    },
    showHeader: {
        type: Boolean,
        default: true
    },
    style: {
        type: String,
        default: "filled"
    }
});

const columnsValues = () => {
    const rowValues = [...props.rows];
    const columnValues = [...props.columns].map(c => __(c));
    const maxColumns = Math.max(columnValues.length, ...rowValues.map(r => r.elements.length));

    // Fill with empty strings if the row is shorter than maxColumns
    while (columnValues.length < maxColumns) {
        columnValues.unshift("");
    }

    return columnValues;
};
</script>

<template>
    <table :class="`table__${style}`">
        <thead v-if="showHeader">
        <tr>
            <th v-for="(column, index) in columnsValues()" :key="index">{{ column }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(row, rowIndex) in rows" :key="rowIndex" @click.prevent.stop="row.onclick">
            <td v-for="(value, colIndex) in row.elements" :key="colIndex">
                <component v-if="typeof value === 'object'" :is="value"/>
                <div v-else v-html="value"></div>
            </td>
        </tr>
        </tbody>
    </table>
</template>

<style scoped>
.table__filled {
    @apply table-auto text-white
}

.table__filled > thead {
    @apply text-gray-300 bg-darkside-900 h-12
}

.table__filled > tbody > tr {
    @apply h-12 text-gray-200 text-center cursor-pointer hover:bg-darkside-500 odd:bg-darkside-600 even:bg-darkside-700
}
</style>
