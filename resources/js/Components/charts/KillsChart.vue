<script setup>

import {onMounted} from "vue";
import {Chart} from "chart.js/auto";
import ChartWrapper from "@/Components/charts/ChartWrapper.vue";
import {__} from "../../trans.js";

const props = defineProps({
    statistics: {
        type: Object,
        required: true
    }
});

// Prepare data for chart
const chartData = Object.entries(props.statistics.kills).map(([key, value]) => ({
    type: key,
    count: value
}));

onMounted(() => {
    new Chart(kills, {
        type: 'doughnut',
        data: {
            labels: chartData.map(row => row.type),
            datasets: [
                {
                    label: __('general.killed'),
                    data: chartData.map(row => row.count),
                    backgroundColor: [
                        'rgb(47, 77, 237)',
                        'rgb(60,237,47)',
                        'rgb(237,47,63)',
                    ],
                    borderWidth: 0,
                },
            ],
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    display: true,
                    position: 'bottom',
                    labels: {
                        padding: 24,
                        color: 'white',
                    }
                }
            },
            cutout: '30%',
        }
    });
});
</script>

<template>
    <ChartWrapper>
        {{ __('components.charts.kills_by_type') }}
        <template #chart>
            <canvas id="kills"></canvas>
        </template>
    </ChartWrapper>
</template>
