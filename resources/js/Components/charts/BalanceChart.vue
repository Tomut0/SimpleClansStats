<script setup>
import ChartWrapper from "@/Components/charts/ChartWrapper.vue";
import {Chart} from "chart.js/auto";
import {onMounted} from "vue";
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
    balance: values.balance
}));

onMounted(() => {
    new Chart(clanBalance, {
        type: 'line',
        data: {
            datasets: [{
                data: chartData.map(item => ({x: item.x, y: item.balance})),
                fill: true,
                borderColor: 'rgb(47, 77, 237)',
                tension: 0.1
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
                y: {
                    beginAtZero: true,
                },
                x: {
                    type: 'time',
                    time: {
                        unit: props.unit,
                    },
                }
            },
        }
    });
});

</script>

<template>
    <ChartWrapper>
        {{ __('components.charts.avg_balance') }}
        <template #chart>
            <canvas id="clanBalance"></canvas>
        </template>
    </ChartWrapper>
</template>

