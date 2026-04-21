import { Route, Routes, Navigate } from 'react-router-dom'
import { ProtectedRoute } from './components/ProtectedRoute/ProtectedRoute'
import LoginPage from './pages/LoginPage'
import RegistrationPage from './pages/RegistrationPage'
import HomePage from './pages/HomePage'
import ProfilesPage from './pages/ProfilesPage'
import SearchPage from './pages/SearchPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/home" />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/registration" element={<RegistrationPage />} />
      <Route path="/home" element={<HomePage />} />
      <Route path="/search" element={<SearchPage />} />

      <Route element={<ProtectedRoute />}>
        <Route path="/profiles" element={<ProfilesPage />} />
      </Route>

      <Route path="*" element={<h1>404</h1>} />
    </Routes>
  )
}

export default App
