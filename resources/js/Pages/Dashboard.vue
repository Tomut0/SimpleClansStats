<script setup>
import SCSLayout from "@/Layouts/SCSLayout.vue";
import {Head} from "@inertiajs/vue3";
import {__} from "../trans.js";
import DashboardLayout from "@/Layouts/DashboardLayout.vue";

defineOptions({layout: [SCSLayout, DashboardLayout]});
</script>

<template>
    <Head :title="__('pages.dashboard')"/>
</template>
