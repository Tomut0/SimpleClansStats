<script setup>
import Selector from "@/Components/Selector.vue";
import {usePage} from "@inertiajs/vue3";

const selectors = usePage().props.selectors;
</script>

<template>
    <div class="flex justify-between mt-8">
        <Selector :items="selectors.all.sortSelector" query="sortBy" can-bake-icons/>
        <Selector :items="selectors.all.intervalSelector" query="period"/>
    </div>

    <slot/>
</template>

<style scoped>

</style>
