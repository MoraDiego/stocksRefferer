<script setup lang="ts">
import { useStockStore } from '@/stores/stocks'
const store = useStockStore()
const columns  = store.Columns.map(col => col.key)
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
              {{col}}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            class="hover:bg-green-50"
          >
            <td v-for="col in columns" class="px-4 py-2 border">
              <span>
                {{col}}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
  </div>
</template>