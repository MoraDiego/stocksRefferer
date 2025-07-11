import type { Stock } from '@/types/stocks'

const BASE_URL = 'http://localhost:8080'

export async function getAcciones(): Promise<Stock[]> {
  const res = await fetch(`${BASE_URL}/acciones`)
  if (!res.ok) throw new Error(`Error ${res.status}: ${res.statusText}`)
  return res.json()
}
export async function getAccionesPorRiesgo(riesgo: number) {
  const res = await fetch(`http://localhost:8080/accionesRiesgo?risk=${riesgo}`)
  if (!res.ok) throw new Error('Error al obtener acciones')
  return await res.json()
}