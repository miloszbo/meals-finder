<template>
  <!-- Container for the star rating -->
  <div class="flex items-center space-x-1">
    <!-- Loop over 5 stars -->
    <i
      v-for="star in 5"
      :key="star"
      class="fa fa-star fa-lg cursor-pointer transition-colors"
      :class="{
        /* Highlight in yellow if hovered or selected */
        'text-yellow-400': (hoverRating || internalRating) >= star,
        /* Otherwise show as gray */
        'text-gray-400': (hoverRating || internalRating) < star
      }"
      @mouseover="hoverRating = star" 
      @mouseleave="hoverRating = 0"
      @click="updateRating(star)"

    />
     <!-- 
     @mouseover="hoverRating = n"    ← on hover, set hoverRating 
     @mouseleave="hoverRating = 0"   ← clear hover state 
     @click="updateRating(n)"        ← commit this star as the rating 
    -->
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// v-model prop: initial rating value from parent
const props = defineProps({
  modelValue: {
    type: Number,
    default: 0
  }
})

// enable emitting update:modelValue to support v-model
const emit = defineEmits(['update:modelValue'])

// reactive state: the current rating and hover index
const internalRating = ref(props.modelValue)
const hoverRating = ref(0)

// sync when parent changes
watch(() => props.modelValue, val => {
  internalRating.value = val
})

function updateRating(value) {
  internalRating.value = value
  emit('update:modelValue', value)
}
</script>
