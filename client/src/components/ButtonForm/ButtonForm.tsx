import styles from './ButtonForm.module.css'

interface ButtonFormProps {
  buttonName: string
  disabled?: boolean
}

export default function ButtonForm({ buttonName, disabled }: ButtonFormProps) {
  return (
    <button type="submit" className={styles.buttonForm} disabled={disabled}>
      {buttonName}
    </button>
  )
}
