import { Route, Routes, Navigate } from 'react-router-dom'
import LoginPage from './pages/LoginPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" />} />
      <Route path="/login" element={<LoginPage />} />

      <Route path="*" element={<h1>404</h1>} />
    </Routes>
  )
}

export default App
