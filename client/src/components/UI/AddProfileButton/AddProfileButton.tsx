import styles from './AddProfileButton.module.css'
import { Link } from 'react-router-dom'

export default function AddProfile() {
  return (
    <div className={styles.addProfile}>
      <Link to="/login">
        <button>+</button>
      </Link>
      <span className={styles.label}>Add new</span>
    </div>
  )
}
