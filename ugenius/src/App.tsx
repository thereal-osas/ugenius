import { Routes, Route } from 'react-router-dom'
import { Toaster } from '@/components/ui/toaster'
import { AuthProvider } from '@/contexts/AuthContext'
import ProtectedRoute  from '@/components/ProtectedRoute'
import EntryLoading from '@/components/EntryLoading'
import ScrollToTop from '@/components/ScrollToTop'
import IntroAnimation from '@/components/IntroAnimation'
import { useState } from 'react'
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
import AdminLogin from '@/pages/AdminLogin'
import AdminRegister from '@/pages/AdminRegister'
import DashboardUnavailable from '@/pages/DashboardUnavailable'
import AdminDashboard from '@/pages/AdminDashboard'
import CreateAdmin from '@/pages/CreateAdmin'
import ReadingHours from '@/pages/ReadingHours'
import Goals from '@/pages/Goals'
import StudyGroups from '@/pages/StudyGroups'
import NotFound from '@/pages/NotFound'

function App() {
  const [showIntro, setShowIntro] = useState(false); // Set to false to disable intro

  const handleIntroComplete = () => {
    setShowIntro(false);
  };

  return (
    <AuthProvider>
      <ScrollToTop />
      {showIntro && <IntroAnimation onComplete={handleIntroComplete} />}
      <EntryLoading />
      <Routes>
        {/* Public routes */}
        <Route path="/" element={<Index />} />
        <Route path="/about" element={<About />} />
        <Route path="/benefits" element={<Benefits />} />
        <Route path="/resources" element={<Resources />} />
        <Route path="/testimonials" element={<Testimonials />} />
        {/* <Route path="/contact" element={<Contact />} /> */}
        <Route path="/join" element={<Join />} />
        <Route path="/thank-you" element={<ThankYou />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/admin/login" element={<AdminLogin />} />
        <Route path="/admin/register" element={<AdminRegister />} />

        {/* Protected routes */}
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <DashboardUnavailable />
            </ProtectedRoute>
          }
        />
        <Route
          path="/admin/dashboard"
          element={
            <ProtectedRoute adminOnly>
              <AdminDashboard />
            </ProtectedRoute>
          }
        />
        <Route
          path="/admin/create-admin"
          element={
            <ProtectedRoute adminOnly>
              <CreateAdmin />
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
