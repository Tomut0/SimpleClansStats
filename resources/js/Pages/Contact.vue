<script setup>
import {Head, Link} from "@inertiajs/vue3";
import Translator from "@/Components/Translator.vue";

</script>

<template>
    <Head title="Translate test"/>

    {{ $page.props.errors }}
    {{ $page.props.locale }}
    {{ $page.props.applicationLocale }}

    <h1>
        <Translator></Translator>
    </h1>

    <br>
    <Link :href="route('welcome')">Go back</Link>
    <div class="min-h-screen"></div>
</template>

<style scoped>

</style>
