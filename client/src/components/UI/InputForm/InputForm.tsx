import { useState } from 'react'
import styles from './InputForm.module.css'
import showIcon from '../../../assets/icons/showPassword.png'
import hideIcon from '../../../assets/icons/hidePassword.png'

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
  const [showPassword, setShowPassword] = useState(false)

  const currentType =
    name === 'password' ? (showPassword ? 'text' : 'password') : type

  return (
    <div className={styles.field}>
      <label className={styles.label} htmlFor={name}>
        {label}
      </label>
      <div className={styles.inputWrapper}>
        <input
          id={name}
          name={name}
          type={currentType}
          placeholder={isError ? `Invalid ` + name : placeholder}
          className={`${styles.input} ${isError ? styles.inputError : ''}`}
        />
        {name === 'password' && (
          <button
            type="button"
            className={styles.toggleBtn}
            onClick={() => setShowPassword(!showPassword)}
            tabIndex={-1}
          >
            <img
              src={showPassword ? hideIcon : showIcon}
              className={styles.icon}
            />
          </button>
        )}
      </div>
    </div>
  )
}
