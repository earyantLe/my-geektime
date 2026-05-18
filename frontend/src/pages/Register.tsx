import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { BookOpen, Eye, EyeOff, Loader2, User, Lock, AlertCircle, Mail, CheckCircle } from 'lucide-react'
import { register, getConfig, ConfigData, sendEmailCode } from '@/api/auth'

type RegisterMethod = 'name' | 'email'

const Bubble = ({ className, style }: { className: string; style?: React.CSSProperties }) => (
  <div className={`absolute rounded-full bg-white/10 backdrop-blur-3xl ${className}`} style={style} />
)

const isValidEmail = (email: string) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

export const Register: React.FC = () => {
  const [registerMethod, setRegisterMethod] = useState<RegisterMethod>('name')
  const [account, setAccount] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [verificationCode, setVerificationCode] = useState('')
  const [loading, setLoading] = useState(false)
  const [sendingCode, setSendingCode] = useState(false)
  const [countdown, setCountdown] = useState(0)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)
  const [config, setConfig] = useState<ConfigData | null>(null)
  const [configLoading, setConfigLoading] = useState(true)
  
  const navigate = useNavigate()

  useEffect(() => {
    const fetchConfig = async () => {
      try {
        const data = await getConfig()
        setConfig(data)
        // register_types 为数组，如果为空则表示不允许注册
        if (data.register_types && Array.isArray(data.register_types) && data.register_types.length > 0) {
          setRegisterMethod(data.register_types[0] as 'name' | 'email')
        }
      } catch (err: any) {
        console.error('获取配置失败:', err)
        // 设置默认配置，避免页面崩溃
        setConfig({
          register_types: [],
          login_types: ['name'],
          login_guests: [],
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
    setSuccess('')
    try {
      await sendEmailCode(email)
      setCountdown(60)
      setSuccess('验证码已发送到您的邮箱')
    } catch (err: any) {
      setError(err.message || '发送验证码失败')
    } finally {
      setSendingCode(false)
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setSuccess('')

    const currentAccount = registerMethod === 'email' ? email : account
    if (!currentAccount || !password) {
      setError('请填写所有必填项')
      return
    }

    if (registerMethod === 'name' && currentAccount.length < 6) {
      setError('用户名不少于6个字符')
      return
    }

    if (registerMethod === 'email' && !isValidEmail(currentAccount)) {
      setError('请输入有效的邮箱地址')
      return
    }

    if (password.length < 6) {
      setError('密码不少于6个字符')
      return
    }

    if (password !== confirmPassword) {
      setError('两次输入的密码不一致')
      return
    }

    if (registerMethod === 'email' && (!verificationCode || verificationCode.length < 6)) {
      setError('请输入6位验证码')
      return
    }

    setLoading(true)

    try {
      const registerType = registerMethod === 'email' ? 'email' : 'name'
      
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
      
      setSuccess('注册成功！即将跳转到登录页面...')
      setTimeout(() => {
        navigate('/login')
      }, 2000)
    } catch (err: any) {
      setError(err.message || '注册失败，请重试')
    } finally {
      setLoading(false)
    }
  }

  const showEmailRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.includes('email')
  const showNameRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.includes('name')
  // register_types 为空时表示不允许注册
  const canRegister = config?.register_types && Array.isArray(config.register_types) && config.register_types.length > 0

  return (
    <div className="min-h-screen flex bg-gradient-to-br from-primary-50 via-primary-100 to-primary-50 relative overflow-hidden">
      <Bubble className="w-96 h-96 -top-48 -left-48 animate-pulse" />
      <Bubble className="w-64 h-64 top-1/4 -right-32 animate-bounce" style={{ animationDuration: '3s' }} />
      <Bubble className="w-80 h-80 -bottom-40 left-1/3 animate-pulse" style={{ animationDuration: '4s' }} />
      <Bubble className="w-48 h-48 bottom-1/4 right-1/4 animate-bounce" style={{ animationDuration: '5s' }} />
      
      <div className="hidden lg:flex lg:w-1/2 relative z-10">
        <div className="max-w-lg mx-auto px-12 py-16 flex flex-col justify-center text-primary-900">
          <div className="flex items-center gap-4 mb-12">
            <div className="p-4 bg-primary-600/20 backdrop-blur-sm rounded-2xl">
              <BookOpen className="w-10 h-10 text-primary-600" />
            </div>
            <h1 className="text-4xl font-bold">我的极客时间</h1>
          </div>
          
          <h2 className="text-5xl font-bold mb-6 leading-tight">
            开启你的
            <br />
            <span className="bg-gradient-to-r from-primary-500 to-primary-600 bg-clip-text text-transparent">
              技术学习之旅
            </span>
          </h2>
          <p className="text-primary-700/80 text-xl mb-12 leading-relaxed">
            加入我们的社区，与成千上万的开发者一起学习成长。
            获取最新的技术资讯、实战教程和行业洞察。
          </p>

          <div className="space-y-6">
            <div className="flex items-center gap-4 p-4 bg-white/50 backdrop-blur-sm rounded-xl">
              <div className="p-3 bg-green-100 rounded-xl">
                <CheckCircle className="w-6 h-6 text-green-600" />
              </div>
              <div>
                <h3 className="font-semibold text-lg">免费注册</h3>
                <p className="text-primary-600/70">立即开始你的学习旅程</p>
              </div>
            </div>
            <div className="flex items-center gap-4 p-4 bg-white/50 backdrop-blur-sm rounded-xl">
              <div className="p-3 bg-blue-100 rounded-xl">
                <Mail className="w-6 h-6 text-blue-600" />
              </div>
              <div>
                <h3 className="font-semibold text-lg">邮箱验证</h3>
                <p className="text-primary-600/70">安全可靠的账户保护</p>
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
                <Loader2 className="w-8 h-8 animate-spin text-primary-600" />
              </div>
            ) : !canRegister ? (
              <>
                <div className="lg:hidden flex justify-center mb-8">
                  <div className="flex items-center gap-3">
                    <div className="p-3 bg-primary-100 rounded-xl">
                      <BookOpen className="w-6 h-6 text-primary-600" />
                    </div>
                    <span className="text-xl font-bold text-primary-700">
                      我的极客时间
                    </span>
                  </div>
                </div>

                <div className="text-center mb-8">
                  <div className="p-4 bg-gray-100 rounded-2xl inline-block mb-4">
                    <AlertCircle className="w-12 h-12 text-gray-500" />
                  </div>
                  <h2 className="text-3xl font-bold text-gray-900 mb-2">
                    暂不开放注册
                  </h2>
                  <p className="text-gray-500">
                    当前系统暂不开放新用户注册
                  </p>
                </div>

                <p className="text-center text-sm text-gray-500">
                  已有账号？
                  <button 
                    onClick={() => navigate('/login')} 
                    className="text-primary-600 hover:text-primary-700 font-semibold transition-colors"
                  >
                    立即登录
                  </button>
                </p>
              </>
            ) : (
              <>
                <div className="lg:hidden flex justify-center mb-8">
                  <div className="flex items-center gap-3">
                    <div className="p-3 bg-primary-100 rounded-xl">
                      <BookOpen className="w-6 h-6 text-primary-600" />
                    </div>
                    <span className="text-xl font-bold text-primary-700">
                      我的极客时间
                    </span>
                  </div>
                </div>

                <div className="text-center mb-8">
                  <h2 className="text-3xl font-bold text-gray-900 mb-2">
                    创建账号
                  </h2>
                  <p className="text-gray-500">
                    加入我们，开始新的技术探索
                  </p>
                </div>

                {(showEmailRegister && showNameRegister) && (
                  <div className="flex bg-gray-50 rounded-xl p-1 mb-6">
                    {showNameRegister && (
                      <button
                        onClick={() => setRegisterMethod('name')}
                        className={`flex-1 py-2 px-3 rounded-lg text-sm font-medium transition-all flex items-center justify-center gap-2 ${
                          registerMethod === 'name'
                            ? 'bg-white text-primary-600 shadow-sm'
                            : 'text-gray-500 hover:text-gray-700'
                        }`}
                      >
                        <User className="w-4 h-4" />
                        用户名
                      </button>
                    )}
                    {showEmailRegister && (
                      <button
                        onClick={() => setRegisterMethod('email')}
                        className={`flex-1 py-2 px-3 rounded-lg text-sm font-medium transition-all flex items-center justify-center gap-2 ${
                          registerMethod === 'email'
                            ? 'bg-white text-primary-600 shadow-sm'
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

                {success && (
                  <div className="mb-6 p-4 bg-green-50 border border-green-200 rounded-xl flex items-start gap-3">
                    <CheckCircle className="w-5 h-5 text-green-500 flex-shrink-0 mt-0.5" />
                    <span className="text-green-700 text-sm">{success}</span>
                  </div>
                )}

                <form onSubmit={handleSubmit} className="space-y-5">
                  {registerMethod === 'name' ? (
                    <div className="relative">
                      <User className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                      <input
                        type="text"
                        placeholder="请输入用户名（至少6个字符）"
                        value={account}
                        onFocus={() => setError('')}
                        onChange={(e) => setAccount(e.target.value)}
                        onBlur={() => {
                          if (!account) return
                          if (account.length < 6) {
                            setError('用户名不少于6个字符')
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

                  {registerMethod === 'email' && (
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
                        className="absolute right-2 top-1/2 -translate-y-1/2 px-3 py-1.5 text-sm font-medium text-primary-600 hover:text-primary-700 disabled:text-gray-400 disabled:cursor-not-allowed transition-colors"
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
                      placeholder="请输入密码（至少6个字符）"
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

                  <button
                    type="submit"
                    disabled={loading}
                    className="w-full mt-4 py-3.5 text-base font-semibold text-primary-700 bg-primary-50 hover:bg-primary-100 border-2 border-primary-200 rounded-xl transition-all duration-300 shadow-sm hover:shadow-md disabled:opacity-70 disabled:cursor-not-allowed flex items-center justify-center gap-2"
                  >
                    {loading ? (
                      <>
                        <Loader2 className="w-5 h-5 animate-spin" />
                        注册中...
                      </>
                    ) : (
                      '创建账号'
                    )}
                  </button>
                </form>

                <p className="text-center text-sm text-gray-500 mt-8">
                  已有账号？
                  <button 
                    onClick={() => navigate('/login')} 
                    className="text-primary-600 hover:text-primary-700 font-semibold transition-colors"
                  >
                    立即登录
                  </button>
                </p>
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}