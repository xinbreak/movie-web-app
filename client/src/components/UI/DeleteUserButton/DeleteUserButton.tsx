import styles from './DeleteUserButton.module.css'
import type { User } from '../../../api/authService'
import { useNavigate } from 'react-router-dom'
import { deleteUser } from '../../../api/deleteUser'

interface DeleteUserProps {
  userData: User
}

export default function DeleteUserButton({ userData }: DeleteUserProps) {
  const navigate = useNavigate()

  const handleClick = async () => {
    if (!window.confirm('Are you sure you want to delete your account?')) return

    try {
      await deleteUser(userData)
      localStorage.removeItem('current_user')
      localStorage.removeItem('isAuthorized')
      const savedUsersRaw = localStorage.getItem('saved_users')
      let remainingUsers: User[] = []

      if (savedUsersRaw) {
        const savedUsers: User[] = JSON.parse(savedUsersRaw)
        remainingUsers = savedUsers.filter((u) => u.id !== userData.id)

        localStorage.setItem('saved_users', JSON.stringify(remainingUsers))
      }

      remainingUsers.length === 0 ? navigate('/login') : navigate('/profiles')
    } catch (error) {
      console.error('Delete error:', error)
    }
  }

  return (
    <button className={styles.deleteUserButton} onClick={handleClick}>
      DELETE ACCOUNT
    </button>
  )
}
