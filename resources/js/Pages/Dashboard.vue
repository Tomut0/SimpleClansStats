<script setup>
import {Head} from '@inertiajs/vue3';
import SCHeader from "@/Components/SCHeader.vue";
import {__} from "../trans.js";

</script>

<template>
    <Head :title="__('pages.dashboard')"/>

    <div class="px-16 pt-12">
        <SCHeader/>
    </div>
</template>
