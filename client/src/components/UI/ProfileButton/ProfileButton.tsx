import { useNavigate } from 'react-router-dom'
import styles from './ProfileButton.module.css'
import defaultUser from '../../../assets/icons/DefaultUser.svg'
import type { User } from '../../../api/authService'

interface ProfileButtonProps {
  user: User
}

export default function ProfileButton({ user }: ProfileButtonProps) {
  const navigate = useNavigate()

  const profileImage = !user.avatar_url ? defaultUser : user.avatar_url

  const handleProfileClick = () => {
    localStorage.setItem('current_user', JSON.stringify(user))
    localStorage.setItem('isAuthorized', 'true')

    navigate('/home')
  }

  return (
    <div className={styles.addProfile}>
      <button className={styles.btn} onClick={handleProfileClick}>
        <img src={profileImage} alt={user.username} />
      </button>
      <span className={styles.label}>{user.username}</span>
    </div>
  )
}
