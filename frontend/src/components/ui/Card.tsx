import React from 'react'
import { clsx } from 'clsx'

interface CardProps {
  children?: React.ReactNode
  className?: string
  header?: React.ReactNode
  actions?: React.ReactNode
}

export const Card: React.FC<CardProps> = ({ children, className, header, actions }) => {
  return (
    <div className={clsx('bg-white/90 backdrop-blur-xl rounded-2xl shadow-xl overflow-hidden', className)}>
      {header && (
        <div className="px-5 py-4 border-b border-white/20 flex justify-between items-center bg-gradient-to-r from-purple-500/10 to-pink-500/10">
          <div className="font-semibold text-gray-800">{header}</div>
          {actions && <div>{actions}</div>}
        </div>
      )}
      {children}
    </div>
  )
}
