import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginRequest, getUsers } from '../api/authService'
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
      const users = await getUsers()
      console.log('Список пользователей:', users)
      return { isError: true }
    }

    try {
      await loginRequest(email, password)
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
