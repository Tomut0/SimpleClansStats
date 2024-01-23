<script setup>
import {inject} from "vue";

const appName = inject('appName');
</script>

<template>
    <div class="flex items-center">
        <img src="images/logo.svg" alt="logo">
        <span class="ml-6 text-white font-bold text-3xl font-sans">{{ appName }}</span>
    </div>
</template>
