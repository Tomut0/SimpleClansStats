<script setup>

import SCHeader from "@/Components/SCHeader.vue";
import {Head} from "@inertiajs/vue3";
import SCFooter from "@/Components/SCFooter.vue";
</script>

<template>
    <Head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
        <meta name="robots" content="noindex">
        <link rel="shortcut icon" href="/favicons/favicon.ico">
    </Head>

    <div class="px-6 lg:px-16 pt-12">
        <SCHeader/>
        <slot />
    </div>

    <SCFooter />
</template>

<style scoped>

</style>
