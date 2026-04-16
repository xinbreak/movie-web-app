const API_URL = import.meta.env.VITE_API_URL

export interface User {
  id: string
  username: string
  email: string
  avatar_url: string | null
}

export const loginRequest = async (
  email: string,
  password: string
): Promise<User> => {
  const response = await fetch(`${API_URL}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  })

  const data = await response.json().catch(() => ({}))

  if (!response.ok) {
    throw new Error(data.message || 'Login failed')
  }

  return data as User
}

export const registerRequest = async (
  userData: Record<string, string>
): Promise<void> => {
  const response = await fetch(`${API_URL}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(userData)
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.message || 'Registration failed')
  }
}
