import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '@/contexts/AuthContext';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useToast } from '@/hooks/use-toast';
import { Shield, Mail, Lock, User, Building } from 'lucide-react';
import { api, type Campus } from '@/lib/api';

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1";

export default function AdminRegister() {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    confirmPassword: '',
    firstName: '',
    lastName: '',
    campusId: '',
  });
  const [campuses, setCampuses] = useState<Campus[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const { login } = useAuth();
  const { toast } = useToast();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCampuses = async () => {
      try {
        const response = await api.getCampuses();
        setCampuses(response.data || []);
      } catch (error) {
        toast({
          title: 'Error',
          description: 'Failed to load campuses',
          variant: 'destructive',
        });
      }
    };

    fetchCampuses();
  }, [toast]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (formData.password !== formData.confirmPassword) {
      toast({
        title: 'Passwords do not match',
        description: 'Please make sure both passwords are the same.',
        variant: 'destructive',
      });
      return;
    }

    if (formData.password.length < 8) {
      toast({
        title: 'Password too short',
        description: 'Password must be at least 8 characters long.',
        variant: 'destructive',
      });
      return;
    }

    setIsLoading(true);
    try {
      // Register as admin using the admin-specific endpoint
      await fetch(`${API_URL}/auth/register-admin`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: formData.email,
          password: formData.password,
          first_name: formData.firstName,
          last_name: formData.lastName,
          campus_id: formData.campusId || undefined,
          institution: 'U-Genius Platform',
          department: 'Administration',
          level: 'Administrator',
        }),
      });

      toast({
        title: 'Registration successful!',
        description: 'Your admin account has been created. Please login.',
      });

      // Auto-login after successful registration
      await login(formData.email, formData.password, true);
      navigate('/admin/dashboard');
    } catch (error: any) {
      toast({
        title: 'Registration failed',
        description: error.response?.data?.message || 'Something went wrong. Please try again.',
        variant: 'destructive',
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]: e.target.value
    }));
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gold-50 to-white px-4">
      <div className="w-full max-w-md">
        <div className="bg-white rounded-2xl shadow-xl p-8">
          <div className="text-center mb-8">
            <div className="w-16 h-16 rounded-full bg-gradient-gold flex items-center justify-center mx-auto mb-4">
              <Shield className="w-8 h-8 text-primary" />
            </div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">Admin Registration</h1>
            <p className="text-gray-600">Create your administrator account</p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="firstName">First Name</Label>
                <div className="relative">
                  <User className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                  <Input
                    id="firstName"
                    name="firstName"
                    type="text"
                    placeholder="John"
                    value={formData.firstName}
                    onChange={handleChange}
                    required
                    className="pl-10"
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="lastName">Last Name</Label>
                <div className="relative">
                  <User className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                  <Input
                    id="lastName"
                    name="lastName"
                    type="text"
                    placeholder="Doe"
                    value={formData.lastName}
                    onChange={handleChange}
                    required
                    className="pl-10"
                  />
                </div>
              </div>
            </div>

            <div className="space-y-2">
              <Label htmlFor="email">Email</Label>
              <div className="relative">
                <Mail className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                <Input
                  id="email"
                  name="email"
                  type="email"
                  placeholder="admin@example.com"
                  value={formData.email}
                  onChange={handleChange}
                  required
                  className="pl-10"
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label htmlFor="campusId">Campus</Label>
              <div className="relative">
                <Building className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                <select
                  id="campusId"
                  name="campusId"
                  value={formData.campusId}
                  onChange={handleChange}
                  className="w-full px-10 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gold-500 focus:border-transparent"
                  required
                >
                  <option value="">Select a campus</option>
                  {campuses.map((campus) => (
                    <option key={campus.id} value={campus.id}>
                      {campus.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <div className="relative">
                  <Lock className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                  <Input
                    id="password"
                    name="password"
                    type="password"
                    placeholder="••••••••••"
                    value={formData.password}
                    onChange={handleChange}
                    required
                    className="pl-10"
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="confirmPassword">Confirm Password</Label>
                <div className="relative">
                  <Lock className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
                  <Input
                    id="confirmPassword"
                    name="confirmPassword"
                    type="password"
                    placeholder="••••••••••"
                    value={formData.confirmPassword}
                    onChange={handleChange}
                    required
                    className="pl-10"
                  />
                </div>
              </div>
            </div>

            <Button type="submit" className="w-full" variant="gold" disabled={isLoading}>
              {isLoading ? 'Creating Account...' : 'Create Admin Account'}
            </Button>
          </form>

          <div className="mt-6 text-center space-y-2">
            <p className="text-sm text-gray-600">
              Already have an admin account?{' '}
              <a href="/admin/login" className="text-gold-600 hover:text-gold-700 font-medium">
                Sign in here
              </a>
            </p>
            <p className="text-sm text-gray-600">
              <a href="/login" className="text-gray-600 hover:text-gray-700 font-medium">
                ← Back to Student Login
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
