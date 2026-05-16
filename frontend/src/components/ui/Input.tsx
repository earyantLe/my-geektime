import React from 'react'
import { clsx } from 'clsx'

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
}

export const Input: React.FC<InputProps> = ({
  label,
  error,
  className,
  ...props
}) => {
  return (
    <div className="w-full">
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      <input
        className={clsx(
          'w-full h-10 px-3 text-sm border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-400 bg-white/80 backdrop-blur-sm transition-all',
          error ? 'border-red-400' : 'border-gray-200 hover:border-purple-300',
          className
        )}
        {...props}
      />
      {error && <p className="mt-1 text-sm text-red-500">{error}</p>}
    </div>
  )
}
