import React from 'react'
import { clsx } from 'clsx'
import { AlertCircle, CheckCircle, Info, XCircle, X } from 'lucide-react'

interface AlertProps {
  children: React.ReactNode
  variant?: 'info' | 'success' | 'warning' | 'error'
  showIcon?: boolean
  closable?: boolean
  onClose?: () => void
  className?: string
}

export const Alert: React.FC<AlertProps> = ({
  children,
  variant = 'info',
  showIcon = true,
  closable = false,
  onClose,
  className,
}) => {
  const variants = {
    info: 'bg-gradient-to-r from-blue-500/10 to-cyan-500/10 text-blue-800 border-blue-200/50',
    success: 'bg-gradient-to-r from-green-500/10 to-emerald-500/10 text-green-800 border-green-200/50',
    warning: 'bg-gradient-to-r from-yellow-500/10 to-orange-500/10 text-yellow-800 border-yellow-200/50',
    error: 'bg-gradient-to-r from-red-500/10 to-primary-500/10 text-red-800 border-red-200/50',
  }

  const icons = {
    info: Info,
    success: CheckCircle,
    warning: AlertCircle,
    error: XCircle,
  }

  const iconColors = {
    info: 'text-blue-500',
    success: 'text-green-500',
    warning: 'text-yellow-500',
    error: 'text-red-500',
  }

  const Icon = icons[variant]

  return (
    <div
      className={clsx(
        'flex items-start gap-3 p-4 border rounded-xl backdrop-blur-sm shadow-lg',
        variants[variant],
        className
      )}
    >
      {showIcon && <Icon size={20} className={clsx('flex-shrink-0 mt-0.5', iconColors[variant])} />}
      <div className="flex-1">{children}</div>
      {closable && (
        <button
          onClick={onClose}
          className="flex-shrink-0 hover:opacity-70 transition-opacity"
        >
          <X size={18} />
        </button>
      )}
    </div>
  )
}
