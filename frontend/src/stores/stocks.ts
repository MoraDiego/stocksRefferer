import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useStockStore = defineStore('stocks', () => {
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
  return {Columns}
})
