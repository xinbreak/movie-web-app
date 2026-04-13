import { useActionState } from 'react'
import { validateEmail, validatePassword } from '../utils/validation'
import { registerRequest } from '../api/authService'
import { useNavigate } from 'react-router-dom'

export const useReg = () => {
  const navigate = useNavigate()

  const registerAction = async (_prevState: any, formData: FormData) => {
    const data = Object.fromEntries(formData.entries())

    const email = validateEmail(data.email as string)
    const password = validatePassword(data.password as string)

    if (!email || !password || !data.firstName || !data.lastName) {
      return { succes: true }
    }

    try {
      const response = await registerRequest(data)
      console.log('Data: ', response.data, 'Succes: ', response.success)
      navigate('/login')
      return { success: true }
    } catch (error: any) {
      return {
        success: false
      }
    }
  }

  const [state, formAction, isPending] = useActionState(registerAction, {
    succes: false
  })

  return {
    succes: state.succes,
    formAction,
    isLoading: isPending
  }
}
