import styles from './FooterButton.module.css'
import { Link } from 'react-router-dom'

interface FooterButtonProps {
  linkTo: string
  footerIcon: string
  isActive: boolean
}

export default function FooterButton({
  linkTo,
  footerIcon,
  isActive
}: FooterButtonProps) {
  return (
    <button className={styles.footerButton}>
      <Link to={linkTo}>
        <img
          src={footerIcon}
          className={isActive ? styles.footerIconActive : styles.footerIcon}
        />
      </Link>
    </button>
  )
}
