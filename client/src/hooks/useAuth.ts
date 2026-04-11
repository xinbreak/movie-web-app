import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginRequest } from '../api/authService'
import { validateEmail } from '../utils/validation'

interface FormState {
  error: string | null
  success: boolean
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
      return { error: 'Invalid email', success: false }
    }

    try {
      const response = await loginRequest(email, password)
      localStorage.setItem('accessToken', response.token) //настройка токена
      navigate('/home')
      return { error: null, success: true }
    } catch (err: any) {
      return {
        error: 'Invalid password',
        success: false
      }
    }
  }

  const [state, formAction, isPending] = useActionState(loginAction, {
    error: null,
    success: false
  })

  return { state, formAction, isPending }
}
