<script setup lang="ts">
import { useStockStore } from '@/stores/stocks'
import { storeToRefs } from 'pinia'
import { computed, ref, watch } from 'vue'
const store = useStockStore()
const {acciones} = storeToRefs(store)
//Paginacion de la tabla
const currentPage = ref(1)
const itemsPerPage = 5
const totalPages = computed(() =>
  Math.ceil(acciones.value.length / itemsPerPage)
)
const paginatedStocks = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  return acciones.value.slice(start, start + itemsPerPage)
})
function nextPage() {
  if (currentPage.value < totalPages.value) currentPage.value++
}
function prevPage() {
  if (currentPage.value > 1) currentPage.value--
}
watch(() => store.acciones.length, () => {
  if (currentPage.value > totalPages.value) {
    currentPage.value = totalPages.value || 1
  }
})
</script>
<template>
    <div class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200">
        <thead>
          <tr class="bg-green-100 text-left text-sm font-medium text-green-800 uppercase">
            <th
              v-for="col in store.columnasVisibles"
              class="px-4 py-2 border"
            >
              {{ col.label }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="stock in paginatedStocks"
            class="hover:bg-green-50"
          >
            <td v-for="col in store.columnasVisibles" :key="col.key" class="px-4 py-2 border">
              <span v-if="col.key === 'rating'">
                {{ stock.rating_from }} → {{ stock.rating_to }}
              </span>
              <span v-else>
                {{ stock[col.key as keyof typeof stock] }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="flex justify-between items-center text-sm">
        <button
          class="bg-gray-200 px-3 py-1 rounded disabled:opacity-50"
          @click="prevPage"
          :disabled="currentPage === 1"
        >
          Anterior
        </button>

        <span>Página {{ currentPage }} de {{ totalPages }}</span>

        <button
          class="bg-gray-200 px-3 py-1 rounded disabled:opacity-50"
          @click="nextPage"
          :disabled="currentPage === totalPages"
        >
          Siguiente
        </button>
      </div>
  </div>
</template>