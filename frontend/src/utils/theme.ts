// 主题定义
export interface Theme {
  name: string
  label: string
  description?: string
  primary: {
    50: string
    100: string
    200: string
    300: string
    400: string
    500: string
    600: string
    700: string
    800: string
    900: string
  }
}

// 预定义主题 - 基于 UI 设计最佳实践 2024/2025
// 遵循 60-30-10 法则，色彩心理学，WCAG 对比度标准
export const themes: Record<string, Theme> = {
  default: {
    name: 'default',
    label: '科技蓝',
    description: '专业、安全、可靠 - 适合科技类产品',
    primary: {
      // 蓝色系：传递信任、专业感（金融、科技类应用）
      50: '#eff6ff',
      100: '#dbeafe',
      200: '#bfdbfe',
      300: '#93c5fd',
      400: '#60a5fa',
      500: '#3b82f6', // 主色 - 60%
      600: '#2563eb',
      700: '#1d4ed8',
      800: '#1e40af',
      900: '#1e3a8a', // 深色 - 满足 WCAG 对比度 4.5:1
    },
  },
  purple: {
    name: 'purple',
    label: '优雅紫',
    description: '创意、奢华、神秘 - 适合设计类应用',
    primary: {
      // 紫色系：创意、奢华、高端感（设计、奢侈品应用）
      50: '#faf5ff',
      100: '#f3e8ff',
      200: '#e9d5ff',
      300: '#d8b4fe',
      400: '#c084fc',
      500: '#a855f7', // 主色 - 60%
      600: '#9333ea',
      700: '#7e22ce',
      800: '#6b21a8',
      900: '#581c87', // 深色 - 满足 WCAG 对比度 4.5:1
    },
  },
  green: {
    name: 'green',
    label: '清新绿',
    description: '自然、健康、成长 - 适合健康、环保类应用',
    primary: {
      // 绿色系：自然、健康、生机（健康、农业、环保产品）
      50: '#f0fdf4',
      100: '#dcfce7',
      200: '#bbf7d0',
      300: '#86efac',
      400: '#4ade80',
      500: '#22c55e', // 主色 - 60%
      600: '#16a34a',
      700: '#15803d',
      800: '#166534',
      900: '#14532d', // 深色 - 满足 WCAG 对比度 4.5:1
    },
  },
  rose: {
    name: 'rose',
    label: '浪漫粉',
    description: '温暖、亲和、活力 - 适合电商、社交类应用',
    primary: {
      // 玫瑰色系：温暖、亲和、活力（电商、社交应用）
      50: '#fff1f2',
      100: '#ffe4e6',
      200: '#fecdd3',
      300: '#fda4af',
      400: '#fb7185',
      500: '#f43f5e', // 主色 - 60%
      600: '#e11d48',
      700: '#be123c',
      800: '#9f1239',
      900: '#881337', // 深色 - 满足 WCAG 对比度 4.5:1
    },
  },
  orange: {
    name: 'orange',
    label: '活力橙',
    description: '温暖、活力、积极 - 适合电商、儿童类应用',
    primary: {
      // 橙色系：温暖、活力、注意力（电商、儿童产品）
      50: '#fff7ed',
      100: '#ffedd5',
      200: '#fed7aa',
      300: '#fdba74',
      400: '#fb923c',
      500: '#f97316', // 主色 - 60%
      600: '#ea580c',
      700: '#c2410c',
      800: '#9a3412',
      900: '#7c2d12', // 深色 - 满足 WCAG 对比度 4.5:1
    },
  },
}

// localStorage key
const THEME_STORAGE_KEY = 'app-theme'

// 获取当前主题
export const getCurrentTheme = (): string => {
  return localStorage.getItem(THEME_STORAGE_KEY) || 'default'
}

// 设置主题
export const setTheme = (themeName: string): void => {
  const theme = themes[themeName]
  if (!theme) {
    console.warn(`Theme "${themeName}" not found`)
    return
  }

  localStorage.setItem(THEME_STORAGE_KEY, themeName)

  // 更新 CSS 变量 - 使用 :root 选择器
  const root = document.documentElement
  
  // 先移除旧的主题类
  Object.keys(themes).forEach(key => {
    root.classList.remove(`theme-${key}`)
  })
  
  // 添加新主题类
  root.classList.add(`theme-${themeName}`)
  
  // 同时更新 CSS 变量（作为备用方案）
  Object.entries(theme.primary).forEach(([key, value]) => {
    root.style.setProperty(`--color-primary-${key}`, value)
  })

  // 触发自定义事件，通知主题变更
  window.dispatchEvent(new CustomEvent('themeChange', { detail: themeName }))
}

// 初始化主题
export const initTheme = (): void => {
  const currentTheme = getCurrentTheme()
  setTheme(currentTheme)
}
