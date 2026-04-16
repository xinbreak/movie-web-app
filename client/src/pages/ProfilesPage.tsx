import { useNavigate } from 'react-router-dom'
import styles from '../styles/ProfilesPage.module.css'
import AddProfile from '../components/UI/AddProfile/AddProfile'
import ProfileButton from '../components/UI/ProfileButton/ProfileButton'
import EditButton from '../components/UI/EditProfileButton/EditProfileButton'
import type { User } from '../api/authService'

export default function ProfilesPage() {
  const navigate = useNavigate()
  const users: User[] = JSON.parse(
    localStorage.getItem('saved_users') || '[]'
  ).slice(0, 4)

  return (
    <div className={styles.pageStyle}>
      <button className={styles.closeButtonStyle} onClick={() => navigate(-1)}>
        ✕
      </button>

      <h1>Change profile</h1>

      <div className={styles.profilesContainer}>
        {users.map((user) => (
          <ProfileButton
            key={user.id}
            avatar={user.avatar_url}
            username={user.username}
          />
        ))}
        <AddProfile />
      </div>
      <EditButton />
    </div>
  )
}
