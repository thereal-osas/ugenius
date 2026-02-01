import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { useToast } from '@/hooks/use-toast';

interface StudyGroup {
  id: string;
  name: string;
  description: string;
  subject: string;
  max_members: number;
  member_count: number;
  is_member: boolean;
  created_at: string;
}

export default function StudyGroups() {
  const { toast } = useToast();
  const [groups, setGroups] = useState<StudyGroup[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    subject: '',
    max_members: 10,
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    fetchGroups();
  }, []);

  const fetchGroups = async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'}/study-groups`, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
        },
      });
      const data = await response.json();
      if (data.data) {
        setGroups(data.data);
      }
    } catch (error) {
      console.error('Failed to fetch study groups:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'}/study-groups`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
        },
        body: JSON.stringify(formData),
      });
      if (!response.ok) throw new Error('Failed to create group');
      toast({ title: 'Study group created!', description: 'Invite your friends to join.' });
      setShowForm(false);
      setFormData({ name: '', description: '', subject: '', max_members: 10 });
      fetchGroups();
    } catch (error) {
      toast({ title: 'Failed to create group', variant: 'destructive' });
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleJoinLeave = async (groupId: string, isMember: boolean) => {
    try {
      const endpoint = isMember ? 'leave' : 'join';
      const response = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'}/study-groups/${groupId}/${endpoint}`, {
        method: 'POST',
        headers: { 'Authorization': `Bearer ${localStorage.getItem('access_token')}` },
      });
      if (!response.ok) throw new Error('Failed');
      toast({ title: isMember ? 'Left group' : 'Joined group!' });
      fetchGroups();
    } catch (error) {
      toast({ title: 'Action failed', variant: 'destructive' });
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex justify-between items-center">
          <div className="flex items-center gap-4">
            <Link to="/dashboard" className="text-gray-500 hover:text-gray-700">← Back</Link>
            <h1 className="text-2xl font-bold text-gray-900">Study Groups</h1>
          </div>
          <Button variant="gold" onClick={() => setShowForm(true)}>Create Group</Button>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {showForm && (
          <div className="bg-white rounded-xl shadow-sm p-6 mb-8">
            <h2 className="text-lg font-semibold mb-4">Create Study Group</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="name">Group Name</Label>
                  <Input id="name" value={formData.name} onChange={(e) => setFormData({ ...formData, name: e.target.value })} required />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="subject">Subject</Label>
                  <Input id="subject" value={formData.subject} onChange={(e) => setFormData({ ...formData, subject: e.target.value })} required />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="description">Description</Label>
                <Textarea id="description" value={formData.description} onChange={(e) => setFormData({ ...formData, description: e.target.value })} />
              </div>
              <div className="space-y-2">
                <Label htmlFor="max_members">Max Members</Label>
                <Input id="max_members" type="number" min="2" max="50" value={formData.max_members} onChange={(e) => setFormData({ ...formData, max_members: parseInt(e.target.value) })} />
              </div>
              <div className="flex gap-4">
                <Button type="submit" variant="gold" disabled={isSubmitting}>{isSubmitting ? 'Creating...' : 'Create'}</Button>
                <Button type="button" variant="outline" onClick={() => setShowForm(false)}>Cancel</Button>
              </div>
            </form>
          </div>
        )}

        {isLoading ? (
          <div className="text-center py-8">Loading...</div>
        ) : groups.length === 0 ? (
          <div className="bg-white rounded-xl shadow-sm p-8 text-center">
            <p className="text-gray-500 mb-4">No study groups yet. Create one to start collaborating!</p>
            <Button variant="gold" onClick={() => setShowForm(true)}>Create First Group</Button>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {groups.map((group) => (
              <div key={group.id} className="bg-white rounded-xl shadow-sm p-6">
                <h3 className="font-semibold text-lg mb-2">{group.name}</h3>
                <p className="text-sm text-gold-600 mb-2">{group.subject}</p>
                <p className="text-gray-500 text-sm mb-4">{group.description}</p>
                <div className="flex justify-between items-center">
                  <span className="text-sm text-gray-500">{group.member_count}/{group.max_members} members</span>
                  <Button variant={group.is_member ? 'outline' : 'gold'} size="sm" onClick={() => handleJoinLeave(group.id, group.is_member)}>
                    {group.is_member ? 'Leave' : 'Join'}
                  </Button>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}

