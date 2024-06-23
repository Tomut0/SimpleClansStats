<script setup>
import Selector from "@/Components/Selector.vue";
import {usePage} from "@inertiajs/vue3";

const selectors = usePage().props.selectors;
</script>

<template>
    <div class="flex flex-col items-center justify-between mt-8 space-y-4 lg:flex-row">
        <Selector :items="selectors.all.sortSelector" query="sortBy" can-bake-icons/>
        <Selector :items="selectors.all.intervalSelector" query="period"/>
    </div>

    <slot/>
</template>

<style scoped>

</style>
