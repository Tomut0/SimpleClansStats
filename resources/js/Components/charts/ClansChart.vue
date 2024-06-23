<script setup>

import ChartWrapper from "@/Components/charts/ChartWrapper.vue";

import {onMounted} from "vue";
import {Chart} from "chart.js/auto";
import {__} from "@/trans.js";

const props = defineProps({
    statistics: {
        type: Object,
        required: true
    },
    unit: {
        type: String,
        required: true
    }
});

// Prepare data for chart
const chartData = Object.entries(props.statistics).map(([date, values]) => ({
    x: date,
    clans: values.clans
}));

onMounted(() => {
    new Chart(clans, {
        type: 'bar',
        data: {
            datasets: [{
                data: chartData.map(item => ({x: item.x, y: item.clans})),
                fill: true,
                backgroundColor: 'rgb(47, 77, 237, 0.6)',
                borderColor: 'rgb(47, 77, 237)',
                borderWidth: 1,
            }]
        },
        options: {
            plugins: {
                legend: {
                    display: false,
                }
            },
            responsive: true,
            scales: {
                x: {
                    type: 'time',
                    time: {
                        unit: props.unit,
                    },
                }
            },
        }
    });
})
</script>

<template>
    <ChartWrapper>
        {{ __('components.charts.clans') }}
        <template #chart>
            <canvas id="clans"></canvas>
        </template>
    </ChartWrapper>
</template>
