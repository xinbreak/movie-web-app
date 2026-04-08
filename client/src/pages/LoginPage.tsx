import LoginForm from '../components/LoginForm/LoginForm'

export default function LoginPage() {
  const pageStyle: React.CSSProperties = {
    minHeight: '100vh',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '20px'
  }

  return (
    <div style={pageStyle}>
      <LoginForm />
    </div>
  )
}
