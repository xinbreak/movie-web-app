import styles from './RegistrationForm.module.css'
import InputForm from '../InputForm/InputForm'
import ButtonForm from '../ButtonForm/ButtonForm'
import { Link } from 'react-router-dom'
import { useReg } from '../../hooks/useReg'

export default function RegistrationForm() {
  const { succes, formAction, isLoading } = useReg()

  return (
    <div className={styles.formCard}>
      <h1 className={styles.title}>REGISTRATION</h1>

      <form action={formAction} className={styles.inputGroup} noValidate>
        <InputForm
          label="FIRST NAME"
          name="firstName"
          placeholder="First name"
          isError={succes}
        />

        <InputForm
          label="LAST NAME"
          name="lastName"
          placeholder="Last name"
          isError={succes}
        />

        <InputForm
          label="EMAIL"
          name="email"
          placeholder="email@email.com"
          type="email"
          isError={succes}
        />

        <InputForm
          label="PASSWORD"
          name="password"
          placeholder="password"
          type="password"
          isError={succes}
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
