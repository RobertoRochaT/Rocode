import { CheckCircle2, Circle, PlayCircle } from 'lucide-react';

const PatternCard = ({ pattern, progress, onClick }) => {
  const totalProblems = pattern.problemCount || 0;
  const completedProblems = progress?.completed || 0;
  const progressPercentage = totalProblems > 0 ? (completedProblems / totalProblems) * 100 : 0;

  const getStatusClass = () => {
    if (completedProblems === 0) return 'pattern-node-incomplete';
    if (completedProblems === totalProblems) return 'pattern-node-complete';
    return 'pattern-node-in-progress';
  };

  const getStatusIcon = () => {
    if (completedProblems === totalProblems && totalProblems > 0) {
      return <CheckCircle2 className="w-5 h-5 text-green-400" />;
    }
    if (completedProblems > 0) {
      return <PlayCircle className="w-5 h-5 text-primary-400" />;
    }
    return <Circle className="w-5 h-5 text-dark-600" />;
  };

  return (
    <div
      onClick={() => onClick(pattern)}
      className={`pattern-node ${getStatusClass()} group`}
    >
      {/* Header */}
      <div className="flex items-start justify-between mb-3">
        <div className="flex items-center space-x-2">
          {getStatusIcon()}
          <span className="text-xs font-medium text-dark-400">
            Pattern {pattern.id}
          </span>
        </div>
        {totalProblems > 0 && (
          <span className="text-xs font-medium text-dark-500">
            {completedProblems}/{totalProblems}
          </span>
        )}
      </div>

      {/* Pattern Name */}
      <h3 className="text-sm font-semibold text-dark-100 mb-2 group-hover:text-primary-400 transition-colors">
        {pattern.name}
      </h3>

      {/* Description */}
      {pattern.description && (
        <p className="text-xs text-dark-400 mb-3 line-clamp-2">
          {pattern.description}
        </p>
      )}

      {/* Progress Bar */}
      {totalProblems > 0 && (
        <div className="progress-bar">
          <div
            className="progress-fill"
            style={{ width: `${progressPercentage}%` }}
          />
        </div>
      )}

      {/* Completion Badge */}
      {completedProblems === totalProblems && totalProblems > 0 && (
        <div className="absolute -top-2 -right-2 bg-green-600 text-white text-xs font-bold px-2 py-1 rounded-full shadow-lg">
          ✓
        </div>
      )}
    </div>
  );
};

export default PatternCard;
