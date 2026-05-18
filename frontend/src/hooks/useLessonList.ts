import { useState, useCallback } from 'react'
import { getTaskList, TaskItem } from '@/api/task'

interface LessonFilters {
  xstatus: number
  keywords: string
}


export const useLessonList = () => {
  const [lessonList, setLessonList] = useState<any[]>([])
  const [lessonLoading, setLessonLoading] = useState(false)
  const [lessonPage, setLessonPage] = useState(1)
  const [lessonTotal, setLessonTotal] = useState(0)
  const [lessonFilters, setLessonFilters] = useState<LessonFilters>({ xstatus: 0, keywords: '' })
  const [lessonLoadMoreLoading, setLessonLoadMoreLoading] = useState(false)
  const [lessonHasMore, setLessonHasMore] = useState(true)
  const [lessonShowBackToTop] = useState(false)
  const [selectedTask, setSelectedTask] = useState<TaskItem | null>(null)
  const [showLessonDrawer, setShowLessonDrawer] = useState(false)

  const loadLessonList = useCallback(async (taskId: string, page: number, filters: LessonFilters, isLoadMore = false) => {
    if (isLoadMore) {
      setLessonLoadMoreLoading(true)
    } else {
      setLessonLoading(true)
    }
    try {
      const params: any = { page, perPage: 10, task_pid: taskId }
      if (filters.xstatus) params.xstatus = filters.xstatus
      if (filters.keywords) params.keywords = filters.keywords

      const res = await getTaskList(params)
      const newRows = res.rows || []
      if (isLoadMore) {
        setLessonList((prev) => [...prev, ...newRows])
      } else {
        setLessonList(newRows)
      }
      setLessonTotal(res.count || 0)
      setLessonPage(page)
      const totalCount = res.count || 0
      const currentLoadedCount = isLoadMore ? lessonList.length + newRows.length : newRows.length
      setLessonHasMore(currentLoadedCount < totalCount && newRows.length > 0)
    } catch (error) {
      console.error('Failed to load lesson list', error)
    } finally {
      if (isLoadMore) {
        setLessonLoadMoreLoading(false)
      } else {
        setLessonLoading(false)
      }
    }
  }, [])

  const openLessonDrawer = useCallback((item: TaskItem) => {
    setSelectedTask(item)
    setShowLessonDrawer(true)
    setLessonPage(1)
    setLessonHasMore(true)
    setLessonFilters({ xstatus: 0, keywords: '' })
    loadLessonList(item.task_id, 1, { xstatus: 0, keywords: '' })
  }, [loadLessonList])

  const closeLessonDrawer = useCallback(() => {
    setShowLessonDrawer(false)
    setSelectedTask(null)
  }, [])

  const handleLessonPageChange = useCallback((page: number) => {
    if (selectedTask) {
      loadLessonList(selectedTask.task_id, page, lessonFilters, false)
    }
  }, [selectedTask, lessonFilters, loadLessonList])

  const handleLessonFilter = useCallback(() => {
    if (selectedTask) {
      setLessonPage(1)
      setLessonHasMore(true)
      loadLessonList(selectedTask.task_id, 1, lessonFilters)
    }
  }, [selectedTask, lessonFilters, loadLessonList])







  return {
    lessonList,
    lessonLoading,
    lessonPage,
    lessonTotal,
    lessonFilters,
    setLessonFilters,
    lessonLoadMoreLoading,
    lessonHasMore,
    lessonShowBackToTop,
    selectedTask,
    showLessonDrawer,
    setShowLessonDrawer,
    openLessonDrawer,
    closeLessonDrawer,
    handleLessonPageChange,
    handleLessonFilter,
    lessonLoadMoreRef: null,
    lessonScrollToTop: () => {},
  }
}
