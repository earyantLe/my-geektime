import React, { useEffect, useState } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { Loader2 } from 'lucide-react'
import { useAuthStore } from '@/store/auth'
import { login } from '@/api/auth'

export const GitHubCallback: React.FC = () => {
  const [searchParams] = useSearchParams()
  const navigate = useNavigate()
  const setAuth = useAuthStore((state) => state.setAuth)
  const setGeekAuth = useAuthStore((state) => state.setGeekAuth)
  const [error, setError] = useState('')

  useEffect(() => {
    const code = searchParams.get('code')
    if (!code) {
      setError('授权码无效')
      return
    }

    const handleGitHubLogin = async () => {
      try {
        const response = await login({
          type: 'github',
          data: { code },
        })

        if (response.status === 0 && response.token) {
          setAuth(response.token, response.user)
          if (response.user.geek_auth) {
            setGeekAuth(true)
          }
          navigate('/')
        } else {
          setError('登录失败')
        }
      } catch (err: any) {
        console.error('GitHub 登录错误:', err)
        setError(err.response?.data?.msg || err.message || 'GitHub 登录失败')
      }
    }

    handleGitHubLogin()
  }, [searchParams, navigate, setAuth, setGeekAuth])

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-primary-50 via-primary-100 to-primary-50">
        <div className="bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8 max-w-md text-center">
          <h2 className="text-2xl font-bold text-red-600 mb-4">登录失败</h2>
          <p className="text-gray-600 mb-6">{error}</p>
          <button
            onClick={() => navigate('/login')}
            className="px-6 py-3 bg-primary-600 text-white rounded-xl hover:bg-primary-700 transition-colors"
          >
            返回登录
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-primary-50 via-primary-100 to-primary-50">
      <div className="bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8 max-w-md text-center">
        <Loader2 className="w-10 h-10 animate-spin text-primary-600 mx-auto mb-4" />
        <h2 className="text-2xl font-bold text-gray-900 mb-2">正在处理 GitHub 登录...</h2>
        <p className="text-gray-500">请稍候</p>
      </div>
    </div>
  )
}

export default GitHubCallback
