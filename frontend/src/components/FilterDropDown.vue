
<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useStockStore } from '@/stores/stocks'

const store = useStockStore()
const mostrar = ref(false)
const dropdown = ref<HTMLElement | null>(null)

function handleClickOutside(event: MouseEvent) {
  const target = event.target as HTMLElement
  const button = document.getElementById('filter-button')

  if (
    mostrar.value &&
    !dropdown.value?.contains(target) &&
    !button?.contains(target)
  ) {
    mostrar.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <section class="relative grid place-items-center">
    <div class="relative inline-block text-left">
      <div>
        <button
          type="button"
          class="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
          id="filter-button"
          aria-haspopup="true"
          @click="mostrar = !mostrar"
        >
          Columnas
          <svg class="-mr-1 h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path
              fill-rule="evenodd"
              d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>

      <div
        v-if="mostrar"
        ref="dropdown"
        class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
        role="menu"
        aria-orientation="vertical"
        aria-labelledby="filter-button"
      >
        <div class="py-1" role="none">
          <label
            v-for="col in store.columnas"
            :key="col.key"
            class="inline-flex items-center px-4 py-2 text-sm text-gray-700 w-full cursor-pointer"
          >
            <input
              type="checkbox"
              v-model="col.visible"
              class="form-checkbox h-5 w-5 text-gray-600"
            />
            <span class="ml-2">{{ col.label }}</span>
          </label>
        </div>
      </div>
    </div>
  </section>
</template>
