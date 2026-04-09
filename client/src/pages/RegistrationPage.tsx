import RegistrationForm from '../components/RegistrationForm/RegistrationForm'

export default function RegistrationPage() {
  const pageStyle: React.CSSProperties = {
    minHeight: '100vh',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '20px'
  }

  return (
    <div style={pageStyle}>
      <RegistrationForm />
    </div>
  )
}
