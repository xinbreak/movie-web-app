import styles from './Header.module.css'
import logo from '/logo.svg'
import { Link } from 'react-router-dom'
import { Navigation } from '../UI/Navigation/Navigation'
import { UserAction } from '../UI/UserAction/UserAction'

export default function Header() {
  return (
    <header>
      <div className={styles.navBar}>
        <Link to="/home">
          <img src={logo} alt="Logo" />
        </Link>
        <Navigation />
      </div>
      <UserAction />
    </header>
  )
}
