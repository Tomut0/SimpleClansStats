<script setup>

import SortSelector from "@/Components/SortSelector.vue";
import SCHeader from "@/Components/SCHeader.vue";
import IntervalSelector from "@/Components/IntervalSelector.vue";
import {Head} from "@inertiajs/vue3";
</script>

<template>
    <Head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
        <meta name="robots" content="noindex">
        <link rel="shortcut icon" href="/favicons/favicon.ico">
    </Head>

    <div class="px-16 pt-12">
        <SCHeader/>

        <div class="flex justify-between mt-8">
            <SortSelector/>
            <IntervalSelector/>
        </div>

        <slot/>
    </div>
</template>

<style scoped>

</style>
