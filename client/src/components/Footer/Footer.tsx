import styles from './Footer.module.css'
import FooterButton from '../UI/FooterButton/FooterButton'
import { useLocation } from 'react-router-dom'
import download from '../../assets/icons/Download.svg'
import favorites from '../../assets/icons/Favorites.svg'
import home from '../../assets/icons/Home.svg'
import search from '../../assets/icons/Search.svg'

export default function Footer() {
  const location = useLocation()

  return (
    <footer className={styles.footer}>
      <div className={styles.userAction}>
        <FooterButton
          linkTo="/home"
          footerIcon={home}
          isActive={location.pathname === '/home'}
        />
        <FooterButton
          linkTo="/search"
          footerIcon={search}
          isActive={location.pathname === '/search'}
        />
        <FooterButton
          linkTo="/favorites"
          footerIcon={favorites}
          isActive={location.pathname === '/favorites'}
        />
        <FooterButton
          linkTo="/download"
          footerIcon={download}
          isActive={location.pathname === '/download'}
        />
      </div>
    </footer>
  )
}
