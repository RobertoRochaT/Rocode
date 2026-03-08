import { Link } from 'react-router-dom';
import { Code2, Target, TrendingUp, Users, ArrowRight, CheckCircle } from 'lucide-react';
import { useAuth } from '../contexts/AuthContext';

const Home = () => {
  const { isAuthenticated } = useAuth();

  const features = [
    {
      icon: Target,
      title: '94 Coding Patterns',
      description: 'Master essential coding patterns covering all major topics from Two Pointers to Dynamic Programming.',
    },
    {
      icon: TrendingUp,
      title: 'Track Your Progress',
      description: 'Monitor your learning journey with detailed statistics and completion tracking.',
    },
    {
      icon: Users,
      title: 'Structured Roadmap',
      description: 'Follow a well-organized roadmap designed to take you from beginner to expert.',
    },
  ];

  const stats = [
    { value: '94', label: 'Patterns' },
    { value: '600+', label: 'Problems' },
    { value: '15', label: 'Categories' },
  ];

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative overflow-hidden bg-gradient-to-br from-dark-950 via-dark-900 to-dark-950 pt-20 pb-32">
        {/* Background decoration */}
        <div className="absolute inset-0 overflow-hidden">
          <div className="absolute -top-40 -right-40 w-80 h-80 bg-primary-500/10 rounded-full blur-3xl" />
          <div className="absolute -bottom-40 -left-40 w-80 h-80 bg-primary-500/10 rounded-full blur-3xl" />
        </div>

        <div className="container mx-auto px-4 relative z-10">
          <div className="max-w-4xl mx-auto text-center">
            {/* Main heading */}
            <div className="mb-6 inline-flex items-center space-x-2 px-4 py-2 bg-primary-900/30 border border-primary-800/50 rounded-full">
              <Code2 className="w-4 h-4 text-primary-400" />
              <span className="text-sm text-primary-300">Master LeetCode Patterns</span>
            </div>

            <h1 className="text-5xl md:text-7xl font-bold mb-6 animate-fade-in">
              Your Roadmap to
              <span className="text-gradient block mt-2">Coding Excellence</span>
            </h1>

            <p className="text-xl text-dark-300 mb-10 max-w-2xl mx-auto animate-fade-in animation-delay-200">
              Master 94 essential coding patterns with a structured, visual roadmap.
              Track your progress and conquer LeetCode problems systematically.
            </p>

            {/* CTA Buttons */}
            <div className="flex flex-col sm:flex-row items-center justify-center gap-4 animate-fade-in animation-delay-400">
              <Link
                to={isAuthenticated ? '/roadmap' : '/register'}
                className="btn btn-primary text-lg px-8 py-3 flex items-center space-x-2 group"
              >
                <span>{isAuthenticated ? 'View Roadmap' : 'Get Started Free'}</span>
                <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
              </Link>
              {!isAuthenticated && (
                <Link
                  to="/login"
                  className="btn btn-outline text-lg px-8 py-3"
                >
                  Sign In
                </Link>
              )}
            </div>

            {/* Stats */}
            <div className="grid grid-cols-3 gap-8 mt-16 max-w-2xl mx-auto">
              {stats.map((stat, index) => (
                <div
                  key={stat.label}
                  className="text-center animate-slide-up"
                  style={{ animationDelay: `${600 + index * 100}ms` }}
                >
                  <div className="text-4xl font-bold text-gradient mb-2">
                    {stat.value}
                  </div>
                  <div className="text-sm text-dark-400">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 bg-dark-900/50">
        <div className="container mx-auto px-4">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              Why Choose <span className="text-gradient">RoCode</span>?
            </h2>
            <p className="text-dark-400 max-w-2xl mx-auto">
              A systematic approach to mastering coding interviews
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8 max-w-6xl mx-auto">
            {features.map((feature, index) => (
              <div
                key={feature.title}
                className="card card-hover p-8 text-center animate-slide-up"
                style={{ animationDelay: `${index * 100}ms` }}
              >
                <div className="inline-flex items-center justify-center w-16 h-16 bg-primary-900/30 rounded-2xl mb-6">
                  <feature.icon className="w-8 h-8 text-primary-400" />
                </div>
                <h3 className="text-xl font-semibold mb-3 text-dark-100">
                  {feature.title}
                </h3>
                <p className="text-dark-400">{feature.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* How It Works Section */}
      <section className="py-20">
        <div className="container mx-auto px-4">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              How It <span className="text-gradient">Works</span>
            </h2>
            <p className="text-dark-400 max-w-2xl mx-auto">
              Simple steps to start your coding journey
            </p>
          </div>

          <div className="max-w-4xl mx-auto space-y-8">
            {[
              {
                step: '01',
                title: 'Choose a Pattern',
                description: 'Browse through 94 carefully curated coding patterns organized by category.',
              },
              {
                step: '02',
                title: 'Solve Problems',
                description: 'Practice with handpicked LeetCode problems for each pattern. Click to open directly in LeetCode.',
              },
              {
                step: '03',
                title: 'Track Progress',
                description: 'Mark problems as complete and add notes. Watch your progress grow on the visual roadmap.',
              },
            ].map((step, index) => (
              <div
                key={step.step}
                className="flex items-start space-x-6 animate-slide-up"
                style={{ animationDelay: `${index * 150}ms` }}
              >
                <div className="flex-shrink-0 w-16 h-16 bg-gradient-primary rounded-xl flex items-center justify-center">
                  <span className="text-2xl font-bold text-white">{step.step}</span>
                </div>
                <div className="flex-1">
                  <h3 className="text-xl font-semibold mb-2 text-dark-100">
                    {step.title}
                  </h3>
                  <p className="text-dark-400">{step.description}</p>
                </div>
                <div className="hidden md:block">
                  <CheckCircle className="w-6 h-6 text-primary-500" />
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 bg-dark-900/50">
        <div className="container mx-auto px-4">
          <div className="max-w-3xl mx-auto text-center">
            <h2 className="text-3xl md:text-4xl font-bold mb-6">
              Ready to Master Coding Patterns?
            </h2>
            <p className="text-dark-400 mb-8 text-lg">
              Join thousands of developers improving their problem-solving skills
            </p>
            <Link
              to={isAuthenticated ? '/roadmap' : '/register'}
              className="btn btn-primary text-lg px-10 py-4 inline-flex items-center space-x-2 group"
            >
              <span>{isAuthenticated ? 'Go to Roadmap' : 'Start Learning Now'}</span>
              <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Home;
