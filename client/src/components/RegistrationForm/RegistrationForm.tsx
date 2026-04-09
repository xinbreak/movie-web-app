import styles from './RegistrationForm.module.css'
import InputForm from '../InputForm/InputForm'
import ButtonForm from '../ButtonForm/ButtonForm'
import { useState } from 'react'
import { Link } from 'react-router-dom'

export default function RegistrationForm() {
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  return (
    <div className={styles.formCard}>
      <h1 className={styles.title}>REGISTRATION</h1>

      <form className={styles.inputGroup}>
        <InputForm
          label="FIRST NAME"
          placeholder="First name"
          value={firstName}
          onChange={setFirstName}
        />

        <InputForm
          label="LAST NAME"
          placeholder="Last name"
          value={lastName}
          onChange={setLastName}
        />

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

        <ButtonForm buttonName="SIGN UP" />
      </form>

      <div className={styles.footer}>
        <span>Already have an account?</span>
        <Link to="/login" className={styles.linkWrapper}>
          <button className={styles.signInBtn}>Sign In</button>
        </Link>
      </div>
    </div>
  )
}
