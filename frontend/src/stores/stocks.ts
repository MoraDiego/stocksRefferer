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

  const Columns = [
  { key: 'ticker', label: 'Ticker' },
  { key: 'company', label: 'Empresa' },
  { key: 'action', label: 'Acción' },
  { key: 'TargetFrom', label: 'Desde' },
  { key: 'TargetTo', label: 'Hacia' },
  { key: 'rating', label: 'Calificación' },
  { key: 'brokerage', label: 'Broker' },
  { key: 'time', label: 'Fecha' },
] as const

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

  return {Columns, riesgo, etiquetasRiesgo, acciones, cargarAcciones}
})
