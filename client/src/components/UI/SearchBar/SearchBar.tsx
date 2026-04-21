import { useState, useEffect, type ChangeEvent } from 'react'
import styles from './SearchBar.module.css'
import searchIcon from '../../../assets/icons/Search.svg'

export default function SearchBar() {
  const [query, setQuery] = useState<string>('')

  useEffect(() => {
    if (!query.trim()) return

    const handler = setTimeout(() => {
      fetchData(query)
    }, 800)

    return () => clearTimeout(handler)
  }, [query])

  const fetchData = (value: string) => {
    console.log(value)
  }

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value)
  }

  return (
    <div className={styles.searchBar}>
      <img src={searchIcon} alt="Search" className={styles.searchIcon} />
      <input
        type="text"
        className={styles.searchInput}
        placeholder="Search any movie or serie"
        value={query}
        onChange={handleChange}
      />
    </div>
  )
}
