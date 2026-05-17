import React from 'react'
import { clsx } from 'clsx'

interface SpinnerProps {
  size?: 'sm' | 'md' | 'lg'
  className?: string
}

export const Spinner: React.FC<SpinnerProps> = ({ size = 'md', className }) => {
  const sizes = {
    sm: 'w-4 h-4',
    md: 'w-8 h-8',
    lg: 'w-12 h-12',
  }

  return (
    <div
      className={clsx(
        'relative animate-spin',
        sizes[size],
        className
      )}
      style={{
        animationDuration: '0.8s',
      }}
    >
      {/* 外层圆环 */}
      <div className="absolute inset-0 rounded-full border-2 border-gray-200" />
      {/* 菊花旋转部分 - 使用主题色 */}
      <div 
        className="absolute inset-0 rounded-full" 
        style={{
          background: 'conic-gradient(from 0deg, transparent 0%, transparent 70%, var(--color-primary-500, #3b82f6) 70%, var(--color-primary-600, #2563eb) 100%)',
        }}
      />
      {/* 内部遮罩，创建圆环效果 */}
      <div className="absolute inset-[2px] rounded-full bg-white" />
    </div>
  )
}
