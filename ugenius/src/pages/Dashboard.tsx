import { useEffect, useState } from 'react';
import { useAuth } from '@/contexts/AuthContext';
import { api } from '@/lib/api';
import type { ReadingHourStats, LeaderboardEntry, Goal, Achievement } from '@/lib/api';
import { Button } from '@/components/ui/button';
import { Link } from 'react-router-dom';

export default function Dashboard() {
  const { user, logout } = useAuth();
  const [stats, setStats] = useState<ReadingHourStats | null>(null);
  const [leaderboard, setLeaderboard] = useState<LeaderboardEntry[]>([]);
  const [goals, setGoals] = useState<Goal[]>([]);
  const [achievements, setAchievements] = useState<Achievement[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [statsRes, leaderboardRes, goalsRes, achievementsRes] = await Promise.all([
          api.getReadingHourStats(),
          api.getLeaderboard(user?.campus_id),
          api.getGoals(),
          api.getAchievements(),
        ]);
        if (statsRes.data) setStats(statsRes.data);
        if (leaderboardRes.data) setLeaderboard(leaderboardRes.data.slice(0, 5));
        if (goalsRes.data) setGoals(goalsRes.data);
        if (achievementsRes.data) setAchievements(achievementsRes.data);
      } catch (error) {
        console.error('Failed to fetch dashboard data:', error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchData();
  }, [user?.campus_id]);

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gold-600"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex justify-between items-center">
          <h1 className="text-2xl font-bold text-gray-900">U-Genius Dashboard</h1>
          <div className="flex items-center gap-4">
            <span className="text-gray-600">Welcome, {user?.first_name}!</span>
            <Button variant="outline" onClick={logout}>Logout</Button>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <StatCard title="Total Hours" value={stats?.total_hours?.toFixed(1) || '0'} icon="📚" />
          <StatCard title="This Week" value={stats?.weekly_hours?.toFixed(1) || '0'} icon="📅" />
          <StatCard title="Submissions" value={stats?.total_submissions?.toString() || '0'} icon="✅" />
          <StatCard title="Badges Earned" value={achievements.length.toString()} icon="🏆" />
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Quick Actions */}
          <div className="bg-white rounded-xl shadow-sm p-6">
            <h2 className="text-lg font-semibold mb-4">Quick Actions</h2>
            <div className="space-y-3">
              <Link to="/reading-hours/new">
                <Button variant="gold" className="w-full">Log Reading Hours</Button>
              </Link>
              <Link to="/goals/new">
                <Button variant="outline" className="w-full">Set New Goal</Button>
              </Link>
              <Link to="/study-groups">
                <Button variant="outline" className="w-full">Find Study Groups</Button>
              </Link>
            </div>
          </div>

          {/* Active Goals */}
          <div className="bg-white rounded-xl shadow-sm p-6">
            <h2 className="text-lg font-semibold mb-4">Active Goals</h2>
            {goals.length === 0 ? (
              <p className="text-gray-500 text-center py-4">No active goals. Set one to get started!</p>
            ) : (
              <div className="space-y-3">
                {goals.slice(0, 3).map((goal) => (
                  <div key={goal.id} className="border rounded-lg p-3">
                    <div className="flex justify-between items-center mb-2">
                      <span className="font-medium">{goal.title}</span>
                      <span className="text-sm text-gray-500">
                        {goal.current_value}/{goal.target_value}
                      </span>
                    </div>
                    <div className="w-full bg-gray-200 rounded-full h-2">
                      <div
                        className="bg-gold-500 h-2 rounded-full"
                        style={{ width: `${Math.min((goal.current_value / goal.target_value) * 100, 100)}%` }}
                      ></div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>

          {/* Leaderboard */}
          <div className="bg-white rounded-xl shadow-sm p-6">
            <h2 className="text-lg font-semibold mb-4">Weekly Leaderboard</h2>
            {leaderboard.length === 0 ? (
              <p className="text-gray-500 text-center py-4">No data yet. Start logging hours!</p>
            ) : (
              <div className="space-y-3">
                {leaderboard.map((entry) => (
                  <div key={entry.user_id} className="flex items-center gap-3">
                    <span className={`w-6 h-6 rounded-full flex items-center justify-center text-sm font-bold ${
                      entry.rank === 1 ? 'bg-gold-500 text-white' :
                      entry.rank === 2 ? 'bg-gray-300' :
                      entry.rank === 3 ? 'bg-amber-600 text-white' : 'bg-gray-100'
                    }`}>
                      {entry.rank}
                    </span>
                    <div className="flex-1">
                      <p className="font-medium">{entry.first_name} {entry.last_name}</p>
                      <p className="text-sm text-gray-500">{entry.total_hours.toFixed(1)} hours</p>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>

        {/* Recent Achievements */}
        {achievements.length > 0 && (
          <div className="mt-8 bg-white rounded-xl shadow-sm p-6">
            <h2 className="text-lg font-semibold mb-4">Recent Achievements</h2>
            <div className="flex flex-wrap gap-4">
              {achievements.slice(0, 6).map((achievement) => (
                <div key={achievement.id} className="flex items-center gap-2 bg-gold-50 rounded-lg px-4 py-2">
                  <span className="text-2xl">🏅</span>
                  <div>
                    <p className="font-medium">{achievement.badge_name}</p>
                    <p className="text-sm text-gray-500">{achievement.badge_description}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </main>
    </div>
  );
}

function StatCard({ title, value, icon }: { title: string; value: string; icon: string }) {
  return (
    <div className="bg-white rounded-xl shadow-sm p-6">
      <div className="flex items-center gap-4">
        <span className="text-3xl">{icon}</span>
        <div>
          <p className="text-sm text-gray-500">{title}</p>
          <p className="text-2xl font-bold text-gray-900">{value}</p>
        </div>
      </div>
    </div>
  );
}

