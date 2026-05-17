import React from 'react'
import { Menu } from 'lucide-react'
import { ThemeSwitcher } from '@/components/ThemeSwitcher'

interface HeaderProps {
  onMenuClick?: () => void
}

export const Header: React.FC<HeaderProps> = ({ onMenuClick }) => {
  return (
    <header className="h-14 bg-white/80 backdrop-blur-xl border-b border-white/20 flex items-center px-4 lg:hidden">
      <button
        onClick={onMenuClick}
        className="p-2 hover:bg-white/50 rounded-lg transition-colors"
      >
        <Menu size={24} className="text-gray-600" />
      </button>
      <h2 className="ml-3 text-lg font-semibold bg-gradient-to-r from-primary-600 to-primary-500 bg-clip-text text-transparent">
        管理后台
      </h2>
      <div className="ml-auto">
        <ThemeSwitcher />
      </div>
    </header>
  )
}
