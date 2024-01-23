<script setup>
/**
 * The ExpandedButton.vue component is used to create buttons
 * that can expand or collapse additional content.
 */
import {ref} from "vue";

const props = defineProps({
    is: {
        default: "Button"
    },
    afterExpand: {
        type: Function,
        default: function () {
        }
    },
    /**
     * Specifies whether the expandable content is positioned to the left or right of the button.
     * @values left, right
     */
    direction: {
        type: String,
        default: "right"
    },
    /**
     *  A boolean that controls whether the content is always shown or can be toggled.
     *  if true, the content will always be visible.
     *  @values false, true
     */
    isLocked: {
        type: Boolean,
        default: false
    },
    /**
     * A boolean that disables the toggle functionality of the button if true.
     * @values false, true
     */
    disabled: {
        type: Boolean,
        default: false,
    },
    /**
     *  A boolean that controls the positioning of the icon slot.
     *  If true and the content is expanded, the icon will be positioned absolutely.
     *  @values false, true
     */
    isAbsoluteIcon: Boolean
});

const shown = ref(false);

function toggle() {
    if (props.disabled) {
        return;
    }

    shown.value = !shown.value;

    if (shown.value) props.afterExpand(shown);
}
</script>

<template>
    <component :is="is">
        <span class="flex items-center" v-if="direction === 'left'">
            <span v-show="shown || isLocked"><slot></slot></span>
            <span @click.prevent.stop="toggle" :class="{'absolute':isAbsoluteIcon && shown}"><slot
                name="icon"></slot></span>
        </span>
        <span class="flex items-center" v-else-if="direction === 'right'">
            <span @click.prevent.stop="toggle" :class="{'absolute':isAbsoluteIcon && shown}"><slot
                name="icon"></slot></span>
            <span v-show="shown || isLocked"><slot></slot></span>
        </span>
    </component>
</template>

<style scoped>

</style>
