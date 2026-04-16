import styles from './EditProfileButton.module.css'
import type { User } from '../../../api/authService'

export default function EditButton() {
  const user: User[] = JSON.parse(localStorage.getItem('current_user') || '[]')
  return <button className={styles.editButton}>Edit profile</button>
}
