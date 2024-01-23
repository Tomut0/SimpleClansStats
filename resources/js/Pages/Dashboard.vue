<script setup>
import SCSLayout from "@/Layouts/SCSLayout.vue";
import {Head} from "@inertiajs/vue3";
import {__} from "../trans.js";

defineOptions({layout: SCSLayout});
</script>

<template>
    <Head :title="__('pages.dashboard')"/>

</template>
