import React, { useEffect, useState, useRef } from 'react'
import { getUserList, updateUserStatus, createUser, deleteUser, UserItem } from '@/api/user'
import { Button, Card, Select, Table, Pagination, Spinner, Modal, Input } from '@/components/ui'
import { useToast } from '@/components/ui/Toast'
import { Eye, EyeOff } from 'lucide-react'

const statusOptions = [
  { label: '全部', value: 0 },
  { label: '正常', value: 1 },
  { label: '禁用', value: 2 },
]

export const UserList: React.FC = () => {
  const [loading, setLoading] = useState(false)
  const [items, setItems] = useState<UserItem[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [perPage, setPerPage] = useState(10)
  const [status, setStatus] = useState(0)
  const [prevStatus, setPrevStatus] = useState<number | null>(null)

  // 使用 useRef 来跟踪是否已经加载过数据
  const hasLoadedRef = useRef(false)
  // 用于跟踪是否是用户主动改变页码
  const isUserPageChangeRef = useRef(false)
  const [showStatusModal, setShowStatusModal] = useState(false)
  const [statusItem, setStatusItem] = useState<UserItem | null>(null)
  const [showCreateModal, setShowCreateModal] = useState(false)
  const [showDeleteModal, setShowDeleteModal] = useState(false)
  const [deleteItem, setDeleteItem] = useState<UserItem | null>(null)
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)
  const [userNameError, setUserNameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [confirmPasswordError, setConfirmPasswordError] = useState('')
  const [createFormData, setCreateFormData] = useState({
    user_name: '',
    nick_name: '',
    password: '',
    confirmPassword: '',
  })
  const { addToast } = useToast()

  useEffect(() => {
    // 只有当 status 真正发生变化时才重新加载数据
    if (prevStatus !== null && prevStatus === status) {
      return
    }
    
    // 初始化时不立即加载，避免与下面的 useEffect 冲突
    if (!hasLoadedRef.current) {
      hasLoadedRef.current = true
      // 初始化时直接加载数据
      loadData()
      setPrevStatus(status)
      return
    }
    
    // 非初始化时，标记为用户操作，重置页码并加载第一页数据
    isUserPageChangeRef.current = true
    setPrevStatus(status)
    setPage(1)
  }, [status])

  useEffect(() => {
    // 初始化时由第一个 useEffect 处理，这里只处理用户主动改变页码
    if (!hasLoadedRef.current) {
      return
    }
    
    // 如果是用户主动改变页码（通过分页组件），则加载数据
    if (isUserPageChangeRef.current || page > 1) {
      isUserPageChangeRef.current = false
      loadData()
    }
  }, [page, perPage])

  const loadData = async () => {
    setLoading(true)
    try {
      const params: any = { page, perPage }
      if (status) params.status = status

      const res = await getUserList(params)
      setItems(res.rows || [])
      setTotal(res.count || 0)
    } catch (error) {
      console.error('Failed to load users', error)
    } finally {
      setLoading(false)
    }
  }

  const handleStatusChange = (item: UserItem) => {
    setStatusItem(item)
    setShowStatusModal(true)
  }

  const handleStatusConfirm = async () => {
    if (!statusItem) return
    const newStatus = statusItem.status === 1 ? 2 : 1
    const action = newStatus === 2 ? '禁用' : '启用'
    
    try {
      await updateUserStatus({ uid: statusItem.uid, status: newStatus })
      addToast(`${action}成功`, 'success')
      setShowStatusModal(false)
      setStatusItem(null)
      loadData()
    } catch (error) {
      console.error('Failed to update user status', error)
    }
  }

  const handleCreate = () => {
    setCreateFormData({ user_name: '', nick_name: '', password: '', confirmPassword: '' })
    setShowPassword(false)
    setShowConfirmPassword(false)
    setUserNameError('')
    setPasswordError('')
    setConfirmPasswordError('')
    setShowCreateModal(true)
  }

  const handleUserNameBlur = () => {
    if (createFormData.user_name && createFormData.user_name.length < 6) {
      setUserNameError('登录名至少需要6个字符')
    } else {
      setUserNameError('')
    }
  }

  const handlePasswordBlur = () => {
    if (createFormData.password && createFormData.password.length < 6) {
      setPasswordError('密码至少需要6个字符')
    } else {
      setPasswordError('')
    }
  }

  const handleConfirmPasswordBlur = () => {
    if (createFormData.confirmPassword && createFormData.password !== createFormData.confirmPassword) {
      setConfirmPasswordError('两次输入的密码不一致')
    } else {
      setConfirmPasswordError('')
    }
  }

  const handleCreateConfirm = async () => {
    if (!createFormData.user_name || !createFormData.nick_name || !createFormData.password || !createFormData.confirmPassword) {
      addToast('请填写所有必填项', 'warning')
      return
    }
    if (createFormData.user_name.length < 6) {
      addToast('登录名至少需要6个字符', 'warning')
      return
    }
    if (createFormData.password.length < 6) {
      addToast('密码至少需要6个字符', 'warning')
      return
    }
    if (createFormData.password !== createFormData.confirmPassword) {
      addToast('两次输入的密码不一致', 'warning')
      return
    }
    try {
      await createUser({
        user_name: createFormData.user_name,
        nick_name: createFormData.nick_name,
        password: createFormData.password,
      })
      addToast('创建成功', 'success')
      setShowCreateModal(false)
      loadData()
    } catch (error: any) {
      const errorMsg = error?.response?.msg || error?.message || '创建失败'
      addToast(errorMsg, 'error')
    }
  }

  const handleDelete = (item: UserItem) => {
    setDeleteItem(item)
    setShowDeleteModal(true)
  }

  const handleDeleteConfirm = async () => {
    if (!deleteItem) return
    try {
      await deleteUser({ uid: deleteItem.uid })
      addToast('删除成功', 'success')
      setShowDeleteModal(false)
      setDeleteItem(null)
      loadData()
    } catch (error: any) {
      const errorMsg = error?.response?.msg || error?.message || '删除失败'
      addToast(errorMsg, 'error')
    }
  }

  const columns = [
    {
      key: 'uid',
      label: 'UID',
      render: (item: UserItem) => (
        <span className="font-mono text-sm">{item.uid}</span>
      ),
    },
    {
      key: 'user_name',
      label: '登录名',
    },
    {
      key: 'nick_name',
      label: '昵称',
    },
    {
      key: 'status',
      label: '状态',
      render: (item: UserItem) => (
        <span
          className={`px-2 py-1 rounded text-xs ${
            item.status === 1
              ? 'bg-green-100 text-green-800'
              : 'bg-red-100 text-red-800'
          }`}
        >
          {item.status === 1 ? '正常' : '禁用'}
        </span>
      ),
    },
    {
      key: 'created_at',
      label: '创建时间',
      render: (item: UserItem) =>
        new Date(Number(item.created_at) * 1000).toLocaleString('zh-CN'),
    },
    {
      key: 'updated_at',
      label: '更新时间',
      render: (item: UserItem) =>
        new Date(Number(item.updated_at) * 1000).toLocaleString('zh-CN'),
    },
    {
      key: 'action',
      label: '操作',
      render: (item: UserItem) => (
        <div className="flex gap-2">
          <Button
            variant={item.status === 1 ? 'warning' : 'success'}
            size="sm"
            onClick={() => handleStatusChange(item)}
          >
            {item.status === 1 ? '禁用' : '启用'}
          </Button>
          <Button
            variant="danger"
            size="sm"
            onClick={() => handleDelete(item)}
          >
            删除
          </Button>
        </div>
      ),
    },
  ]

  return (
    <div>
      <Card 
        header="用户管理"
        actions={
          <Button onClick={handleCreate}>创建用户</Button>
        }
      />

      <Card className="mt-4">
        <div className="pl-4 pr-4 pt-4 mb-4">
          <Select
            label="状态"
            options={statusOptions}
            value={status}
            onChange={(e) => setStatus(Number(e.target.value))}
          />
        </div>

        {loading ? (
          <div className="flex justify-center py-8">
            <Spinner />
          </div>
        ) : (
          <>
            <Table
              columns={columns}
              data={items}
              rowKey="uid"
              className="table-db table-striped"
            />
            {items.length === 0 && !loading && (
              <div className="flex flex-col items-center justify-center py-20 text-gray-500">
                <div className="text-lg">暂无数据</div>
              </div>
            )}
            <Pagination
              current={page}
              total={total}
              pageSize={perPage}
              onChange={setPage}
              onPageSizeChange={setPerPage}
            />
          </>
        )}
      </Card>

      <Modal
        isOpen={showStatusModal}
        onClose={() => {
          setShowStatusModal(false)
          setStatusItem(null)
        }}
        title="确认操作"
        size="sm"
      >
        <div className="space-y-4">
          <p className="text-sm text-gray-600">
            确定要{statusItem?.status === 1 ? '禁用' : '启用'}该用户吗？
          </p>
          <div className="flex gap-3">
            <Button
              variant="secondary"
              onClick={() => {
                setShowStatusModal(false)
                setStatusItem(null)
              }}
              className="flex-1"
            >
              取消
            </Button>
            <Button
              variant={statusItem?.status === 1 ? 'warning' : 'success'}
              onClick={handleStatusConfirm}
              className="flex-1"
            >
              确定
            </Button>
          </div>
        </div>
      </Modal>

      <Modal
        isOpen={showCreateModal}
        onClose={() => {
          setShowCreateModal(false)
          setCreateFormData({ user_name: '', nick_name: '', password: '', confirmPassword: '' })
          setShowPassword(false)
          setShowConfirmPassword(false)
        }}
        title="创建用户"
        size="md"
      >
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              登录名
            </label>
            <Input
              placeholder="请输入登录名"
              value={createFormData.user_name}
              onChange={(e) => setCreateFormData({ ...createFormData, user_name: e.target.value })}
              onBlur={handleUserNameBlur}
            />
            {userNameError && <p className="mt-1 text-sm text-red-500">{userNameError}</p>}
          </div>
          <Input
            label="昵称"
            placeholder="请输入昵称"
            value={createFormData.nick_name}
            onChange={(e) => setCreateFormData({ ...createFormData, nick_name: e.target.value })}
          />
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              密码
            </label>
            <div className="relative">
              <input
                type={showPassword ? 'text' : 'password'}
                className={`w-full h-10 px-3 text-sm border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-400 bg-white/80 backdrop-blur-sm transition-all pr-10 ${
                  passwordError ? 'border-red-400' : 'border-gray-200 hover:border-purple-300'
                }`}
                placeholder="请输入密码"
                value={createFormData.password}
                onChange={(e) => setCreateFormData({ ...createFormData, password: e.target.value })}
                onBlur={handlePasswordBlur}
              />
              <button
                type="button"
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                onClick={() => setShowPassword(!showPassword)}
              >
                {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
              </button>
            </div>
            {passwordError && <p className="mt-1 text-sm text-red-500">{passwordError}</p>}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              确认密码
            </label>
            <div className="relative">
              <input
                type={showConfirmPassword ? 'text' : 'password'}
                className={`w-full h-10 px-3 text-sm border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-400 bg-white/80 backdrop-blur-sm transition-all pr-10 ${
                  confirmPasswordError ? 'border-red-400' : 'border-gray-200 hover:border-purple-300'
                }`}
                placeholder="请再次输入密码"
                value={createFormData.confirmPassword}
                onChange={(e) => setCreateFormData({ ...createFormData, confirmPassword: e.target.value })}
                onBlur={handleConfirmPasswordBlur}
              />
              <button
                type="button"
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                onClick={() => setShowConfirmPassword(!showConfirmPassword)}
              >
                {showConfirmPassword ? <EyeOff size={18} /> : <Eye size={18} />}
              </button>
            </div>
            {confirmPasswordError && <p className="mt-1 text-sm text-red-500">{confirmPasswordError}</p>}
          </div>
          <div className="flex gap-3">
            <Button
              variant="secondary"
              onClick={() => {
                setShowCreateModal(false)
                setCreateFormData({ user_name: '', nick_name: '', password: '', confirmPassword: '' })
                setShowPassword(false)
                setShowConfirmPassword(false)
              }}
              className="flex-1"
            >
              取消
            </Button>
            <Button
              variant="primary"
              onClick={handleCreateConfirm}
              className="flex-1"
            >
              确定
            </Button>
          </div>
        </div>
      </Modal>

      <Modal
        isOpen={showDeleteModal}
        onClose={() => {
          setShowDeleteModal(false)
          setDeleteItem(null)
        }}
        title="确认删除"
        size="sm"
      >
        <div className="space-y-4">
          <p className="text-sm text-gray-600">
            确定要删除用户 {deleteItem?.nick_name} 吗？此操作不可恢复。
          </p>
          <div className="flex gap-3">
            <Button
              variant="secondary"
              onClick={() => {
                setShowDeleteModal(false)
                setDeleteItem(null)
              }}
              className="flex-1"
            >
              取消
            </Button>
            <Button
              variant="danger"
              onClick={handleDeleteConfirm}
              className="flex-1"
            >
              确定删除
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  )
}
