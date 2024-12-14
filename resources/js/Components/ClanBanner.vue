<script setup>

import {computed, onMounted, ref} from "vue";
import {createBanner, isNotBanner} from "@/banner.js";

const emit = defineEmits(['created']);

const props = defineProps({
    clan: {
        type: Object,
        required: true
    },
    width: {
        type: Number,
        default: 20
    },
    height: {
        type: Number,
        default: 40
    }
});

const hasBannerMeta = computed(() => !isNotBanner(props.clan.banner));

const bannerDOM = ref();

onMounted(() => {
    if (bannerDOM.value) {
        const ctx = bannerDOM.value.getContext("2d");
        axios.get(route('v1:images.index'), {
            params: {
                category: "banner"
            }
        }).then((response) => {
            createBanner(props.clan.banner, response.data, ctx, props.width, props.height).then(() => {
                emit('created', bannerDOM);
            });
        }).catch(console.error);
    }
});
</script>

<template>
    <canvas ref="bannerDOM" v-if="hasBannerMeta" :width="width" :height="height" class="h-full banner"
            style="align-self: center"></canvas>
</template>

<style scoped>

</style>
