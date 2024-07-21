<script setup>

const props = defineProps({
    translationStatus: String,
    languageTag: {
        type: String,
        required: true
    },
    displayName: {
        type: String,
        required: true
    },
});

const tag = props.languageTag.split("_"),
    locale = tag[0],
    country = locale === "en" ? `us` : tag[1].toLowerCase();
</script>

<template>
    <div class="flex justify-between items-center">
        <div class="inline-flex items-center">
            <div :class="`fib fi-${country} flex-none`" class="w-12 h-12"></div>
            <div class="mx-4 shrink text-start leading-tight">{{ displayName }}</div>
        </div>
        <!-- fixme: some translation lines shouldn't be translated and current status doesn't consider it -->
        <!-- <div>{{ translationStatus }}</div>-->
    </div>
</template>

<style scoped>

</style>
