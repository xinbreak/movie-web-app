import { Link } from 'react-router-dom'
import styles from './Navigation.module.css'

export const Navigation = () => (
  <ul className={styles.mainNav}>
    <li>
      <Link to="/movies">Movies</Link>
    </li>
    <li>
      <Link to="/series">Series</Link>
    </li>
  </ul>
)
