import type { User } from './authService'

const API_URL = import.meta.env.VITE_API_URL

export const uploadAvatar = async (file: File): Promise<User> => {
  const user = JSON.parse(localStorage.getItem('current_user') || '{}')

  // 1. Читаем файл в Base64
  const base64 = await new Promise<string>((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = () => resolve(reader.result as string)
    reader.onerror = reject
    reader.readAsDataURL(file)
  })

  // 2. Отправляем на бэк
  const response = await fetch(`${API_URL}/users/${user.id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      avatar_url: base64,
      username: user.username,
      password: user.password
    })
  })

  if (!response.ok) throw new Error('Update failed')

  // 3. ТАК КАК БЭК НЕ ВОЗВРАЩАЕТ ЮЗЕРА:
  // Создаем обновленный объект сами на основе старого + новый аватар
  const updatedUser = { ...user, avatar_url: base64 }

  // Сохраняем в localStorage
  localStorage.setItem('current_user', JSON.stringify(updatedUser))

  return updatedUser // Возвращаем именно объект пользователя
}
