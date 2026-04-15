import styles from './FilmCard.module.css'

export default function FilmCard({ img, name }: { img: string; name: string }) {
  return (
    <div className={styles.filmCard}>
      <img src={img} alt="" />
      <h3>{name}</h3>
    </div>
  )
}
