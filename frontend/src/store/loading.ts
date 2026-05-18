import { create } from 'zustand'

interface LoadingState {
  isLoading: boolean
  showLoading: () => void
  hideLoading: () => void
}

// 使用简单的计数器变量而不是 state
let requestCount = 0

export const useLoadingStore = create<LoadingState>((set) => ({
  isLoading: false,
  showLoading: () => {
    requestCount++
    set({ isLoading: true })
  },
  hideLoading: () => {
    requestCount = Math.max(0, requestCount - 1)
    if (requestCount === 0) {
      set({ isLoading: false })
    }
  },
}))
