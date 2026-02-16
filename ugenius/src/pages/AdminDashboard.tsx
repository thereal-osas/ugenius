import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '@/contexts/AuthContext';
import { api, type User, type Campus } from '@/lib/api';
import { Button } from '@/components/ui/button';
import { useToast } from '@/hooks/use-toast';
import { UserPlus, Users, Building } from 'lucide-react';

export default function AdminDashboard() {
  const { user } = useAuth();
  const [users, setUsers] = useState<User[]>([]);
  const [campuses, setCampuses] = useState<Campus[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [promotingUserId, setPromotingUserId] = useState<string | null>(null);
  const { toast } = useToast();

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const [usersResponse, campusesResponse] = await Promise.all([
          api.getCampusUsers(),
          api.getCampuses()
        ]);
        setUsers(usersResponse.data || []);
        setCampuses(campusesResponse.data || []);
      } catch {
        toast({
          title: 'Error fetching data',
          description: 'Could not load the data.',
          variant: 'destructive',
        });
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, [toast]);

  const handlePromoteUser = async (userId: string, campusId: string) => {
    setPromotingUserId(userId);
    try {
      await api.promoteUser({ user_id: userId, campus_id: campusId });
      toast({
        title: 'User promoted successfully!',
        description: 'The user has been promoted to campus admin.',
      });
      
      // Refresh users list
      const response = await api.getCampusUsers();
      setUsers(response.data || []);
    } catch (error) {
      toast({
        title: 'Failed to promote user',
        description: error instanceof Error ? error.message : 'An error occurred',
        variant: 'destructive',
      });
    } finally {
      setPromotingUserId(null);
    }
  };

  const isAdmin = user?.role === 'campus_admin';
  const isSuperAdmin = user?.role === 'super_admin';

  const getUserListTitle = () => {
    if (isSuperAdmin) {
      return 'All Users';
    }
    if (isAdmin && user?.campus?.name) {
      return `${user.campus.name} Users`;
    }
    return 'Users';
  };

  const getUserListDescription = () => {
    if (isSuperAdmin) {
      return 'A list of all registered users across all campuses.';
    }
    return 'A list of all registered users in your campus.';
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8 flex justify-between items-center">
          <div className="flex items-center">
            <Users className="h-6 w-6 text-gold-600 mr-3" />
            <div>
              <h1 className="text-2xl font-bold text-gray-900">Admin Dashboard</h1>
              <p className="text-sm text-gray-500">
                Welcome, {user?.first_name}. You are a {user?.role.replace('_', ' ')}.
              </p>
            </div>
          </div>
          {isSuperAdmin && (
            <Button asChild variant="default">
              <Link to="/admin/create-admin">
                <UserPlus className="h-4 w-4 mr-2" />
                Create New Admin
              </Link>
            </Button>
          )}
        </div>
      </header>
      <main className="py-10">
        <div className="max-w-7xl mx-auto sm:px-6 lg:px-8">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg">
            <div className="px-4 py-5 sm:px-6">
              <div className="flex items-center">
                <Building className="h-5 w-5 text-gold-600 mr-2" />
                <h3 className="text-lg leading-6 font-medium text-gray-900">
                  {getUserListTitle()}
                </h3>
              </div>
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                {getUserListDescription()}
              </p>
            </div>
            <div className="border-t border-gray-200">
              {isLoading ? (
                <p className="p-6 text-center">Loading users...</p>
              ) : (
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Phone</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Campus</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Department</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Level</th>
                        <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Joined</th>
                        {isSuperAdmin && <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>}
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {users.map((u) => (
                        <tr key={u.id}>
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{u.first_name} {u.last_name}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{u.email}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{u.phone || 'N/A'}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            <span className={`px-2 py-1 text-xs font-medium rounded-full ${
                              u.role === 'super_admin' ? 'bg-purple-100 text-purple-800' :
                              u.role === 'campus_admin' ? 'bg-blue-100 text-blue-800' :
                              'bg-green-100 text-green-800'
                            }`}>
                              {u.role.replace('_', ' ')}
                            </span>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{u.campus?.name || 'N/A'}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{u.department || 'N/A'}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{u.level || 'N/A'}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {new Date(u.created_at).toLocaleDateString()}
                          </td>
                          {isSuperAdmin && (
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {u.role === 'student' && (
                                <select
                                  className="text-xs border border-gray-300 rounded px-2 py-1 mr-2"
                                  onChange={(e) => {
                                    if (e.target.value) {
                                      handlePromoteUser(u.id, e.target.value);
                                    }
                                  }}
                                  disabled={promotingUserId === u.id}
                                  value=""
                                >
                                  <option value="" disabled>Promote</option>
                                  {campuses.map((campus) => (
                                    <option key={campus.id} value={campus.id}>
                                      Promote to {campus.name}
                                    </option>
                                  ))}
                                </select>
                              )}
                              {u.role !== 'student' && (
                                <span className="text-xs text-gray-400">Already admin</span>
                              )}
                            </td>
                          )}
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              )}
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
