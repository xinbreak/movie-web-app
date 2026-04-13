import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginRequest } from '../api/authService'
import { validateEmail } from '../utils/validation'

interface FormState {
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
      return { success: false }
    }

    try {
      const response = await loginRequest(email, password)
      localStorage.setItem('accessToken', response.token) //настройка токена
      navigate('/home')
      return { success: true }
    } catch (error: any) {
      return {
        success: false
      }
    }
  }

  const [state, formAction, isPending] = useActionState(loginAction, {
    success: false
  })

  return { state, formAction, isPending }
}
