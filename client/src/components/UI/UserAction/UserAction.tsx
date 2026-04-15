import { useState } from 'react'
import { Link } from 'react-router-dom'
import styles from './UserAction.module.css'
import notification from '../../../assets/icons/notification.svg'
import userProfile from '../../../assets/icons/userProfile.svg'

export const UserAction = () => {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <>
      <div className={styles.rightSection}>
        <div className={styles.desktopIcons}>
          <Link to="/notifications" onClick={() => setIsOpen(false)}>
            <img src={notification} alt="not" width="24" />
          </Link>
          <Link to="/profile" onClick={() => setIsOpen(false)}>
            <img src={userProfile} alt="user" width="24" />
          </Link>
        </div>

        <div className={styles.mobileInteractiveGroup}>
          <div className={`${styles.miniMenu} ${isOpen ? styles.active : ''}`}>
            <Link to="/notifications" onClick={() => setIsOpen(false)}>
              <img src={notification} alt="not" width="24" />
            </Link>
            <Link to="/profile" onClick={() => setIsOpen(false)}>
              <img src={userProfile} alt="user" width="24" />
            </Link>
          </div>
          <button
            className={`${styles.burger} ${isOpen ? styles.burgerOpen : ''}`}
            onClick={() => setIsOpen(!isOpen)}
          >
            <span></span>
          </button>
        </div>
      </div>
      {isOpen && (
        <div className={styles.overlay} onClick={() => setIsOpen(false)} />
      )}
    </>
  )
}
