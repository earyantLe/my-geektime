import request from '@/utils/request'

export interface TaskListParams {
  page?: number
  perPage?: number
  direction?: number
  tag?: number
  product_type?: number
  product_form?: number
  xstatus?: number
  keywords?: string
}

export interface TaskItem {
  task_id: string
  task_name: string
  subtitle: string
  cover: string
  author: {
    name: string
    intro: string
  }
  is_audio?: boolean
  is_video?: boolean
  is_finish: boolean
  sale_type: number
  sale: number
  other_type: number
  other_group: number
  status: number
  article: {
    count: number
  }
  statistics: {
    items: Record<number, number>
  }
  doc?: string
  redirect?: string
  dir?: string
}

export const getTaskList = (params?: TaskListParams) => {
  return request.get<any, { rows: TaskItem[]; count: number }>('/task/list', { params })
}

export const deleteTask = (ids: string[]) => {
  return request.delete('/task/delete', { data: { ids } })
}

export const retryTask = (params: { pid?: string; ids?: string[]; retry?: boolean }) => {
  return request.post('/task/retry', params)
}

export const exportTask = (params: { pid: string; type: string }) => {
  if (params.type === 'markdown') {
    return request.get<any, Blob>('/task/export', { params, responseType: 'blob' })
  }
  return request.get('/task/export', { params })
}

export interface TaskInfo {
  id: string
  task_name: string
  subtitle: string
  cover: string
  intro_html?: string
  author?: {
    name: string
    intro: string
  }
  article?: {
    count: number
  }
  is_audio?: boolean
  is_video?: boolean
  is_finish: boolean
  redirect?: string
}

export const getTaskInfo = (id: string) => {
  return request.get<any, TaskInfo>('/task/info', { params: { id } })
}

export const getArticleComments = (params: { aid: string; page?: number; perPage?: number }) => {
  return request.get('/task/article/comments', { params })
}

export const getCommentDiscussions = (params: {
  target_id: string
  target_type: number
  page?: number
  perPage?: number
  use_likes_order?: boolean
}) => {
  return request.get('/task/article/discussions', { params })
}
