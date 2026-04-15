import Header from '../components/Header/Header'
import Footer from '../components/Footer/Footer'

export default function HomePage() {
  const pageStyle: React.CSSProperties = {
    minHeight: '100vh',
    paddingTop: '60px',
    paddingBottom: '60px'
  }

  return (
    <div style={pageStyle}>
      <Header />
      <Footer />
    </div>
  )
}
