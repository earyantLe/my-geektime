import React from 'react'
import { clsx } from 'clsx'

interface SwitchProps {
  checked: boolean
  onChange: (checked: boolean) => void
  label?: string
  onText?: string
  offText?: string
  disabled?: boolean
}

export const Switch: React.FC<SwitchProps> = ({
  checked,
  onChange,
  label,
  onText = '开启',
  offText = '关闭',
  disabled = false,
}) => {
  return (
    <label className="flex items-center gap-2 cursor-pointer">
      {label && <span className="text-sm text-gray-700">{label}</span>}
      <div className="relative">
        <input
          type="checkbox"
          checked={checked}
          onChange={(e) => onChange(e.target.checked)}
          disabled={disabled}
          className="sr-only"
        />
        <div
          className={clsx(
            'w-12 h-6 rounded-full transition-colors',
            checked ? 'bg-primary-500' : 'bg-gray-300',
            disabled && 'opacity-50 cursor-not-allowed'
          )}
        />
        <div
          className={clsx(
            'absolute top-1 w-4 h-4 bg-white rounded-full transition-transform',
            checked ? 'translate-x-7' : 'translate-x-1'
          )}
        />
      </div>
      <span className="text-sm text-gray-600">{checked ? onText : offText}</span>
    </label>
  )
}
