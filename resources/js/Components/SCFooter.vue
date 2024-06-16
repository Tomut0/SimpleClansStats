<script setup>

import ApplicationLogo from "@/Components/ui/ApplicationLogo.vue";
import {__} from "../trans.js";
</script>

<template>
    <footer class="flex justify-between bg-darkside-500 px-6 lg:px-16 py-12 mt-12">
        <ApplicationLogo/>

        <div class="flex flex-col text-right">
            <span class="text-gray-200 font-medium text-xl">{{ __('components.footer.copyright', { year: new Date().getFullYear() }) }}</span>
            <span class="text-gray-300 text-lg">{{ __('components.footer.credits') }}</span>
        </div>
    </footer>
</template>

<style scoped>

</style>
