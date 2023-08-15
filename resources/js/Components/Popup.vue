<script setup>
defineProps({
    seen: {
        type: Boolean,
        default: true
    },
});

</script>

<template>
    <div v-show="seen" class="py-2 px-4 shadow-md">
        <slot></slot>
    </div>
</template>

