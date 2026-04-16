import styles from './ProfileButton.module.css'
import defaultUser from '../../../assets/icons/DefaultUser.svg'
import { Link } from 'react-router-dom'

interface ProfileButtonProps {
  avatar: string | null
  username: string
}

export default function ProfileButton({
  avatar,
  username
}: ProfileButtonProps) {
  const profileImage = !avatar ? defaultUser : avatar

  return (
    <div className={styles.addProfile}>
      <Link to="/registration">
        <button>
          <img src={profileImage} />
        </button>
      </Link>
      <span className={styles.label}>{username}</span>
    </div>
  )
}
