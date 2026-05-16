import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { BookOpen, Eye, EyeOff, Loader2, User, Lock, AlertCircle, Sparkles, Target, Heart, Github, Mail } from 'lucide-react'
import { useAuthStore } from '@/store/auth'
import { login, register, getConfig, ConfigData, sendEmailCode } from '@/api/auth'

type TabType = 'login' | 'register'

const Bubble = ({ className, style }: { className: string; style?: React.CSSProperties }) => (
  <div className={`absolute rounded-full bg-white/10 backdrop-blur-3xl ${className}`} style={style} />
)

const isValidEmail = (email: string) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

export const Login: React.FC = () => {
  const [activeTab, setActiveTab] = useState<TabType>('login')
  const [loginMethod, setLoginMethod] = useState<'name' | 'email'>('name')
  const [account, setAccount] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [verificationCode, setVerificationCode] = useState('')
  const [loading, setLoading] = useState(false)
  const [sendingCode, setSendingCode] = useState(false)
  const [countdown, setCountdown] = useState(0)
  const [error, setError] = useState('')
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)
  const [config, setConfig] = useState<ConfigData | null>(null)
  const [configLoading, setConfigLoading] = useState(true)
  
  const navigate = useNavigate()
  const setAuth = useAuthStore((state) => state.setAuth)
  const setGeekAuth = useAuthStore((state) => state.setGeekAuth)

  useEffect(() => {
    const fetchConfig = async () => {
      try {
        const data = await getConfig()
        setConfig(data)
        if (data.login_types && Array.isArray(data.login_types) && data.login_types.length > 0) {
          setLoginMethod(data.login_types[0] as 'name' | 'email')
        }
      } catch (err: any) {
        console.error('获取配置失败:', err)
        // 设置默认配置，避免页面崩溃
        setConfig({
          login_types: ['name'],
          register_types: ['name'],
          login_guest: [],
          email_login: { enabled: false, verify_code: false },
          github_login: { enabled: false }
        })
      } finally {
        setConfigLoading(false)
      }
    }
    fetchConfig()
  }, [])

  useEffect(() => {
    if (activeTab === 'register') {
      setAccount('')
      setEmail('')
      setPassword('')
      setConfirmPassword('')
      setVerificationCode('')
      setError('')
    }
  }, [activeTab])

  useEffect(() => {
    if (config?.login_guest && Array.isArray(config.login_guest) && activeTab === 'login') {
      config.login_guest.forEach((guest: { type: string; account: string; password: string }) => {
        if (loginMethod === 'name' && guest.type === 'name') {
          setAccount(guest.account)
          setPassword(guest.password)
        } else if (loginMethod === 'email' && guest.type === 'email') {
          setEmail(guest.account)
          setPassword(guest.password)
        }
      })
    }
  }, [activeTab, config, loginMethod])

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000)
      return () => clearTimeout(timer)
    }
  }, [countdown])

  const sendCode = async () => {
    if (!email || !isValidEmail(email)) {
      setError('请输入有效的邮箱地址')
      return
    }
    setSendingCode(true)
    setError('')
    try {
      await sendEmailCode(email)
      setCountdown(60)
    } catch (err: any) {
      setError(err.message || '发送验证码失败')
    } finally {
      setSendingCode(false)
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')

    const currentAccount = loginMethod === 'email' ? email : account
    if (!currentAccount || !password) {
      setError('请填写所有必填项')
      return
    }

    if (loginMethod === 'name' && currentAccount.length < 6) {
      setError('登录名不少于6个字符')
      return
    }

    if (loginMethod === 'email' && !isValidEmail(currentAccount)) {
      setError('请输入有效的邮箱地址')
      return
    }

    if (password.length < 6) {
      setError('密码不少于6个字符')
      return
    }

    if (activeTab === 'register' && password !== confirmPassword) {
      setError('两次输入的密码不一致')
      return
    }

    if (activeTab === 'register' && loginMethod === 'email' && (!verificationCode || verificationCode.length < 6)) {
      setError('请输入6位验证码')
      return
    }

    setLoading(true)

    try {
      if (activeTab === 'login') {
        const loginType = loginMethod === 'email' ? 'email' : 'name'
        let loginData: any
        
        if (loginType === 'email') {
          loginData = { email: currentAccount, password }
        } else {
          loginData = { account: currentAccount, password }
        }
        
        const response = await login({
          type: loginType,
          data: loginData,
        })
        
        if (response.status === 0 && response.token) {
          setAuth(response.token, response.user)
          if (response.user.geek_auth) {
            setGeekAuth(true)
          }
          navigate('/')
        }
      } else {
        const registerType = loginMethod === 'email' ? 'email' : 'name'
        
        if (registerType === 'email') {
          await register({
            type: 'email',
            data: { email: currentAccount, code: verificationCode, password },
          })
        } else {
          await register({
            type: 'name',
            data: { account: currentAccount, password },
          })
        }
        setActiveTab('login')
        setError('')
      }
    } catch (err: any) {
      setError(err.message || '操作失败，请重试')
    } finally {
      setLoading(false)
    }
  }

  const canRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.length > 0
  const showEmailLogin = config?.login_types && Array.isArray(config.login_types) && config.login_types.includes('email')
  const showNameLogin = config?.login_types && Array.isArray(config.login_types) && config.login_types.includes('name')
  const showEmailRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.includes('email')
  const showNameRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.includes('name')
  const showGitHubLogin = config?.github_login?.enabled

  return (
    <div className="min-h-screen flex bg-gradient-to-br from-purple-600 via-pink-500 to-orange-400 relative overflow-hidden">
      <Bubble className="w-96 h-96 -top-48 -left-48 animate-pulse" />
      <Bubble className="w-64 h-64 top-1/4 -right-32 animate-bounce" style={{ animationDuration: '3s' }} />
      <Bubble className="w-80 h-80 -bottom-40 left-1/3 animate-pulse" style={{ animationDuration: '4s' }} />
      <Bubble className="w-48 h-48 bottom-1/4 right-1/4 animate-bounce" style={{ animationDuration: '5s' }} />
      
      <div className="hidden lg:flex lg:w-1/2 relative z-10">
        <div className="max-w-lg mx-auto px-12 py-16 flex flex-col justify-center text-white">
          <div className="flex items-center gap-4 mb-12">
            <div className="p-4 bg-white/20 backdrop-blur-sm rounded-2xl">
              <BookOpen className="w-10 h-10" />
            </div>
            <h1 className="text-4xl font-bold">我的极客时间</h1>
          </div>
          
          <h2 className="text-5xl font-bold mb-6 leading-tight">
            让学习成为
            <br />
            <span className="bg-gradient-to-r from-yellow-200 to-pink-200 bg-clip-text text-transparent">
              一种习惯
            </span>
          </h2>
          <p className="text-white/80 text-xl mb-12 leading-relaxed">
            汇聚行业专家，提供高质量的技术内容。
            无论你是刚入门的新手还是经验丰富的工程师，
            这里都有适合你的学习路径。
          </p>

          <div className="space-y-6">
            <div className="flex items-center gap-4">
              <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                <Sparkles className="w-6 h-6" />
              </div>
              <div>
                <h3 className="font-semibold text-lg">精选内容</h3>
                <p className="text-white/70">定期更新的技术专栏和实战课程</p>
              </div>
            </div>
            <div className="flex items-center gap-4">
              <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                <Target className="w-6 h-6" />
              </div>
              <div>
                <h3 className="font-semibold text-lg">学习路径</h3>
                <p className="text-white/70">系统化的课程设计，助你快速成长</p>
              </div>
            </div>
            <div className="flex items-center gap-4">
              <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                <Heart className="w-6 h-6" />
              </div>
              <div>
                <h3 className="font-semibold text-lg">社区互动</h3>
                <p className="text-white/70">与技术爱好者共同学习</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="w-full lg:w-1/2 flex items-center justify-center px-4 py-8 relative z-10">
        <div className="w-full max-w-md">
          <div className="bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8">
            {configLoading ? (
              <div className="flex justify-center items-center py-12">
                <Loader2 className="w-8 h-8 animate-spin text-purple-600" />
              </div>
            ) : (
              <>
            <div className="lg:hidden flex justify-center mb-8">
              <div className="flex items-center gap-3">
                <div className="p-3 bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl">
                  <BookOpen className="w-6 h-6 text-white" />
                </div>
                <span className="text-xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent">
                  我的极客时间
                </span>
              </div>
            </div>

            <div className="text-center mb-8">
              <h2 className="text-3xl font-bold text-gray-900 mb-2">
                {activeTab === 'login' ? '欢迎回来' : '创建账号'}
              </h2>
              <p className="text-gray-500">
                {activeTab === 'login' ? '登录以继续您的学习之旅' : '加入我们，开始新的技术探索'}
              </p>
            </div>

            {((activeTab === 'login' && showEmailLogin && showNameLogin) || (activeTab === 'register' && showEmailRegister && showNameRegister)) && (
              <div className="flex bg-gray-50 rounded-xl p-1 mb-6">
                {((activeTab === 'login' && showNameLogin) || (activeTab === 'register' && showNameRegister)) && (
                  <button
                    onClick={() => setLoginMethod('name')}
                    className={`flex-1 py-2 px-3 rounded-lg text-sm font-medium transition-all flex items-center justify-center gap-2 ${
                      loginMethod === 'name'
                        ? 'bg-white text-purple-600 shadow-sm'
                        : 'text-gray-500 hover:text-gray-700'
                    }`}
                  >
                    <User className="w-4 h-4" />
                    登录名
                  </button>
                )}
                {((activeTab === 'login' && showEmailLogin) || (activeTab === 'register' && showEmailRegister)) && (
                  <button
                    onClick={() => setLoginMethod('email')}
                    className={`flex-1 py-2 px-3 rounded-lg text-sm font-medium transition-all flex items-center justify-center gap-2 ${
                      loginMethod === 'email'
                        ? 'bg-white text-purple-600 shadow-sm'
                        : 'text-gray-500 hover:text-gray-700'
                    }`}
                  >
                    <Mail className="w-4 h-4" />
                    邮箱
                  </button>
                )}
              </div>
            )}

            {error && (
              <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-xl flex items-start gap-3">
                <AlertCircle className="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
                <span className="text-red-700 text-sm">{error}</span>
              </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-5">
              {loginMethod === 'name' ? (
                <div className="relative">
                  <User className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                  <input
                    type="text"
                    placeholder="请输入登录名"
                    value={account}
                    onFocus={() => setError('')}
                    onChange={(e) => setAccount(e.target.value)}
                    onBlur={() => {
                      if (!account) return
                      if (account.length < 6) {
                        setError('登录名不少于6个字符')
                      } else {
                        setError('')
                      }
                    }}
                    className="w-full pl-12 pr-4 py-3.5 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all bg-gray-50 focus:bg-white"
                    required
                  />
                </div>
              ) : (
                <div className="relative">
                  <Mail className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                  <input
                    type="email"
                    placeholder="请输入邮箱地址"
                    value={email}
                    onFocus={() => setError('')}
                    onChange={(e) => setEmail(e.target.value)}
                    onBlur={() => {
                      if (!email) return
                      if (!isValidEmail(email)) {
                        setError('请输入有效的邮箱地址')
                      } else {
                        setError('')
                      }
                    }}
                    className="w-full pl-12 pr-4 py-3.5 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all bg-gray-50 focus:bg-white"
                    required
                  />
                </div>
              )}

              {loginMethod === 'email' && activeTab === 'register' && (
                <div className="relative">
                  <input
                    type="text"
                    placeholder="请输入6位验证码"
                    value={verificationCode}
                    onChange={(e) => setVerificationCode(e.target.value.replace(/\D/g, '').slice(0, 6))}
                    className="w-full pl-4 pr-24 py-3.5 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all bg-gray-50 focus:bg-white"
                    required
                    maxLength={6}
                  />
                  <button
                    type="button"
                    onClick={sendCode}
                    disabled={countdown > 0 || sendingCode}
                    className="absolute right-2 top-1/2 -translate-y-1/2 px-3 py-1.5 text-sm font-medium text-purple-600 hover:text-purple-700 disabled:text-gray-400 disabled:cursor-not-allowed transition-colors"
                  >
                    {sendingCode ? (
                      <Loader2 className="w-4 h-4 animate-spin" />
                    ) : countdown > 0 ? (
                      `${countdown}s`
                    ) : (
                      '获取验证码'
                    )}
                  </button>
                </div>
              )}

              <div className="relative">
                <Lock className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                <input
                  type={showPassword ? 'text' : 'password'}
                  placeholder="请输入密码"
                  value={password}
                  onFocus={() => setError('')}
                  onChange={(e) => setPassword(e.target.value)}
                  onBlur={() => {
                    if (!password) return
                    if (password.length < 6) {
                      setError('密码不少于6个字符')
                    } else {
                      setError('')
                    }
                  }}
                  className="w-full pl-12 pr-12 py-3.5 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all bg-gray-50 focus:bg-white"
                  required
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
                >
                  {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
                </button>
              </div>

              {activeTab === 'register' && (
                <div className="relative">
                  <Lock className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                  <input
                    type={showConfirmPassword ? 'text' : 'password'}
                    placeholder="请再次输入密码"
                    value={confirmPassword}
                    onFocus={() => setError('')}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    onBlur={() => {
                      if (!confirmPassword) return
                      if (confirmPassword !== password) {
                        setError('两次输入的密码不一致')
                      } else {
                        setError('')
                      }
                    }}
                    className="w-full pl-12 pr-12 py-3.5 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all bg-gray-50 focus:bg-white"
                    required
                  />
                  <button
                    type="button"
                    onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                    className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
                  >
                    {showConfirmPassword ? <EyeOff size={20} /> : <Eye size={20} />}
                  </button>
                </div>
              )}

              <button
                type="submit"
                disabled={loading}
                className="w-full mt-4 py-3.5 text-base font-semibold text-white bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 rounded-xl transition-all duration-300 shadow-lg hover:shadow-xl disabled:opacity-70 disabled:cursor-not-allowed flex items-center justify-center gap-2"
              >
                {loading ? (
                  <>
                    <Loader2 className="w-5 h-5 animate-spin" />
                    处理中...
                  </>
                ) : (
                  activeTab === 'login' ? '登录' : '注册'
                )}
              </button>
            </form>

            {canRegister && (
              <p className="text-center text-sm text-gray-500 mt-8">
                {activeTab === 'login' ? (
                  <>还没有账号？<button onClick={() => setActiveTab('register')} className="text-purple-600 hover:text-purple-700 font-semibold transition-colors">立即注册</button></>
                ) : (
                  <>已有账号？<button onClick={() => setActiveTab('login')} className="text-purple-600 hover:text-purple-700 font-semibold transition-colors">立即登录</button></>
                )}
              </p>
            )}

            {showGitHubLogin && (
              <>
                <div className="relative my-6">
                  <div className="absolute inset-0 flex items-center">
                    <div className="w-full border-t border-gray-200"></div>
                  </div>
                  <div className="relative flex justify-center text-sm">
                    <span className="px-4 bg-white text-gray-500">或者</span>
                  </div>
                </div>

                <button
                  type="button"
                  onClick={() => {
                    if (config?.github_login?.client_id) {
                      const authUrl = `${config.github_login.auth_url}?client_id=${config.github_login.client_id}&redirect_uri=${encodeURIComponent(config.github_login.redirect_url || '')}&scope=${config.github_login.scope || 'read:user user:email'}`
                      window.location.href = authUrl
                    } else {
                      setError('GitHub 登录未配置')
                    }
                  }}
                  className="w-full py-3 px-4 border border-gray-200 rounded-xl hover:bg-gray-50 transition-colors flex items-center justify-center gap-3 font-medium text-gray-700"
                >
                  <Github className="w-5 h-5" />
                  使用 GitHub 登录
                </button>
              </>
            )}
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
