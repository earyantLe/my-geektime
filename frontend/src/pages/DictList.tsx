import React, { useEffect, useState, useRef } from 'react'
import {
  getDictList,
  createDict,
  updateDict,
  deleteDict,
  DictItem,
} from '@/api/dict'
import { Button, Card, Input, Table, Pagination, Spinner, Modal } from '@/components/ui'
import { useToast } from '@/components/ui/Toast'
import { Plus, Edit, Trash2 } from 'lucide-react'

export const DictList: React.FC = () => {
  const [loading, setLoading] = useState(false)
  const [items, setItems] = useState<DictItem[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [perPage, setPerPage] = useState(10)

  // 使用 useRef 来跟踪是否已经加载过数据
  const hasLoadedRef = useRef(false)
  const [showModal, setShowModal] = useState(false)
  const [editingItem, setEditingItem] = useState<DictItem | null>(null)
  const [showDeleteModal, setShowDeleteModal] = useState(false)
  const [deleteItem, setDeleteItem] = useState<DictItem | null>(null)
  const [formData, setFormData] = useState({
    name: '',
    key: '',
    content: [{ type: 'string', value: '' }],
    summary: '',
    sort: 1,
  })
  const { addToast } = useToast()

  useEffect(() => {
    // 初始化时不立即加载，避免重复请求
    if (!hasLoadedRef.current) {
      hasLoadedRef.current = true
      return
    }
    
    loadData()
  }, [page, perPage])

  const loadData = async () => {
    setLoading(true)
    try {
      const res = await getDictList({ page, perPage })
      setItems(res.rows || [])
      setTotal(res.count || 0)
    } catch (error) {
      console.error('Failed to load dict', error)
    } finally {
      setLoading(false)
    }
  }

  const handleAdd = () => {
    setEditingItem(null)
    setFormData({
      name: '',
      key: '',
      content: [{ type: 'string', value: '' }],
      summary: '',
      sort: 1,
    })
    setShowModal(true)
  }

  const handleEdit = (item: DictItem) => {
    setEditingItem(item)
    setFormData({
      name: item.name,
      key: item.key,
      content: item.content || [{ type: 'string', value: '' }],
      summary: item.summary || '',
      sort: item.sort,
    })
    setShowModal(true)
  }

  const handleDelete = (item: DictItem) => {
    setDeleteItem(item)
    setShowDeleteModal(true)
  }

  const handleDeleteConfirm = async () => {
    if (!deleteItem) return
    try {
      await deleteDict(deleteItem.id)
      addToast('删除成功', 'success')
      setShowDeleteModal(false)
      setDeleteItem(null)
      loadData()
    } catch (error) {
      console.error('Failed to delete', error)
    }
  }

  const handleSubmit = async () => {
    if (!formData.name || !formData.key) {
      addToast('请填写名称和唯一标识', 'warning')
      return
    }
    try {
      if (editingItem) {
        await updateDict({
          id: editingItem.id,
          ...formData,
        })
        addToast('更新成功', 'success')
      } else {
        await createDict(formData)
        addToast('创建成功', 'success')
      }
      setShowModal(false)
      loadData()
    } catch (error) {
      console.error('Failed to save', error)
    }
  }

  const handleContentAdd = () => {
    setFormData((prev) => ({
      ...prev,
      content: [...prev.content, { type: 'string', value: '' }],
    }))
  }

  const handleContentRemove = (index: number) => {
    setFormData((prev) => ({
      ...prev,
      content: prev.content.filter((_, i) => i !== index),
    }))
  }

  const handleContentChange = (index: number, field: string, value: any) => {
    setFormData((prev) => ({
      ...prev,
      content: prev.content.map((item, i) =>
        i === index ? { ...item, [field]: value } : item
      ),
    }))
  }

  const columns = [
    { key: 'id', label: 'ID' },
    { key: 'name', label: '名称' },
    { key: 'key', label: '唯一标识' },
    {
      key: 'sort',
      label: '排序',
    },
    {
      key: 'created',
      label: '创建时间',
      render: (item: DictItem) => new Date(item.created).toLocaleString('zh-CN'),
    },
    {
      key: 'updated',
      label: '更新时间',
      render: (item: DictItem) => new Date(item.updated).toLocaleString('zh-CN'),
    },
    {
      key: 'action',
      label: '操作',
      render: (item: DictItem) => (
        <div className="flex gap-2">
          <Button variant="light" size="sm" onClick={() => handleEdit(item)}>
            <Edit size={14} />
          </Button>
          <Button variant="danger" size="sm" onClick={() => handleDelete(item)}>
            <Trash2 size={14} />
          </Button>
        </div>
      ),
    },
  ]

  return (
    <div>
      <Card
        header="字典表"
        actions={
          <Button variant="info" size="sm" onClick={handleAdd}>
            <Plus size={14} className="mr-1" />
            添加字典
          </Button>
        }
      />

      <Card className="mt-4">
        {loading ? (
          <div className="flex justify-center py-8">
            <Spinner />
          </div>
        ) : (
          <>
            <Table
              columns={columns}
              data={items}
              rowKey="id"
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
        isOpen={showModal}
        onClose={() => setShowModal(false)}
        title={editingItem ? '编辑字典' : '添加字典'}
        size="lg"
      >
        <div className="space-y-4">
          <Input
            label="名称"
            value={formData.name}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, name: e.target.value }))
            }
            required
          />
          <Input
            label="唯一标识"
            value={formData.key}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, key: e.target.value }))
            }
            disabled={!!editingItem}
            required
          />

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              内容
            </label>
            {formData.content.map((item, index) => (
              <div key={index} className="flex gap-2 mb-2">
                <select
                  value={item.type}
                  onChange={(e) =>
                    handleContentChange(index, 'type', e.target.value)
                  }
                  className="border rounded px-2 py-1"
                >
                  <option value="string">文本</option>
                  <option value="number">数字</option>
                  <option value="boolean">布尔值</option>
                </select>
                <Input
                  placeholder="值"
                  value={String(item.value)}
                  onChange={(e) =>
                    handleContentChange(index, 'value', e.target.value)
                  }
                />
                <Button
                  variant="danger"
                  size="sm"
                  onClick={() => handleContentRemove(index)}
                >
                  删除
                </Button>
              </div>
            ))}
            <Button variant="light" size="sm" onClick={handleContentAdd}>
              添加项
            </Button>
          </div>

          <textarea
            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="描述"
            rows={3}
            value={formData.summary}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, summary: e.target.value }))
            }
          />

          <Input
            label="排序"
            type="number"
            value={String(formData.sort)}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, sort: Number(e.target.value) }))
            }
          />

          <div className="flex gap-4 pt-4">
            <Button onClick={handleSubmit}>保存</Button>
            <Button variant="light" onClick={() => setShowModal(false)}>
              取消
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
            确定要删除编号【{deleteItem?.id}】？
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
            <Button variant="danger" onClick={handleDeleteConfirm} className="flex-1">
              确定
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  )
}
