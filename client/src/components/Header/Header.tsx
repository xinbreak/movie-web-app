import styles from './Header.module.css'
import logo from '/logo.svg'
import { Link, useLocation } from 'react-router-dom'
import { Navigation } from '../UI/Navigation/Navigation'
import { UserAction } from '../UI/UserAction/UserAction'
import SearchBar from '../UI/SearchBar/SearchBar'

export default function Header() {
  const location = useLocation()
  const isSearchPage = location.pathname === '/search'

  return (
    <header className={styles.header}>
      <div
        className={`${styles.navBar} ${
          isSearchPage ? styles.navBarSearch : ''
        }`}
      >
        <Link to="/home" className={styles.logo}>
          <img src={logo} alt="Logo" />
        </Link>

        {isSearchPage ? <SearchBar /> : <Navigation />}
      </div>

      <div className={styles.userActionWrapper}>
        <UserAction />
      </div>
    </header>
  )
}
