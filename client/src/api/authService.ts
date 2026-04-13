export const loginRequest = async (email: string, password: string) => {
  await new Promise((resolve) => setTimeout(resolve, 800))

  if (email === 'admin@mail.ru' && password === '1234') {
    return { success: true, token: 'fake-jwt-token' }
  }

  throw new Error()
}

export const registerRequest = async (userData: Record<string, any>) => {
  await new Promise((resolve) => setTimeout(resolve, 800))

  if (userData) {
    return { success: true, data: userData }
  }

  throw new Error()
}
