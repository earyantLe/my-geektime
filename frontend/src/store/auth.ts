import { create } from 'zustand'
import { storage } from '@/utils/storage'

interface User {
  uid: string
  user_name: string
  avatar: string
  role_id: number
}

interface AuthState {
  token: string | null
  user: User | null
  geekAuth: boolean
  isAuthenticated: boolean
  setAuth: (token: string, user: User) => void
  setGeekAuth: (auth: boolean) => void
  logout: () => void
  init: () => void
}

export const useAuthStore = create<AuthState>((set) => ({
  token: null,
  user: null,
  geekAuth: false,
  isAuthenticated: false,
  
  setAuth: (token, user) => {
    storage.setToken(token)
    storage.setUser(user)
    set({ token, user, isAuthenticated: true })
  },
  
  setGeekAuth: (auth) => {
    if (auth) {
      storage.setGeekAuth('true')
    } else {
      storage.removeGeekAuth()
    }
    set({ geekAuth: auth })
  },
  
  logout: () => {
    storage.clear()
    set({ token: null, user: null, geekAuth: false, isAuthenticated: false })
  },
  
  init: () => {
    const token = storage.getToken()
    const userData = storage.getUser()
    const geekAuth = storage.getGeekAuth()
    
    if (token && userData.uid) {
      set({
        token,
        user: {
          uid: userData.uid!,
          user_name: userData.uname!,
          avatar: userData.avatar!,
          role_id: Number(userData.roleId),
        },
        geekAuth: !!geekAuth,
        isAuthenticated: true,
      })
    }
  },
}))
