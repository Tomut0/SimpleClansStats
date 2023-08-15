<script setup>
import {Head, Link, usePage} from '@inertiajs/vue3';
</script>

<template>
    <Head title="Language Test"/>

    <h1>{{ usePage().props.locale }}</h1>
    <Link :href="route('contact')">Go to Contact</Link>
</template>
