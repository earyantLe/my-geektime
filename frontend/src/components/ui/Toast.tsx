import React, { createContext, useContext, useState, useCallback, useEffect } from 'react'
import { clsx } from 'clsx'
import { CheckCircle, XCircle, AlertCircle, Info, X } from 'lucide-react'
import { showErrorToastEvent } from '@/utils/request'

export interface Toast {
  id: string
  message: string
  type: 'success' | 'error' | 'warning' | 'info'
  duration?: number
}

interface ToastContextType {
  toasts: Toast[]
  addToast: (message: string, type?: Toast['type'], duration?: number) => void
  removeToast: (id: string) => void
}

const ToastContext = createContext<ToastContextType | undefined>(undefined)

const ToastItem: React.FC<{ toast: Toast; onRemove: () => void }> = ({ toast, onRemove }) => {
  const icons = {
    success: CheckCircle,
    error: XCircle,
    warning: AlertCircle,
    info: Info,
  }

  const variants = {
    success: 'bg-gradient-to-r from-green-500/20 to-emerald-500/20 border-green-400/50 text-green-100',
    error: 'bg-gradient-to-r from-red-500/20 to-pink-500/20 border-red-400/50 text-red-100',
    warning: 'bg-gradient-to-r from-yellow-500/20 to-orange-500/20 border-yellow-400/50 text-yellow-100',
    info: 'bg-gradient-to-r from-blue-500/20 to-cyan-500/20 border-blue-400/50 text-blue-100',
  }

  const iconColors = {
    success: 'text-green-400',
    error: 'text-red-400',
    warning: 'text-yellow-400',
    info: 'text-blue-400',
  }

  const Icon = icons[toast.type]

  React.useEffect(() => {
    const timer = setTimeout(() => {
      onRemove()
    }, toast.duration || 3000)
    return () => clearTimeout(timer)
  }, [toast.duration, onRemove])

  return (
    <div
      className={clsx(
        'flex items-center gap-3 px-4 py-3 rounded-xl border backdrop-blur-md shadow-lg',
        'animate-slide-in min-w-[280px] max-w-[400px]',
        variants[toast.type]
      )}
    >
      <Icon size={20} className={clsx('flex-shrink-0', iconColors[toast.type])} />
      <span className="flex-1 text-sm font-medium">{toast.message}</span>
      <button
        onClick={onRemove}
        className="flex-shrink-0 hover:opacity-70 transition-opacity"
      >
        <X size={16} />
      </button>
    </div>
  )
}

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [toasts, setToasts] = useState<Toast[]>([])

  const addToast = useCallback((message: string, type: Toast['type'] = 'info', duration = 3000) => {
    const id = Math.random().toString(36).substring(2, 9)
    setToasts((prev) => [...prev, { id, message, type, duration }])
  }, [])

  const removeToast = useCallback((id: string) => {
    setToasts((prev) => prev.filter((t) => t.id !== id))
  }, [])

  // 监听 API 错误事件，自动显示 Toast
  useEffect(() => {
    const handler = (event: Event) => {
      const customEvent = event as CustomEvent<string>
      addToast(customEvent.detail, 'error', 5000)
    }
    showErrorToastEvent.addEventListener('showError', handler)
    return () => {
      showErrorToastEvent.removeEventListener('showError', handler)
    }
  }, [addToast])

  return (
    <ToastContext.Provider value={{ toasts, addToast, removeToast }}>
      {children}
      <div className="fixed top-4 right-4 z-[9999] flex flex-col gap-2">
        {toasts.map((toast) => (
          <ToastItem key={toast.id} toast={toast} onRemove={() => removeToast(toast.id)} />
        ))}
      </div>
    </ToastContext.Provider>
  )
}

export const useToast = () => {
  const context = useContext(ToastContext)
  if (!context) {
    throw new Error('useToast must be used within a ToastProvider')
  }
  return context
}
