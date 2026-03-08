import { useState, useEffect } from "react";
import {
  ExternalLink,
  CheckCircle2,
  Circle,
  ChevronDown,
  ChevronUp,
} from "lucide-react";
import { progressAPI } from "../services/api";

const ProblemList = ({ problems, onProgressUpdate }) => {
  const [expandedProblem, setExpandedProblem] = useState(null);
  const [notes, setNotes] = useState({});
  const [localProblems, setLocalProblems] = useState(problems || []);

  // Sync with incoming problems data
  useEffect(() => {
    if (problems && problems.length > 0) {
      setLocalProblems(problems);
    }
  }, [problems]);

  const handleToggleComplete = async (problem) => {
    const newCompletedState = !problem.completed;

    // Optimistic update
    setLocalProblems((prev) =>
      prev.map((p) =>
        p.number === problem.number
          ? { ...p, completed: newCompletedState }
          : p,
      ),
    );

    try {
      await progressAPI.update(
        problem.number,
        problem.pattern_id,
        newCompletedState,
        problem.difficulty,
        notes[problem.number] || "",
      );

      // Notify parent to update progress counts
      if (onProgressUpdate) {
        onProgressUpdate();
      }
    } catch (error) {
      console.error("Failed to update progress:", error);
      // Revert on error
      setLocalProblems((prev) =>
        prev.map((p) =>
          p.number === problem.number
            ? { ...p, completed: !newCompletedState }
            : p,
        ),
      );
    }
  };

  const handleNotesChange = (problemNumber, value) => {
    setNotes((prev) => ({ ...prev, [problemNumber]: value }));
  };

  const handleSaveNotes = async (problem) => {
    try {
      await progressAPI.update(
        problem.number,
        problem.pattern_id,
        problem.completed || false,
        problem.difficulty,
        notes[problem.number] || "",
      );

      if (onProgressUpdate) {
        onProgressUpdate();
      }
    } catch (error) {
      console.error("Failed to save notes:", error);
    }
  };

  const getDifficultyBadge = (difficulty) => {
    const classes = {
      Easy: "badge-easy",
      Medium: "badge-medium",
      Hard: "badge-hard",
    };
    return (
      <span className={`badge ${classes[difficulty] || "badge-medium"}`}>
        {difficulty}
      </span>
    );
  };

  const openLeetCode = (url) => {
    window.open(url, "_blank", "noopener,noreferrer");
  };

  if (!localProblems || localProblems.length === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-dark-400">No problems found for this pattern.</p>
      </div>
    );
  }

  return (
    <div className="space-y-2">
      {localProblems.map((problem) => (
        <div
          key={problem.number}
          className={`card transition-all duration-200 ${
            problem.completed ? "border-green-800/50 bg-green-900/10" : ""
          }`}
        >
          <div className="p-4">
            {/* Main Row */}
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4 flex-1">
                {/* Checkbox - Click to toggle completion */}
                <button
                  onClick={() => handleToggleComplete(problem)}
                  className="focus:outline-none group"
                  title={
                    problem.completed
                      ? "Mark as incomplete"
                      : "Mark as complete"
                  }
                >
                  {problem.completed ? (
                    <CheckCircle2 className="w-6 h-6 text-green-500 group-hover:text-green-400 transition-colors" />
                  ) : (
                    <Circle className="w-6 h-6 text-dark-600 group-hover:text-primary-500 transition-colors" />
                  )}
                </button>

                {/* Problem Number & Title */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-center space-x-3">
                    <span className="text-sm font-mono text-dark-500">
                      #{problem.number}
                    </span>
                    <h3
                      className={`text-sm font-medium ${
                        problem.completed
                          ? "text-dark-400 line-through"
                          : "text-dark-100"
                      }`}
                    >
                      {problem.title}
                    </h3>
                  </div>
                </div>

                {/* Difficulty Badge */}
                <div>{getDifficultyBadge(problem.difficulty)}</div>

                {/* Actions */}
                <div className="flex items-center space-x-2">
                  {/* LeetCode Link */}
                  <button
                    onClick={() => openLeetCode(problem.leetcode_url)}
                    className="p-2 text-dark-400 hover:text-primary-400 hover:bg-dark-800 rounded-lg transition-all"
                    title="Open in LeetCode"
                  >
                    <ExternalLink className="w-4 h-4" />
                  </button>

                  {/* Expand/Collapse */}
                  <button
                    onClick={() =>
                      setExpandedProblem(
                        expandedProblem === problem.number
                          ? null
                          : problem.number,
                      )
                    }
                    className="p-2 text-dark-400 hover:text-dark-200 hover:bg-dark-800 rounded-lg transition-all"
                  >
                    {expandedProblem === problem.number ? (
                      <ChevronUp className="w-4 h-4" />
                    ) : (
                      <ChevronDown className="w-4 h-4" />
                    )}
                  </button>
                </div>
              </div>
            </div>

            {/* Expanded Section - Notes */}
            {expandedProblem === problem.number && (
              <div className="mt-4 pt-4 border-t border-dark-800">
                <label className="block text-sm font-medium text-dark-300 mb-2">
                  Notes
                </label>
                <textarea
                  value={notes[problem.number] || ""}
                  onChange={(e) =>
                    handleNotesChange(problem.number, e.target.value)
                  }
                  onBlur={() => handleSaveNotes(problem)}
                  placeholder="Add your notes, solution approach, or thoughts..."
                  className="input min-h-[100px] resize-y"
                />
                <p className="text-xs text-dark-500 mt-2">
                  Notes are saved automatically when you click outside the text
                  area.
                </p>
              </div>
            )}
          </div>
        </div>
      ))}
    </div>
  );
};

export default ProblemList;
