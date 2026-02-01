import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { api } from '@/lib/api';
import type { Goal } from '@/lib/api';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { useToast } from '@/hooks/use-toast';

const GOAL_TYPES = [
  { value: 'reading_hours', label: 'Reading Hours' },
  { value: 'submissions', label: 'Submissions' },
  { value: 'streak', label: 'Daily Streak' },
  { value: 'custom', label: 'Custom' },
];

export default function Goals() {
  const { toast } = useToast();
  const [goals, setGoals] = useState<Goal[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    goal_type: 'reading_hours',
    target_value: 10,
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    fetchGoals();
  }, []);

  const fetchGoals = async () => {
    try {
      const response = await api.getGoals();
      if (response.data) {
        setGoals(response.data);
      }
    } catch (error) {
      console.error('Failed to fetch goals:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      await api.createGoal(formData);
      toast({
        title: 'Goal created!',
        description: 'Start working towards your new goal.',
      });
      setShowForm(false);
      setFormData({ title: '', description: '', goal_type: 'reading_hours', target_value: 10 });
      fetchGoals();
    } catch (error) {
      toast({
        title: 'Failed to create goal',
        description: error instanceof Error ? error.message : 'An error occurred',
        variant: 'destructive',
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  const getProgress = (goal: Goal) => Math.min((goal.current_value / goal.target_value) * 100, 100);

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex justify-between items-center">
          <div className="flex items-center gap-4">
            <Link to="/dashboard" className="text-gray-500 hover:text-gray-700">← Back</Link>
            <h1 className="text-2xl font-bold text-gray-900">Goals</h1>
          </div>
          <Button variant="gold" onClick={() => setShowForm(true)}>New Goal</Button>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {showForm && (
          <div className="bg-white rounded-xl shadow-sm p-6 mb-8">
            <h2 className="text-lg font-semibold mb-4">Create New Goal</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="title">Goal Title</Label>
                <Input
                  id="title"
                  placeholder="e.g., Read 20 hours this week"
                  value={formData.title}
                  onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                  required
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="description">Description (optional)</Label>
                <Textarea
                  id="description"
                  placeholder="Describe your goal..."
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                />
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="goal_type">Goal Type</Label>
                  <select
                    id="goal_type"
                    value={formData.goal_type}
                    onChange={(e) => setFormData({ ...formData, goal_type: e.target.value })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-gold-500"
                  >
                    {GOAL_TYPES.map((type) => (
                      <option key={type.value} value={type.value}>{type.label}</option>
                    ))}
                  </select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="target_value">Target Value</Label>
                  <Input
                    id="target_value"
                    type="number"
                    min="1"
                    value={formData.target_value}
                    onChange={(e) => setFormData({ ...formData, target_value: parseInt(e.target.value) })}
                    required
                  />
                </div>
              </div>
              <div className="flex gap-4">
                <Button type="submit" variant="gold" disabled={isSubmitting}>
                  {isSubmitting ? 'Creating...' : 'Create Goal'}
                </Button>
                <Button type="button" variant="outline" onClick={() => setShowForm(false)}>Cancel</Button>
              </div>
            </form>
          </div>
        )}

        {isLoading ? (
          <div className="text-center py-8">Loading...</div>
        ) : goals.length === 0 ? (
          <div className="bg-white rounded-xl shadow-sm p-8 text-center">
            <p className="text-gray-500 mb-4">No goals yet. Set your first goal to start tracking progress!</p>
            <Button variant="gold" onClick={() => setShowForm(true)}>Create Your First Goal</Button>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {goals.map((goal) => (
              <div key={goal.id} className={`bg-white rounded-xl shadow-sm p-6 ${goal.is_completed ? 'border-2 border-green-500' : ''}`}>
                <div className="flex justify-between items-start mb-4">
                  <div>
                    <h3 className="font-semibold text-lg">{goal.title}</h3>
                    {goal.description && <p className="text-gray-500 text-sm mt-1">{goal.description}</p>}
                  </div>
                  {goal.is_completed && <span className="text-2xl">🎉</span>}
                </div>
                <div className="mb-2 flex justify-between text-sm">
                  <span className="text-gray-500">{goal.goal_type.replace('_', ' ')}</span>
                  <span className="font-medium">{goal.current_value} / {goal.target_value}</span>
                </div>
                <div className="w-full bg-gray-200 rounded-full h-3">
                  <div
                    className={`h-3 rounded-full ${goal.is_completed ? 'bg-green-500' : 'bg-gold-500'}`}
                    style={{ width: `${getProgress(goal)}%` }}
                  ></div>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}

