import React, { useState, useEffect } from 'react'
import { Palette, Check } from 'lucide-react'
import { themes, getCurrentTheme, setTheme } from '@/utils/theme'

export const ThemeSwitcher: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [currentTheme, setCurrentTheme] = useState(getCurrentTheme())

  useEffect(() => {
    // 监听主题变更事件
    const handleThemeChange = (event: Event) => {
      const customEvent = event as CustomEvent<string>
      setCurrentTheme(customEvent.detail)
    }

    window.addEventListener('themeChange', handleThemeChange)
    return () => window.removeEventListener('themeChange', handleThemeChange)
  }, [])

  const handleThemeChange = (themeName: string) => {
    setTheme(themeName)
    setCurrentTheme(themeName)
    setIsOpen(false)
  }

  return (
    <div className="relative">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="p-2 hover:bg-white/20 rounded-lg transition-all duration-200 group"
        title="切换主题"
      >
        <Palette 
          size={18} 
          className="text-white/80 group-hover:text-white transition-colors" 
        />
      </button>

      {isOpen && (
        <>
          {/* 遮罩层 */}
          <div 
            className="fixed inset-0 z-40" 
            onClick={() => setIsOpen(false)}
          />
          
          {/* 主题选择面板 */}
          <div className="absolute right-0 top-12 z-50 w-48 bg-white/95 backdrop-blur-xl rounded-xl shadow-2xl border border-white/20 overflow-hidden">
            <div className="p-3 border-b border-gray-100">
              <h3 className="text-sm font-semibold text-gray-700">选择主题</h3>
            </div>
            
            <div className="p-2 space-y-1">
              {Object.values(themes).map((theme) => (
                <button
                  key={theme.name}
                  onClick={() => handleThemeChange(theme.name)}
                  className={`w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-all duration-200 ${
                    currentTheme === theme.name
                      ? 'bg-primary-50 text-primary-700'
                      : 'hover:bg-gray-50 text-gray-700'
                  }`}
                >
                  {/* 主题颜色预览 */}
                  <div className="flex gap-0.5">
                    <div 
                      className="w-4 h-4 rounded-full" 
                      style={{ backgroundColor: theme.primary[500] }}
                    />
                    <div 
                      className="w-4 h-4 rounded-full -ml-2 border-2 border-white" 
                      style={{ backgroundColor: theme.primary[600] }}
                    />
                  </div>
                  
                  <span className="flex-1 text-left text-sm font-medium">
                    {theme.label}
                  </span>
                  
                  {currentTheme === theme.name && (
                    <Check size={16} className="text-primary-500" />
                  )}
                </button>
              ))}
            </div>
          </div>
        </>
      )}
    </div>
  )
}
