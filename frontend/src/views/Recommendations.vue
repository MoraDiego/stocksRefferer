<script setup lang="ts">
import FilterDropDown from '@/components/FilterDropDown.vue';
import AppHeader from '@/components/Header.vue'
import Slider from '@/components/Slider.vue'
import StocksTable from '@/components/StocksTable.vue'
</script>
<template>
  <div class="min-h-screen bg-gray-100">
    <AppHeader />
    <main class="max-w-6xl mx-auto bg-white shadow-md rounded-lg p-6 mt-6">
      <h2 class="text-3xl font-bold text-green-700 mb-4">
        Nuestras acciones recomendadas
      </h2>
      <p class="text-lg mb-6">
        Aquí podra seleccionar un nivel de riesgo que este dispuesto a asumir para 
        invertir en acciones, y con el, le daremos las recomendaciones que mas se ajusten
        a su perfil
      </p>
      <div class="relative inline-flex w-full">
        <div class="p-6">
          <FilterDropDown />
        </div>
        <div class="p-6 absolute left-1/2 -translate-x-1/2">
          <Slider />
        </div>
      </div>
      <StocksTable />
    </main>
  </div>
</template>