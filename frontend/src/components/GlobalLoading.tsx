import React from 'react'
import { useLoadingStore } from '@/store/loading'
import { Spinner } from '@/components/ui'

export const GlobalLoading: React.FC = () => {
  const isLoading = useLoadingStore((state) => state.isLoading)

  if (!isLoading) return null

  return (
    <div className="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 z-[9999]">
      <Spinner size="lg" />
    </div>
  )
}
