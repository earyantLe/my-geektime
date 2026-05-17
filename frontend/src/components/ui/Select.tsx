import React from 'react'
import { clsx } from 'clsx'

interface Option {
  label: string
  value: string | number
}

interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {
  label?: string
  error?: string
  options: Option[]
}

export const Select: React.FC<SelectProps> = ({
  label,
  error,
  options,
  className,
  ...props
}) => {
  return (
    <div>
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      <select
        className={clsx(
          'w-full h-10 px-3 text-sm border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-400 bg-white/80 backdrop-blur-sm transition-all',
          error ? 'border-red-400' : 'border-gray-200 hover:border-purple-300',
          className
        )}
        {...props}
      >
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
      {error && <p className="mt-1 text-sm text-red-500">{error}</p>}
    </div>
  )
}
