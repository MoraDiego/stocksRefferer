<script setup lang="ts">
import { computed } from 'vue'
import { useStockStore } from '@/stores/stocks'
import { storeToRefs } from 'pinia'

const store = useStockStore()
const { riesgo } = storeToRefs(store)
const riesgoLabel = computed(() => store.etiquetasRiesgo[riesgo.value])

const sliderColor = computed(() => {

  let r: number, g: number, b: number
    r = 255
    g = 0
    b = 0  
  if (riesgo.value <= 1) {
    // Verde (34, 197, 94) → Naranja (249, 115, 22)
    r = 0
    g = 255
    b = 0
  } else if(riesgo.value == 2){
    // Naranja (249, 115, 22) → Rojo (239, 68, 68)
    r = 255
    g = 125
    b = 0
  }
  return `rgb(${r}, ${g}, ${b})`
})

</script>

<template>
  <div class="flex flex-col items-center gap-2">
    <label for="riesgo" class="text-sm font-semibold text-gray-700">
      Nivel de Riesgo: <span :style="{ color: sliderColor }">{{ riesgoLabel }}</span>
    </label>
    <input
      id="riesgo"
      type="range"
      min="0"
      max="3"
      step="1"
      v-model="riesgo"
      class="w-64 h-3 rounded-lg appearance-none cursor-pointer"
      :style="{ backgroundColor: sliderColor}"
    />
  </div>
</template>