<script setup>
import {computed, ref} from "vue";
import {ChevronLeftIcon, ChevronRightIcon} from "@heroicons/vue/20/solid/index.js";

const currentIndex = ref(0);

const props = defineProps({
    carouselClass: {
        type: String,
        default: ''
    },
    items: {
        type: Array,
        required: true
    },
    showCount: {
        type: Number,
        default: 4
    }
});

const prev = () => {
    if (canPrev.value) {
        currentIndex.value--;
    }
};

const next = () => {
    if (canNext.value) {
        currentIndex.value++;
    }
};

const canPrev = computed(() => {
    return currentIndex.value > 0;
});

const canNext = computed(() => {
    return currentIndex.value < props.items.length - props.showCount;
});

const visibleItems = computed(() => {
    return props.items.slice(currentIndex.value, currentIndex.value + props.showCount);
});
</script>

<template>
    <div class="flex gap-2 items-center select-none lg:gap-4">
        <button class="rounded-full bg-darkside-700 transition-colors duration-300 ease-in-out" @click.prevent.stop="prev" :class="{'bg-darkside-500': canPrev}">
            <ChevronLeftIcon class="w-8 h-8 text-white"/>
        </button>
        <TransitionGroup :class="carouselClass" tag="div" name="slide">
            <component v-for="(item, index) in visibleItems" :is="item.component" v-bind="item.props" :key="currentIndex + index + showCount"/>
        </TransitionGroup>
        <button class="rounded-full bg-darkside-700 transition-colors duration-300 ease-in-out" @click.prevent.stop="next" :class="{'bg-darkside-500': canNext}">
            <ChevronRightIcon class="w-8 h-8 text-white"/>
        </button>
    </div>
</template>

<style scoped>
.slide-move, .slide-enter-active, .slide-leave-active {
    transition: all 0.3s cubic-bezier(.17, .67, .38, 1.13);
}

.slide-enter-from, .slide-leave-to {
    opacity: 0;
    transform: scale(0.3);
}

.slide-leave-active {
    position: absolute;
}
</style>
