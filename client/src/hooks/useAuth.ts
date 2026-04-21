import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginRequest } from '../api/authService'
import type { User } from '../api/authService'
import { validateEmail } from '../utils/validation'

interface FormState {
  isError: boolean
}

export const useAuth = () => {
  const navigate = useNavigate()

  const loginAction = async (
    _prevState: FormState,
    formData: FormData
  ): Promise<FormState> => {
    const email = formData.get('email') as string
    const password = formData.get('password') as string

    if (!validateEmail(email)) {
      return { isError: true }
    }

    try {
      const user = await loginRequest(email, password)

      localStorage.setItem('isAuthorized', 'true')
      localStorage.setItem('current_user', JSON.stringify(user))

      const savedUsers: User[] = JSON.parse(
        localStorage.getItem('saved_users') || '[]'
      )
      const exists = savedUsers.find((u) => u.id === user.id)

      if (!exists) {
        savedUsers.push(user)
        localStorage.setItem('saved_users', JSON.stringify(savedUsers))
      }

      navigate('/home')
      return { isError: false }
    } catch {
      return { isError: true }
    }
  }

  const [state, formAction, isPending] = useActionState(loginAction, {
    isError: false
  })

  return { isError: state.isError, formAction, isPending }
}
