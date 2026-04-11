import styles from './InputForm.module.css'

interface InputFormProps {
  label: string
  name: string
  type?: string
  placeholder?: string
  isError?: boolean
}

export default function InputForm({
  label,
  name,
  type = 'text',
  placeholder,
  isError
}: InputFormProps) {
  return (
    <div className={styles.field}>
      <label className={styles.label} htmlFor={name}>
        {label}
      </label>
      <input
        id={name}
        name={name}
        type={type}
        placeholder={isError ? `Invalid ` + type : placeholder}
        className={`${styles.input} ${isError ? styles.inputError : ''}`}
      />
    </div>
  )
}
