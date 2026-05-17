import React from 'react'

export const ThemeTest: React.FC = () => {
  return (
    <div className="p-8 space-y-4">
      <h1 className="text-2xl font-bold mb-6">主题颜色测试</h1>
      
      {/* 背景色测试 */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold">背景色测试</h2>
        <div className="grid grid-cols-5 gap-2">
          <div className="bg-primary-50 p-4 rounded text-center">primary-50</div>
          <div className="bg-primary-100 p-4 rounded text-center">primary-100</div>
          <div className="bg-primary-200 p-4 rounded text-center">primary-200</div>
          <div className="bg-primary-300 p-4 rounded text-center">primary-300</div>
          <div className="bg-primary-400 p-4 rounded text-center text-white">primary-400</div>
        </div>
        <div className="grid grid-cols-5 gap-2">
          <div className="bg-primary-500 p-4 rounded text-center text-white">primary-500</div>
          <div className="bg-primary-600 p-4 rounded text-center text-white">primary-600</div>
          <div className="bg-primary-700 p-4 rounded text-center text-white">primary-700</div>
          <div className="bg-primary-800 p-4 rounded text-center text-white">primary-800</div>
          <div className="bg-primary-900 p-4 rounded text-center text-white">primary-900</div>
        </div>
      </div>

      {/* 文字色测试 */}
      <div className="space-y-2 mt-6">
        <h2 className="text-lg font-semibold">文字色测试</h2>
        <div className="space-y-1">
          <p className="text-primary-500">text-primary-500</p>
          <p className="text-primary-600">text-primary-600</p>
          <p className="text-primary-700">text-primary-700</p>
        </div>
      </div>

      {/* 边框色测试 */}
      <div className="space-y-2 mt-6">
        <h2 className="text-lg font-semibold">边框色测试</h2>
        <div className="flex gap-4">
          <div className="border-2 border-primary-500 p-4 rounded">border-primary-500</div>
          <div className="border-2 border-primary-600 p-4 rounded">border-primary-600</div>
          <div className="border-2 border-primary-700 p-4 rounded">border-primary-700</div>
        </div>
      </div>

      {/* 渐变测试 */}
      <div className="space-y-2 mt-6">
        <h2 className="text-lg font-semibold">渐变色测试</h2>
        <div className="bg-gradient-to-r from-primary-500 to-primary-700 p-4 rounded text-white">
          from-primary-500 to-primary-700
        </div>
      </div>

      {/* 按钮测试 */}
      <div className="space-y-2 mt-6">
        <h2 className="text-lg font-semibold">按钮测试</h2>
        <div className="flex gap-4">
          <button className="bg-primary-500 hover:bg-primary-600 text-white px-6 py-2 rounded-lg transition-colors">
            Primary Button
          </button>
          <button className="bg-primary-600 hover:bg-primary-700 text-white px-6 py-2 rounded-lg transition-colors">
            Secondary Button
          </button>
        </div>
      </div>
    </div>
  )
}
