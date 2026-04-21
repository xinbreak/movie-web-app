import type { ReactNode } from 'react'
import { Navigate, Outlet } from 'react-router-dom'

interface ProtectedRouteProps {
  children?: ReactNode
}

export const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const isAuth = localStorage.getItem('isAuthorized') === 'true'

  if (!isAuth) {
    return <Navigate to="/login" replace />
  }

  return children ? <>{children}</> : <Outlet />
}
