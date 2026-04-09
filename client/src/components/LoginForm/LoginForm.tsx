import { Link, useNavigate } from 'react-router-dom'
import { useState } from 'react'
import ButtonForm from '../ButtonForm/ButtonForm'
import InputForm from '../InputForm/InputForm'
import styles from './LoginForm.module.css'

export default function LoginForm() {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const handleLogin = () => {
    //доработать
    if (email == 'admin@mail.ru' && password == '1234') {
      navigate('/home')
    }
  }

  return (
    <div className={styles.formCard}>
      <h2 className={styles.title}>LOGIN</h2>

      <form onSubmit={handleLogin} className={styles.inputGroup}>
        <InputForm
          label="EMAIL"
          placeholder="email@email.com"
          type="email"
          value={email}
          onChange={setEmail}
        />

        <InputForm
          label="PASSWORD"
          placeholder="password"
          type="password"
          value={password}
          onChange={setPassword}
        />

        <ButtonForm buttonName="SIGN IN" />
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
