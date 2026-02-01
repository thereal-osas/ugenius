import { Routes, Route } from 'react-router-dom'
import { Toaster } from '@/components/ui/toaster'
import { AuthProvider } from '@/contexts/AuthContext'
import ProtectedRoute from '@/components/ProtectedRoute'
import Index from '@/pages/Index'
import About from '@/pages/About'
import Benefits from '@/pages/Benefits'
import Resources from '@/pages/Resources'
import Testimonials from '@/pages/Testimonials'
import Contact from '@/pages/Contact'
import Join from '@/pages/Join'
import ThankYou from '@/pages/ThankYou'
import Login from '@/pages/Login'
import Register from '@/pages/Register'
import Dashboard from '@/pages/Dashboard'
import ReadingHours from '@/pages/ReadingHours'
import Goals from '@/pages/Goals'
import StudyGroups from '@/pages/StudyGroups'
import NotFound from '@/pages/NotFound'

function App() {
  return (
    <AuthProvider>
      <Routes>
        {/* Public routes */}
        <Route path="/" element={<Index />} />
        <Route path="/about" element={<About />} />
        <Route path="/benefits" element={<Benefits />} />
        <Route path="/resources" element={<Resources />} />
        <Route path="/testimonials" element={<Testimonials />} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/join" element={<Join />} />
        <Route path="/thank-you" element={<ThankYou />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Protected routes */}
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          }
        />
        <Route
          path="/reading-hours"
          element={
            <ProtectedRoute>
              <ReadingHours />
            </ProtectedRoute>
          }
        />
        <Route
          path="/reading-hours/new"
          element={
            <ProtectedRoute>
              <ReadingHours />
            </ProtectedRoute>
          }
        />
        <Route
          path="/goals"
          element={
            <ProtectedRoute>
              <Goals />
            </ProtectedRoute>
          }
        />
        <Route
          path="/goals/new"
          element={
            <ProtectedRoute>
              <Goals />
            </ProtectedRoute>
          }
        />
        <Route
          path="/study-groups"
          element={
            <ProtectedRoute>
              <StudyGroups />
            </ProtectedRoute>
          }
        />

        <Route path="*" element={<NotFound />} />
      </Routes>
      <Toaster />
    </AuthProvider>
  )
}

export default App
