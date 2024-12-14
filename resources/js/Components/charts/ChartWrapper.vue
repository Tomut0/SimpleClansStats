<template>
    <div class="bg-darkside-900 p-6 rounded-2xl">
        <h3 class="text-gray-400 text-xl font-bold mb-4">
            <slot name="default"/>
        </h3>
        <div >
            <slot name="chart"/>
        </div>
    </div>
</template>
