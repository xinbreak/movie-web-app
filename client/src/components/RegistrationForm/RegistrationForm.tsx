import styles from './RegistrationForm.module.css'
import InputForm from '../UI/InputForm/InputForm'
import ButtonForm from '../UI/ButtonForm/ButtonForm'
import { Link } from 'react-router-dom'
import { useReg } from '../../hooks/useReg'

export default function RegistrationForm() {
  const { isError, formAction, isLoading } = useReg()

  return (
    <div className={styles.formCard}>
      <h1 className={styles.title}>REGISTRATION</h1>
      <form action={formAction} className={styles.inputGroup} noValidate>
        <InputForm
          label="USERNAME"
          name="username"
          placeholder="username"
          isError={isError}
        />
        <InputForm
          label="PASSWORD"
          name="password"
          placeholder="password"
          type="password"
          isError={isError}
        />
        <InputForm
          label="EMAIL"
          name="email"
          placeholder="email@email.com"
          type="email"
          isError={isError}
        />
        <InputForm
          label="REPEAT PASSWORD"
          name="passwordRepeat"
          placeholder="password"
          type="password"
          isError={isError}
        />
        <ButtonForm
          buttonName={isLoading ? 'SENDING...' : 'SIGN UP'}
          disabled={isLoading}
        />
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
