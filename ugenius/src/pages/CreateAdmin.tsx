import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '@/contexts/AuthContext';
import { api } from '@/lib/api';
import type {Campus} from '@/lib/api'
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useToast } from '@/hooks/use-toast';
import { ArrowLeft, UserPlus } from 'lucide-react';

export default function CreateAdmin() {
  const [email, setEmail] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [password, setPassword] = useState('');
  const [campusId, setCampusId] = useState('');
  const [campuses, setCampuses] = useState<Campus[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const { user } = useAuth();
  const navigate = useNavigate();
  const { toast } = useToast();

  useEffect(() => {
    const fetchCampuses = async () => {
      try {
        const response = await api.getCampuses();
        if (response.data) {
          setCampuses(response.data.filter((campus: Campus) => campus.is_active));
        }
      } catch (error) {
        console.error('Failed to fetch campuses:', error);
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

    if (!campusId) {
      toast({
        title: 'Campus required',
        description: 'Please select a campus for the admin.',
        variant: 'destructive',
      });
      return;
    }

    if (password.length < 8) {
      toast({
        title: 'Password too short',
        description: 'Password must be at least 8 characters long.',
        variant: 'destructive',
      });
      return;
    }

    setIsLoading(true);

    try {
      await api.createAdmin({
        email,
        first_name: firstName,
        last_name: lastName,
        password,
        campus_id: campusId,
      });
      
      toast({
        title: 'Admin created successfully!',
        description: `${firstName} ${lastName} is now a campus admin.`,
      });
      
      navigate('/admin/dashboard');
    } catch (error) {
      toast({
        title: 'Failed to create admin',
        description: error instanceof Error ? error.message : 'An error occurred',
        variant: 'destructive',
      });
    } finally {
      setIsLoading(false);
    }
  };

  // Only allow super admins to access this page
  if (user?.role !== 'super_admin') {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-gray-900 mb-4">Access Denied</h1>
          <p className="text-gray-600 mb-6">You don't have permission to access this page.</p>
          <Button asChild>
            <Link to="/admin/dashboard">Back to Dashboard</Link>
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
          <div className="flex items-center">
            <Button variant="ghost" asChild className="mr-4">
              <Link to="/admin/dashboard">
                <ArrowLeft className="h-4 w-4 mr-2" />
                Back to Dashboard
              </Link>
            </Button>
            <div>
              <h1 className="text-2xl font-bold text-gray-900">Create Campus Admin</h1>
              <p className="text-sm text-gray-500">Add a new administrator for a campus</p>
            </div>
          </div>
        </div>
      </header>

      <main className="py-10">
        <div className="max-w-2xl mx-auto sm:px-6 lg:px-8">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg">
            <div className="px-4 py-5 sm:px-6">
              <div className="flex items-center">
                <UserPlus className="h-6 w-6 text-gold-600 mr-3" />
                <h3 className="text-lg leading-6 font-medium text-gray-900">
                  Admin Information
                </h3>
              </div>
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                Fill in the details below to create a new campus administrator account.
              </p>
            </div>
            <form onSubmit={handleSubmit} className="px-4 py-5 sm:px-6 space-y-6">
              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="firstName">First Name</Label>
                  <Input
                    id="firstName"
                    value={firstName}
                    onChange={(e) => setFirstName(e.target.value)}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="lastName">Last Name</Label>
                  <Input
                    id="lastName"
                    value={lastName}
                    onChange={(e) => setLastName(e.target.value)}
                    required
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="email">Email Address</Label>
                <Input
                  id="email"
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="admin@university.edu"
                  required
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="campus">Campus Assignment</Label>
                <select
                  id="campus"
                  value={campusId}
                  onChange={(e) => setCampusId(e.target.value)}
                  required
                  className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-gold-500 focus:border-gold-500 sm:text-sm"
                >
                  <option value="">Select a campus</option>
                  {campuses.map((campus) => (
                    <option key={campus.id} value={campus.id}>
                      {campus.name} ({campus.code})
                    </option>
                  ))}
                </select>
                <p className="text-sm text-gray-500">
                  This admin will only have access to users and data from the selected campus.
                </p>
              </div>

              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  placeholder="••••••••"
                  required
                />
              </div>

              <div className="pt-4">
                <Button type="submit" className="w-full" variant="gold" disabled={isLoading}>
                  {isLoading ? 'Creating Admin...' : 'Create Campus Admin'}
                </Button>
              </div>
            </form>
          </div>
        </div>
      </main>
    </div>
  );
}
