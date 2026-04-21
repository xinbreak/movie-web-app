import type { User } from './authService'

const API_URL = import.meta.env.VITE_API_URL

export const deleteUser = async (userData: User) => {
  await fetch(`${API_URL}/users/${userData.id}`, {
    method: 'DELETE',
    headers: { 'Content-Type': 'application/json' }
  })
}
