import Header from '../components/Header/Header'
import Footer from '../components/Footer/Footer'
import styles from '../styles/HomePage.module.css'

export default function HomePage() {
  return (
    <div className={styles.pageStyle}>
      <Header />
      <Footer />
    </div>
  )
}
