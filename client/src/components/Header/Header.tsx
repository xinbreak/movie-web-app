import styles from './Header.module.css'
import logo from '/logo.svg'
import { Navigation } from '../UI/Navigation/Navigation'
import { UserAction } from '../UI/UserAction/UserAction'

export default function Header() {
  return (
    <header className={styles.header}>
      <div className={styles.navBar}>
        <img src={logo} alt="Logo" />
        <Navigation />
      </div>
      <UserAction />
    </header>
  )
}
