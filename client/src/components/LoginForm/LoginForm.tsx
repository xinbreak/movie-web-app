import { Link, useNavigate } from 'react-router-dom'
import { useState } from 'react'
import styles from './LoginForm.module.css'

export default function LoginForm() {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const handleLogin = () => {
    navigate('/home')
  }

  return (
    <div className={styles.formCard}>
      <h2 className={styles.title}>LOGIN</h2>

      <form onSubmit={handleLogin} className={styles.inputGroup}>
        <div className={styles.field}>
          <label className={styles.label}>EMAIL</label>
          <input
            type="email"
            placeholder="email@email.com"
            className={styles.input}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>

        <div className={styles.field}>
          <label className={styles.label}>PASSWORD</label>
          <input
            type="password"
            placeholder="•••••••••••"
            className={styles.input}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <button type="submit" className={styles.loginBtn}>
          SIGN IN
        </button>
      </form>

      <div className={styles.footer}>
        <span>Don't have an account?</span>
        <Link to="/registration" className={styles.linkWrapper}>
          <button className={styles.signUpBtn}>SignUp</button>
        </Link>
      </div>
    </div>
  )
}
