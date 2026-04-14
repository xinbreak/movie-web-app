import { Link } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'
import ButtonForm from '../UI/ButtonForm/ButtonForm'
import InputForm from '../UI/InputForm/InputForm'
import styles from './LoginForm.module.css'

export default function LoginForm() {
  const { isError, formAction, isPending } = useAuth()

  return (
    <div className={styles.formCard}>
      <h1 className={styles.title}>LOGIN</h1>
      <form action={formAction} className={styles.inputGroup} noValidate>
        <InputForm
          label="EMAIL"
          name="email"
          placeholder="email@email.com"
          type="email"
          isError={isError}
        />
        <InputForm
          label="PASSWORD"
          name="password"
          placeholder="password"
          type="password"
          isError={isError}
        />
        <ButtonForm
          buttonName={isPending ? 'CHECKING...' : 'SIGN IN'}
          disabled={isPending}
        />
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
