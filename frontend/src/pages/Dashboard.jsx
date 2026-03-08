import { useState, useEffect } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { progressAPI } from '../services/api';
import { TrendingUp, CheckCircle2, Target, Award, Calendar } from 'lucide-react';

const Dashboard = () => {
  const { user } = useAuth();
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadStats();
  }, []);

  const loadStats = async () => {
    try {
      setLoading(true);
      const data = await progressAPI.getStats();
      setStats(data);
    } catch (error) {
      console.error('Failed to load stats:', error);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'Never';
    const date = new Date(dateString);
    const now = new Date();
    const diffTime = Math.abs(now - date);
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));

    if (diffDays === 0) return 'Today';
    if (diffDays === 1) return 'Yesterday';
    if (diffDays < 7) return `${diffDays} days ago`;
    if (diffDays < 30) return `${Math.floor(diffDays / 7)} weeks ago`;
    if (diffDays < 365) return `${Math.floor(diffDays / 30)} months ago`;
    return date.toLocaleDateString();
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="spinner w-12 h-12 mx-auto mb-4" />
          <p className="text-dark-400">Loading dashboard...</p>
        </div>
      </div>
    );
  }

  const statCards = [
    {
      icon: Target,
      title: 'Total Problems',
      value: stats?.total_problems || 0,
      color: 'text-primary-400',
      bgColor: 'bg-primary-900/30',
    },
    {
      icon: CheckCircle2,
      title: 'Completed',
      value: stats?.completed_problems || 0,
      color: 'text-green-400',
      bgColor: 'bg-green-900/30',
    },
    {
      icon: TrendingUp,
      title: 'Completion Rate',
      value: `${Math.round(stats?.completion_rate || 0)}%`,
      color: 'text-blue-400',
      bgColor: 'bg-blue-900/30',
    },
    {
      icon: Award,
      title: 'Patterns Mastered',
      value: stats?.pattern_stats?.filter(p => p.percentage === 100).length || 0,
      color: 'text-yellow-400',
      bgColor: 'bg-yellow-900/30',
    },
  ];

  return (
    <div className="min-h-screen pb-20">
      {/* Header */}
      <div className="bg-dark-900/50 border-b border-dark-800">
        <div className="container mx-auto px-4 py-8">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold mb-2">
                Welcome back, <span className="text-gradient">{user?.username}</span>!
              </h1>
              <p className="text-dark-400">Track your progress and stay motivated</p>
            </div>
            <div className="hidden md:block">
              <div className="flex items-center space-x-2 px-4 py-2 bg-dark-800 rounded-lg">
                <Calendar className="w-5 h-5 text-primary-400" />
                <span className="text-sm text-dark-300">
                  {new Date().toLocaleDateString('en-US', {
                    weekday: 'long',
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric'
                  })}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-8">
        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          {statCards.map((card, index) => (
            <div
              key={card.title}
              className="card p-6 card-hover animate-slide-up"
              style={{ animationDelay: `${index * 100}ms` }}
            >
              <div className="flex items-center justify-between mb-4">
                <div className={`p-3 rounded-xl ${card.bgColor}`}>
                  <card.icon className={`w-6 h-6 ${card.color}`} />
                </div>
              </div>
              <div className="text-3xl font-bold text-dark-100 mb-1">{card.value}</div>
              <div className="text-sm text-dark-400">{card.title}</div>
            </div>
          ))}
        </div>

        <div className="grid lg:grid-cols-3 gap-6">
          {/* Progress Overview */}
          <div className="lg:col-span-2 card p-6">
            <h2 className="text-xl font-bold text-dark-100 mb-6">Progress Overview</h2>

            {/* Overall Progress Bar */}
            <div className="mb-8">
              <div className="flex items-center justify-between mb-2">
                <span className="text-sm font-medium text-dark-300">Overall Completion</span>
                <span className="text-sm font-bold text-primary-400">
                  {Math.round(stats?.completion_rate || 0)}%
                </span>
              </div>
              <div className="progress-bar h-3">
                <div
                  className="progress-fill h-full"
                  style={{ width: `${stats?.completion_rate || 0}%` }}
                />
              </div>
              <div className="flex items-center justify-between mt-2 text-xs text-dark-500">
                <span>{stats?.completed_problems || 0} completed</span>
                <span>{(stats?.total_problems || 0) - (stats?.completed_problems || 0)} remaining</span>
              </div>
            </div>

            {/* Pattern Progress */}
            <div>
              <h3 className="text-sm font-semibold text-dark-300 mb-4">Pattern Categories</h3>
              <div className="space-y-4 max-h-96 overflow-y-auto custom-scrollbar">
                {stats?.pattern_stats?.slice(0, 10).map((pattern, index) => (
                  <div
                    key={pattern.pattern_id}
                    className="animate-slide-up"
                    style={{ animationDelay: `${index * 50}ms` }}
                  >
                    <div className="flex items-center justify-between mb-2">
                      <div className="flex items-center space-x-2">
                        <span className="text-sm text-dark-300">{pattern.pattern_name}</span>
                        {pattern.percentage === 100 && (
                          <CheckCircle2 className="w-4 h-4 text-green-400" />
                        )}
                      </div>
                      <span className="text-sm font-medium text-dark-400">
                        {pattern.completed}/{pattern.total}
                      </span>
                    </div>
                    <div className="progress-bar h-2">
                      <div
                        className={`h-full transition-all duration-500 ${
                          pattern.percentage === 100
                            ? 'bg-gradient-to-r from-green-600 to-green-500'
                            : 'bg-gradient-to-r from-primary-600 to-primary-500'
                        }`}
                        style={{ width: `${pattern.percentage}%` }}
                      />
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>

          {/* Recent Activity */}
          <div className="card p-6">
            <h2 className="text-xl font-bold text-dark-100 mb-6">Recent Activity</h2>

            {stats?.recent_activity && stats.recent_activity.length > 0 ? (
              <div className="space-y-4">
                {stats.recent_activity.slice(0, 8).map((activity, index) => (
                  <div
                    key={activity.id}
                    className="flex items-start space-x-3 pb-4 border-b border-dark-800 last:border-0 last:pb-0 animate-fade-in"
                    style={{ animationDelay: `${index * 100}ms` }}
                  >
                    <div className={`p-2 rounded-lg flex-shrink-0 ${
                      activity.completed ? 'bg-green-900/30' : 'bg-dark-800'
                    }`}>
                      {activity.completed ? (
                        <CheckCircle2 className="w-4 h-4 text-green-400" />
                      ) : (
                        <Target className="w-4 h-4 text-dark-500" />
                      )}
                    </div>
                    <div className="flex-1 min-w-0">
                      <div className="text-sm font-medium text-dark-200 mb-1">
                        Problem #{activity.problem_number}
                      </div>
                      <div className="text-xs text-dark-500">
                        {activity.completed ? 'Completed' : 'Updated'} {formatDate(activity.updated_at)}
                      </div>
                      {activity.difficulty && (
                        <div className="mt-2">
                          <span className={`badge text-xs ${
                            activity.difficulty === 'Easy' ? 'badge-easy' :
                            activity.difficulty === 'Medium' ? 'badge-medium' :
                            'badge-hard'
                          }`}>
                            {activity.difficulty}
                          </span>
                        </div>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="text-center py-12">
                <Target className="w-12 h-12 text-dark-700 mx-auto mb-4" />
                <p className="text-dark-400 text-sm">No activity yet</p>
                <p className="text-dark-500 text-xs mt-1">Start solving problems to see your progress here</p>
              </div>
            )}
          </div>
        </div>

        {/* Motivation Section */}
        <div className="mt-8 card p-8 text-center bg-gradient-to-br from-primary-900/20 to-dark-900 border-primary-800/30">
          <div className="max-w-2xl mx-auto">
            <Award className="w-12 h-12 text-primary-400 mx-auto mb-4" />
            <h3 className="text-2xl font-bold text-dark-100 mb-3">
              {stats?.completed_problems === 0
                ? 'Start Your Journey!'
                : stats?.completion_rate >= 75
                ? '🎉 Amazing Progress!'
                : stats?.completion_rate >= 50
                ? '🚀 Keep Going!'
                : '💪 You\'re Doing Great!'}
            </h3>
            <p className="text-dark-300 mb-6">
              {stats?.completed_problems === 0
                ? 'Begin your coding journey by tackling your first pattern. Every expert was once a beginner!'
                : stats?.completion_rate >= 75
                ? 'You\'re crushing it! You\'ve completed most of the problems. Keep up the excellent work!'
                : stats?.completion_rate >= 50
                ? 'You\'re halfway there! Consistency is key to mastering these patterns.'
                : 'Great start! Remember, every problem solved is a step towards mastery.'}
            </p>
            <div className="flex items-center justify-center space-x-4">
              <div className="text-center">
                <div className="text-3xl font-bold text-gradient mb-1">
                  {stats?.completed_problems || 0}
                </div>
                <div className="text-xs text-dark-500">Problems Solved</div>
              </div>
              <div className="h-12 w-px bg-dark-700" />
              <div className="text-center">
                <div className="text-3xl font-bold text-gradient mb-1">
                  {stats?.pattern_stats?.filter(p => p.completed > 0).length || 0}
                </div>
                <div className="text-xs text-dark-500">Patterns Started</div>
              </div>
              <div className="h-12 w-px bg-dark-700" />
              <div className="text-center">
                <div className="text-3xl font-bold text-gradient mb-1">
                  {stats?.pattern_stats?.filter(p => p.percentage === 100).length || 0}
                </div>
                <div className="text-xs text-dark-500">Patterns Mastered</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
