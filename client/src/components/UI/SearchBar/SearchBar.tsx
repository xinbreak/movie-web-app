import styles from './SearchBar.module.css'
import searchIcon from '../../../assets/icons/Search.svg'

export default function SearchBar() {
  return (
    <div className={styles.searchBar}>
      <img src={searchIcon} alt="Search" className={styles.searchIcon} />
      <input
        type="text"
        className={styles.searchInput}
        placeholder="Search any movie or serie"
      />
    </div>
  )
}
