import { useActionState } from 'react'
import { useNavigate } from 'react-router-dom'
import { validateEmail, validatePassword } from '../utils/validation'
import { registerRequest } from '../api/authService'

interface FormState {
  isError: boolean
}

export const useReg = () => {
  const navigate = useNavigate()

  const registerAction = async (
    _prevState: FormState,
    formData: FormData
  ): Promise<FormState> => {
    const data = Object.fromEntries(formData.entries())

    const emailValid = validateEmail(data.email as string)
    const passValid = validatePassword(data.password as string)

    if (!emailValid || !passValid || !data.firstName || !data.lastName) {
      return { isError: true }
    }

    try {
      await registerRequest(data)
      navigate('/login')
      return { isError: false }
    } catch {
      return { isError: true }
    }
  }

  const [state, formAction, isPending] = useActionState(registerAction, {
    isError: false
  })

  return { isError: state.isError, formAction, isLoading: isPending }
}
