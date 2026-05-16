import request from '@/utils/request'

export interface LoginParams {
  type: string
  data: {
    account: string
    password: string
  } | {
    code: string
  }
}

export interface RegisterParams {
  type: string
  data: {
    account: string
    password: string
  } | {
    email: string
    code: string
    password: string
  }
}

export interface LoginResponse {
  status: number
  token: string
  user: {
    uid: string
    user_name: string
    avatar: string
    role_id: number
    geek_auth?: boolean
  }
}

export const login = (params: LoginParams) => {
  return request.post<any, LoginResponse>('/base/login', params)
}

export const register = (params: RegisterParams) => {
  return request.post('/base/register', params)
}

export interface ConfigData {
  register_types: string[]
  login_types: string[]
  login_guest: {
    type: string
    account: string
    password: string
  }[]
  email_login: {
    enabled: boolean
    verify_code: boolean
  }
  github_login: {
    enabled: boolean
    client_id?: string
    auth_url?: string
    redirect_url?: string
    scope?: string
  }
}

export const getConfig = () => {
  return request.get<any, ConfigData>('/base/config')
}

export const refreshCookie = (cookie: string) => {
  return request.post('/base/refresh/cookie', { cookie })
}

export const sendEmailCode = (email: string) => {
  return request.post('/base/email/code', { email })
}
