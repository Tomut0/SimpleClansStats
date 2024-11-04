<script setup>
import SCSLayout from "@/Layouts/SCSLayout.vue";
import {Head, usePage} from "@inertiajs/vue3";
import {__} from "../trans.js";
import ShieldPosition from "@/Components/ShieldPosition.vue";
import {chartUnit, getKeyAndValue, queryValue} from "@/helpers.js";
import ClansTable from "@/Components/ClansTable.vue";
import Carousel from "@/Components/Carousel.vue";
import {computed, defineAsyncComponent, ref} from "vue";
import KillsChart from "@/Components/charts/KillsChart.vue";
import BalanceChart from "@/Components/charts/BalanceChart.vue";
import ClansChart from "@/Components/charts/ClansChart.vue";
import ClanModal from "@/Components/modals/ClanModal.vue";
import LeaderboardLayout from "@/Layouts/LeaderboardLayout.vue";

const {selectors, clans, lastKills, statistics, queryClan} = usePage().props;

const currentClan = ref(queryClan);

const currentSort = getKeyAndValue(selectors.current.sortSelector);
const currentInterval = getKeyAndValue(selectors.current.intervalSelector);
const unit = chartUnit(currentInterval.value);

const limit = queryValue(usePage().url, 'limit') ?? 10;

const bestClans = clans.slice(0, 3);

const width = ref(window.innerWidth);
window.addEventListener('resize', () => width.value = window.innerWidth);

const otherClans = computed(() => width.value > 1024 ? clans.slice(3, limit) : clans);
const tableIndex = computed(() => width.value > 1024 ? bestClans.length + 1 : 1);
const cardsCount = computed(() => width.value > 1536 ? 3 : 1);

const positionColors = ['text-gray-400', 'text-yellow-400', 'text-orange-400'];

// swap first and second place if there are 3 best positions
[bestClans[0], bestClans[1]] = [bestClans[1], bestClans[0]];

const killCards = lastKills.map(kill => {
    return {
        component: defineAsyncComponent(() => import('@/Components/KillCard.vue')),
        props: {
            attacker: kill.attacker,
            victim: kill.victim
        }
    }
});

defineOptions({layout: [SCSLayout, LeaderboardLayout]});
</script>

<template>
    <Head :title="__('pages.leaderboard')"/>

    <div class="grid-cols-3 grid-rows-1 mt-4 relative hidden lg:grid">
        <ShieldPosition v-for="(clan, index) in bestClans" :key="clan.tag"
                        :position-icon="currentSort.value.icon" :position-colors="positionColors[index]"
                        :clan="clan" :sort-by="currentSort.key"
                        :class="{'mt-24': index !== 1}" class="relative"/>
    </div>

    <div v-motion-fade-visible class="grid grid-flow-col auto-cols-fr gap-8 mt-24">
        <ClansTable v-if="otherClans.length > 3" v-model="currentClan" :clans="otherClans"
                       :current-sort="currentSort"
                       :add-to-index="tableIndex"/>
    </div>

    <div v-motion-fade-visible class="bg-darkside-900 p-4 rounded-xl mt-8">
        <h3 class="text-gray-400 text-xl font-bold text-center mb-4">{{ __('general.kills.last') }}</h3>
        <Carousel :items="killCards" :show-count="cardsCount"
                  carousel-class="grid grid-flow-col auto-cols-fr w-full gap-4 text-white text-base lg:text-xl"/>
    </div>

    <div v-motion-fade-visible class="grid grid-flow-row auto-cols-fr lg:grid-flow-col lg:grid-cols-[repeat(auto-fit,minmax(0,1fr))] mt-8 gap-8">
        <KillsChart v-if="statistics && !Array.isArray(statistics.kills)" :statistics="statistics"/>
        <BalanceChart v-if="statistics && statistics.byTime" :statistics="statistics.byTime" :unit="unit"/>
        <ClansChart v-if="statistics && statistics.byTime" :statistics="statistics.byTime" :unit="unit"/>
    </div>

    <ClanModal v-model="currentClan" />
</template>
