import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { Lock, Home, User } from 'lucide-react';

export default function DashboardUnavailable() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gold-50 to-white px-4">
      <div className="w-full max-w-md text-center">
        <div className="bg-white rounded-2xl shadow-xl p-8">
          <div className="mx-auto w-16 h-16 bg-gold-100 rounded-full flex items-center justify-center mb-6">
            <Lock className="h-8 w-8 text-gold-600" />
          </div>
          
          <h1 className="text-3xl font-bold text-gray-900 mb-4">
            Dashboard Temporarily Unavailable
          </h1>
          
          <p className="text-gray-600 mb-8">
            The student dashboard is currently under maintenance. We're working on bringing you an enhanced experience!
          </p>
          
          <div className="space-y-4">
            <Button asChild className="w-full" variant="gold">
              <Link to="/">
                <Home className="h-4 w-4 mr-2" />
                Back to Home
              </Link>
            </Button>
            
            <Button asChild className="w-full" variant="outline">
              <Link to="/admin/login">
                <User className="h-4 w-4 mr-2" />
                Admin Login
              </Link>
            </Button>
          </div>
          
          <div className="mt-8 p-4 bg-gray-50 rounded-lg">
            <p className="text-sm text-gray-500">
              <strong>Administrators:</strong> You can access the admin dashboard using the Admin Login button above.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
