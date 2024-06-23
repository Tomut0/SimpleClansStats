<script setup>

import ApplicationLogo from "@/Components/ui/ApplicationLogo.vue";
import Translator from "@/Components/Translator.vue";
import SearchBar from "@/Components/SearchBar.vue";
import {__} from "@/trans.js";
import {ref} from "vue";

const show = ref(true);
</script>

<template>
    <header class="flex justify-between">
        <ApplicationLogo/>

        <nav class="flex items-center justify-center space-x-4">
            <!-- fixme: search bar is work in progress -->

            <!--            <SearchBar :expanded="false" icon-styles="text-white" search-styles="pl-10 bg-transparent border-0 border-b-2 rounded-none
                           focus:outline-none focus:ring-transparent focus:text-indigo-500 focus:border-indigo-500"
                       :placeholder="__('components.searchbar.clan.placeholder')"/>-->
            <Translator/>
        </nav>
    </header>
</template>

<style scoped>

</style>
