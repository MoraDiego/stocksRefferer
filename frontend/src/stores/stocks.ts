import { ref, computed, watch } from 'vue'
import { defineStore } from 'pinia'
import type { Stock } from '@/types/stocks'
import { getAcciones, getAccionesPorRiesgo } from '@/requests/requests'
export const useStockStore = defineStore('stocks', () => {
  const riesgo = ref(0)
  const acciones = ref<Stock[]>([])
  const etiquetasRiesgo: Record<number, string> = {
    0: 'Muy Bajo',
    1: 'Bajo',
    2: 'Moderado',
    3: 'Alto'
  }

  const columnas = ref([
    { key: 'ticker', label: 'Ticker', visible: true },
    { key: 'company', label: 'Empresa', visible: true },
    { key: 'action', label: 'Acción', visible: true },
    { key: 'TargetFrom', label: 'Desde', visible: true },
    { key: 'TargetTo', label: 'Hacia', visible: true },
    { key: 'rating', label: 'Calificación', visible: true },
    { key: 'brokerage', label: 'Broker', visible: true },
    { key: 'time', label: 'Fecha', visible: true }
  ])
  const columnasVisibles = computed(() =>
    columnas.value.filter(col => col.visible)
  )
  async function cargarAcciones() {
    try {
      const datos = await getAcciones()
      acciones.value = datos
    } catch (error) {
      console.error('Error al obtener acciones:', error)
    }
  }
  async function cargarAccionesRiesgo() {
    try {
      const datos = await getAccionesPorRiesgo(riesgo.value)
      acciones.value = datos
    } catch (error) {
      console.error('Error al obtener acciones:', error)
    }
  }
  watch(riesgo, () => {
    cargarAccionesRiesgo()
  }, { immediate: true })

  return {columnas, columnasVisibles, riesgo, etiquetasRiesgo, acciones, cargarAcciones}
})
