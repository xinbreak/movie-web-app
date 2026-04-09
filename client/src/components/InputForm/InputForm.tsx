import styles from './InputForm.module.css'

interface InputFormProps {
  label: string
  type?: string
  placeholder?: string
  value: string
  onChange: (value: string) => void
}

export default function InputForm({
  label,
  type = 'text',
  placeholder,
  value,
  onChange
}: InputFormProps) {
  return (
    <div className={styles.field}>
      <label className={styles.label}>{label}</label>
      <input
        type={type}
        placeholder={placeholder}
        className={styles.input}
        value={value}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  )
}
