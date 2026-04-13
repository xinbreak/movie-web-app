import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginRequest } from '../api/authService'
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
      const response = await loginRequest(email, password)
      localStorage.setItem('accessToken', response.token)
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
