<script setup>

import {addColors} from "@/colors.js";
import {formatNumber, visitViaQuery} from "@/helpers.js";
import Modal from "@/Components/ui/Modal.vue";
import {CalendarDaysIcon, CheckIcon, CircleStackIcon, FireIcon, XMarkIcon} from "@heroicons/vue/24/solid/index.js";
import {usePage} from "@inertiajs/vue3";
import Selector from "@/Components/Selector.vue";
import {__} from "@/trans.js";
import ClanBanner from "@/Components/ClanBanner.vue";

const {locale, selectors} = usePage().props;
const currentClan = defineModel({required: true});

const clanEntities = selectors.current?.clanEntitySelector?.entity;
</script>

<template>
    <Modal :show="currentClan != null" closeable>
        <div class="py-6 px-8">
            <div class="flex">
                <ClanBanner :clan="currentClan" class="mr-4" />
                <div class="w-full">
                    <div class="flex items-center justify-between">
                        <div class="flex items-center">
                            <span class="text-2xl text-white" v-html="addColors(currentClan.color_tag)"></span>
                            <CheckIcon v-if="currentClan.verified" class="ml-2 w-4 h-4 text-darkside-400"/>
                        </div>

                        <XMarkIcon class="h-6 w-6 cursor-pointer text-darkside-400"
                                   @click="visitViaQuery(route('leaderboard.index'), [{name: 'clanTag', value: null}], ['queryClan'])"/>
                    </div>

                    <div class="text-gray-400 text-xl">{{ currentClan.name }}</div>
                </div>
            </div>
            <div class="flex my-2 gap-2 text-gray-300 font-semibold">
                <div class="flex items-center">
                    <CalendarDaysIcon class="w-5 h-5 text-darkside-400 mr-1"/>
                    {{ currentClan.formatted_founded }}
                </div>
                <div class="flex items-center">
                    <CircleStackIcon class="w-5 h-5 text-darkside-400 mr-1"/>
                    {{ formatNumber(currentClan.balance, locale, "currency") }}
                </div>
                <div class="flex items-center">
                    <FireIcon class="w-5 h-5 text-darkside-400 mr-1"/>
                    {{ formatNumber(currentClan.kdr, locale, "decimal", 0, 2) }}
                </div>
            </div>
            <div class="text-gray-300" v-if="currentClan.description">{{ currentClan.description }}</div>

            <Selector :items="selectors.all.clanEntitySelector" selectQuery="clanEntity" class="mt-2 text-center lg:text-left"
                      :additional-query="[{name: 'clanTag', value: currentClan.tag}]"
                      :only="['queryClan', 'selectors']"/>

            <div class="grid grid-cols-1 lg:grid-cols-4 gap-4 mt-4 text-white" v-if="clanEntities && clanEntities.length > 0">
                <div v-for="entity of clanEntities" class="flex items-center">
                    <img v-if="entity.image_src" :src="entity.image_src" :alt="`${entity.name}'s image`" class="h-8 w-8"/>
                    <div v-else class="h-8 w-8 rounded bg-darkside-700"></div>
                    <div class="ml-2">
                        <div class="leading-3">{{ entity.name }}</div>
                        <span class="font-light">{{ entity.value }}</span>
                    </div>
                </div>
            </div>
            <div v-else class="text-white mt-4">{{
                    __('components.clanmodal.entity_not_found', {entity: __(selectors.current.clanEntitySelector.translation)})
                }}
            </div>
        </div>
    </Modal>
</template>

<style scoped>
:deep(.selector-button) {
    background: transparent;
    border-radius: 0 !important;
}

:deep(.selector-button-selected) {
    @apply border-b-2 border-darkside-400 font-semibold
}

:deep(button:not(.selector-button-selected)) {
    @apply border-b-2 border-darkside-500
}

</style>
