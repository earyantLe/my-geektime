import request from '@/utils/request'

export interface DictItem {
  id: string
  name: string
  key: string
  content: any[]
  sort: number
  summary?: string
  pkey?: string
  created: string
  updated: string
}

export const getDictList = (params?: { page?: number; perPage?: number; name?: string; key?: string }) => {
  return request.get<any, { rows: DictItem[]; count: number }>('/sys/dict/list', { params })
}

export const getDictTree = (key: string) => {
  return request.get<any, Record<string, any[]>>('/sys/dict/tree', { params: { key } })
}

export const createDict = (data: Partial<DictItem>) => {
  return request.post('/sys/dict/create', data)
}

export const updateDict = (data: Partial<DictItem>) => {
  return request.put('/sys/dict/update', data)
}

export const deleteDict = (id: string) => {
  return request.delete('/sys/dict/delete', { data: { id } })
}
