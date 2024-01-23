<script setup>
import {LanguageIcon} from "@heroicons/vue/24/solid/index.js";
import SearchBar from "@/Components/SearchBar.vue";
import {__} from "@/trans.js";
import Popup from "@/Components/Popup.vue";
import {ref, watch} from "vue";
import Flag from "@/Components/Flag.vue";
import {Link, usePage} from "@inertiajs/vue3";

const showPopup = ref(false);
let flags = ref(), sortedFlags = ref(new Proxy({}, {}));

// Lazy loading languages
watch(showPopup, async (value) => {
    if (value && !flags.value) {
        const localesPromise = await axios.get(route("locales.status"));
        flags.value = localesPromise.data;
        sortedFlags.value = flags.value;
    }
});

// Needed to get language display name
const translations = JSON.parse(usePage().props.translations);

function searchLanguage(value) {
    if (!value) {
        sortedFlags.value = flags.value;
        return;
    }

    let founded = [];
    for (const lang in translations) {
        if (translations[lang].displayName.toLowerCase().includes(value.toLowerCase())) {
            founded.push(lang);
        }
    }

    if (founded.length > 0) {
        sortedFlags.value = founded.reduce((acc, key) => {
            if (!acc[key]) {
                acc[key] = flags.value[key];
            }
            return acc;
        }, {});
    } else {
        sortedFlags.value = null;
    }
}

</script>

<template>
    <div class="relative">
        <LanguageIcon class="w-6 h-6 ease-in transition text-white cursor-pointer hover:text-blue-700"
                      :class="{ '!text-blue-700':showPopup }"
                      v-on:click="showPopup = !showPopup"/>

        <Popup :seen="showPopup" class="absolute top-8 right-0 bg-white rounded w-64 min-w-full">
            <SearchBar expanded disabled @search="searchLanguage" class="border-2 border-gray-300 rounded-md" icon-styles="ml-2 !w-5 text-black"
                       search-styles="w-full border-none focus:border-none focus:ring-0 !pl-1"
                       :placeholder="__('components.searchbar.locale.placeholder')" />

            <div class="flex flex-col" v-if="sortedFlags">
                <TransitionGroup name="fade">
                    <Link class="p-2 hover:bg-gray-100" v-for="(percent, key) in sortedFlags" :key="key"
                          :href="route('locale.update')"
                          @click="showPopup = false"
                          :data="{'locale': key}" as="button" type="submit" method="post">

                        <Flag :languageTag="key" :display-name="translations[key].displayName"
                              :translation-status="percent"/>
                    </Link>
                </TransitionGroup>
            </div>
            <div class="mt-2" v-else>
                <h4>{{ __('components.translator.locale_not_found') }}</h4>
            </div>
        </Popup>
    </div>
</template>

<style scoped>
.fade-enter-active {
    transition: all 0.3s ease-out;
}

.fade-leave-active {
    transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.fade-enter-from,
.fade-leave-to {
    transform: translateX(20px);
    opacity: 0;
}
</style>
