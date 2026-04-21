import { useNavigate } from 'react-router-dom'
import styles from '../styles/ProfilesPage.module.css'
import AddProfile from '../components/UI/AddProfileButton/AddProfileButton'
import ProfileButton from '../components/UI/ProfileButton/ProfileButton'
import type { User } from '../api/authService'
import DeleteUserButton from '../components/UI/DeleteUserButton/DeleteUserButton'

export default function ProfilesPage() {
  const navigate = useNavigate()
  const users: User[] = JSON.parse(
    localStorage.getItem('saved_users') || '[]'
  ).slice(0, 4)
  const currentUser: User = JSON.parse(
    localStorage.getItem('current_user') || '[]'
  )

  return (
    <div className={styles.pageStyle}>
      <button className={styles.closeButtonStyle} onClick={() => navigate(-1)}>
        ✕
      </button>

      <h1>Change profile</h1>

      <div className={styles.profilesContainer}>
        {users.map((user) => (
          <ProfileButton key={user.id} user={user} />
        ))}
        <AddProfile />
      </div>
      <DeleteUserButton userData={currentUser} />
    </div>
  )
}
