import { Link } from 'react-router-dom'
import ButtonForm from '../ButtonForm/ButtonForm'
import InputForm from '../InputForm/InputForm'
import { useAuth } from '../../hooks/useAuth'
import styles from './LoginForm.module.css'

export default function LoginForm() {
  const { state, formAction, isPending } = useAuth()

  return (
    <div className={styles.formCard}>
      <h1 className={styles.title}>LOGIN</h1>

      <form action={formAction} className={styles.inputGroup} noValidate>
        <InputForm
          label="EMAIL"
          name="email"
          placeholder="email@email.com"
          type="email"
          isError={!!state.success}
        />

        <InputForm
          label="PASSWORD"
          name="password"
          placeholder="password"
          type="password"
          isError={!!state.success}
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
