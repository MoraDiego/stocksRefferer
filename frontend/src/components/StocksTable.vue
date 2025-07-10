<script setup lang="ts">
import { useStockStore } from '@/stores/stocks'
import { storeToRefs } from 'pinia'
const store = useStockStore()
const columns  = store.Columns.map(col => col)
store.cargarAcciones()
const {acciones} = storeToRefs(store)
</script>
<template>
    <div class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200">
        <thead>
          <tr class="bg-green-100 text-left text-sm font-medium text-green-800 uppercase">
            <th
              v-for="col in columns"
              class="px-4 py-2 border"
            >
              {{ col.label }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="stock in acciones"
            class="hover:bg-green-50"
          >
            <td v-for="col in columns" :key="col.key" class="px-4 py-2 border">
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
  </div>
</template>