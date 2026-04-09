import styles from './ButtonForm.module.css'

export default function ButtonForm({ buttonName }: { buttonName: string }) {
  return (
    <button type="submit" className={styles.buttonForm}>
      {buttonName}
    </button>
  )
}
