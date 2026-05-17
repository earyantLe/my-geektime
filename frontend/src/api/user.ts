import request from '@/utils/request'

export interface UserListParams {
  page?: number
  perPage?: number
  status?: number
}

export interface UserItem {
  uid: string
  user_name: string
  nick_name: string
  status: number
  created_at: string
  updated_at: string
}

export const getUserList = (params?: UserListParams) => {
  return request.get<any, { rows: UserItem[]; count: number }>('/user/list', { params })
}

export const updateUserStatus = (params: { uid: string; status: number }) => {
  return request.post('/user/status', params)
}

export const createUser = (params: { user_name: string; nick_name: string; password: string }) => {
  return request.post('/user/create', params)
}

export const deleteUser = (params: { uid: string }) => {
  return request.post('/user/delete', params)
}
