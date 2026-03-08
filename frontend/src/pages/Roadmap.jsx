import { useState, useEffect } from "react";
import { useAuth } from "../contexts/AuthContext";
import { patternsAPI, problemsAPI, progressAPI } from "../services/api";
import PatternCard from "../components/PatternCard";
import ProblemList from "../components/ProblemList";
import { X, Filter, Search, TrendingUp } from "lucide-react";

const Roadmap = () => {
  const { user } = useAuth();
  const [patterns, setPatterns] = useState([]);
  const [problems, setProblems] = useState([]);
  const [userProgress, setUserProgress] = useState({});
  const [selectedPattern, setSelectedPattern] = useState(null);
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [searchQuery, setSearchQuery] = useState("");
  const [loading, setLoading] = useState(true);
  const [modalLoading, setModalLoading] = useState(false);

  // Group patterns by category
  const categories = [
    { id: "all", name: "All Patterns" },
    { id: "I", name: "Two Pointer" },
    { id: "II", name: "Sliding Window" },
    { id: "III", name: "Tree Traversal" },
    { id: "IV", name: "Graph Traversal" },
    { id: "V", name: "Dynamic Programming" },
    { id: "VI", name: "Heap" },
    { id: "VII", name: "Backtracking" },
    { id: "VIII", name: "Greedy" },
    { id: "IX", name: "Binary Search" },
    { id: "X", name: "Stack" },
    { id: "XI", name: "Bit Manipulation" },
    { id: "XII", name: "Linked List" },
    { id: "XIII", name: "Array/Matrix" },
    { id: "XIV", name: "String" },
    { id: "XV", name: "Design" },
  ];

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [patternsData, problemsData, progressData] = await Promise.all([
        patternsAPI.getAll().catch(() => []),
        problemsAPI.getAll().catch(() => []),
        progressAPI.getAll().catch(() => []),
      ]);

      if (!patternsData || !problemsData) {
        console.error("Failed to load patterns or problems");
        return;
      }

      // Count problems per pattern
      const patternProblemCount = {};
      (problemsData || []).forEach((problem) => {
        const patternId = problem.pattern_id;
        patternProblemCount[patternId] =
          (patternProblemCount[patternId] || 0) + 1;
      });

      // Add problem count to patterns
      const enrichedPatterns = (patternsData || []).map((pattern) => ({
        ...pattern,
        problemCount: patternProblemCount[pattern.id] || 0,
      }));

      // Process user progress
      const progressMap = {};
      (progressData || []).forEach((progress) => {
        if (!progressMap[progress.pattern_id]) {
          progressMap[progress.pattern_id] = { completed: 0, total: 0 };
        }
        progressMap[progress.pattern_id].total++;
        if (progress.completed) {
          progressMap[progress.pattern_id].completed++;
        }
      });

      setPatterns(enrichedPatterns);
      setProblems(problemsData);
      setUserProgress(progressMap);
    } catch (error) {
      console.error("Failed to load data:", error);
    } finally {
      setLoading(false);
    }
  };

  const handlePatternClick = async (pattern) => {
    setSelectedPattern(pattern);
    setModalLoading(true);
    try {
      const patternProblems = await problemsAPI.getAll(pattern.id);
      setProblems(patternProblems);
    } catch (error) {
      console.error("Failed to load problems:", error);
    } finally {
      setModalLoading(false);
    }
  };

  const handleCloseModal = () => {
    setSelectedPattern(null);
  };

  const handleProgressUpdate = async () => {
    // Reload progress data to update the pattern cards and overall stats
    // Do NOT reload the problems list - ProblemList manages its own state optimistically
    try {
      const progressData = await progressAPI.getAll();
      const progressMap = {};
      progressData.forEach((progress) => {
        if (!progressMap[progress.pattern_id]) {
          progressMap[progress.pattern_id] = { completed: 0, total: 0 };
        }
        progressMap[progress.pattern_id].total++;
        if (progress.completed) {
          progressMap[progress.pattern_id].completed++;
        }
      });
      setUserProgress(progressMap);

      // Don't reload problems here - this would overwrite the optimistic UI update
      // in ProblemList and cause the checkbox to flicker back to its old state
    } catch (error) {
      console.error("Failed to reload progress:", error);
    }
  };

  const filteredPatterns = patterns.filter((pattern) => {
    // Category filter
    if (selectedCategory !== "all") {
      if (!pattern.category.startsWith(selectedCategory + ".")) {
        return false;
      }
    }

    // Search filter
    if (searchQuery) {
      const query = searchQuery.toLowerCase();
      return (
        pattern.name.toLowerCase().includes(query) ||
        pattern.description?.toLowerCase().includes(query) ||
        pattern.category.toLowerCase().includes(query)
      );
    }

    return true;
  });

  // Group patterns by category for display
  const groupedPatterns = {};
  filteredPatterns.forEach((pattern) => {
    const category = pattern.category;
    if (!groupedPatterns[category]) {
      groupedPatterns[category] = [];
    }
    groupedPatterns[category].push(pattern);
  });

  const totalProblems = patterns.reduce(
    (sum, p) => sum + (p.problemCount || 0),
    0,
  );
  const completedProblems = Object.values(userProgress).reduce(
    (sum, p) => sum + (p.completed || 0),
    0,
  );
  const completionPercentage =
    totalProblems > 0
      ? Math.round((completedProblems / totalProblems) * 100)
      : 0;

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="spinner w-12 h-12 mx-auto mb-4" />
          <p className="text-dark-400">Loading roadmap...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen pb-20">
      {/* Header */}
      <div className="bg-dark-900/50 border-b border-dark-800">
        <div className="container mx-auto px-4 py-8">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h1 className="text-3xl font-bold mb-2">
                Learning <span className="text-gradient">Roadmap</span>
              </h1>
              <p className="text-dark-400">
                Master coding patterns step by step
              </p>
            </div>
            <div className="flex items-center space-x-4">
              <div className="text-right">
                <div className="text-sm text-dark-400 mb-1">
                  Overall Progress
                </div>
                <div className="flex items-center space-x-3">
                  <div className="text-2xl font-bold text-gradient">
                    {completionPercentage}%
                  </div>
                  <div className="text-sm text-dark-500">
                    {completedProblems}/{totalProblems}
                  </div>
                </div>
              </div>
              <div className="w-32 h-32 relative">
                <svg className="w-full h-full transform -rotate-90">
                  <circle
                    cx="64"
                    cy="64"
                    r="56"
                    stroke="currentColor"
                    strokeWidth="8"
                    fill="transparent"
                    className="text-dark-800"
                  />
                  <circle
                    cx="64"
                    cy="64"
                    r="56"
                    stroke="url(#gradient)"
                    strokeWidth="8"
                    fill="transparent"
                    strokeDasharray={`${2 * Math.PI * 56}`}
                    strokeDashoffset={`${
                      2 * Math.PI * 56 * (1 - completionPercentage / 100)
                    }`}
                    className="transition-all duration-1000"
                    strokeLinecap="round"
                  />
                  <defs>
                    <linearGradient
                      id="gradient"
                      x1="0%"
                      y1="0%"
                      x2="100%"
                      y2="100%"
                    >
                      <stop offset="0%" stopColor="#0ea5e9" />
                      <stop offset="100%" stopColor="#06b6d4" />
                    </linearGradient>
                  </defs>
                </svg>
                <div className="absolute inset-0 flex items-center justify-center">
                  <TrendingUp className="w-8 h-8 text-primary-400" />
                </div>
              </div>
            </div>
          </div>

          {/* Filters */}
          <div className="flex flex-col sm:flex-row gap-4">
            {/* Search */}
            <div className="flex-1 relative">
              <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <Search className="w-5 h-5 text-dark-500" />
              </div>
              <input
                type="text"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Search patterns..."
                className="input pl-10 w-full"
              />
            </div>

            {/* Category Filter */}
            <div className="relative sm:w-64">
              <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <Filter className="w-5 h-5 text-dark-500" />
              </div>
              <select
                value={selectedCategory}
                onChange={(e) => setSelectedCategory(e.target.value)}
                className="input pl-10 w-full appearance-none cursor-pointer"
              >
                {categories.map((cat) => (
                  <option key={cat.id} value={cat.id}>
                    {cat.name}
                  </option>
                ))}
              </select>
            </div>
          </div>
        </div>
      </div>

      {/* Pattern Grid */}
      <div className="container mx-auto px-4 py-8">
        {Object.entries(groupedPatterns).map(([category, categoryPatterns]) => (
          <div key={category} className="mb-12">
            <h2 className="text-xl font-bold text-dark-200 mb-6 pb-2 border-b border-dark-800">
              {category}
            </h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
              {categoryPatterns.map((pattern) => (
                <PatternCard
                  key={pattern.id}
                  pattern={pattern}
                  progress={userProgress[pattern.id]}
                  onClick={handlePatternClick}
                />
              ))}
            </div>
          </div>
        ))}

        {filteredPatterns.length === 0 && (
          <div className="text-center py-20">
            <p className="text-dark-400 text-lg">
              No patterns found matching your search.
            </p>
          </div>
        )}
      </div>

      {/* Problem Modal */}
      {selectedPattern && (
        <div className="fixed inset-0 z-50 overflow-hidden">
          {/* Backdrop */}
          <div
            className="absolute inset-0 bg-black/80 backdrop-blur-sm"
            onClick={handleCloseModal}
          />

          {/* Modal */}
          <div className="absolute inset-0 overflow-hidden flex items-center justify-center p-4">
            <div className="relative bg-dark-900 rounded-2xl shadow-2xl border border-dark-800 w-full max-w-4xl max-h-[90vh] flex flex-col animate-slide-up">
              {/* Modal Header */}
              <div className="flex items-start justify-between p-6 border-b border-dark-800">
                <div>
                  <div className="flex items-center space-x-3 mb-2">
                    <span className="text-sm font-medium text-dark-500">
                      Pattern {selectedPattern.id}
                    </span>
                    <span className="text-sm text-dark-600">•</span>
                    <span className="text-sm text-dark-500">
                      {selectedPattern.category}
                    </span>
                  </div>
                  <h2 className="text-2xl font-bold text-dark-100">
                    {selectedPattern.name}
                  </h2>
                  {selectedPattern.description && (
                    <p className="text-dark-400 mt-2">
                      {selectedPattern.description}
                    </p>
                  )}
                </div>
                <button
                  onClick={handleCloseModal}
                  className="p-2 text-dark-400 hover:text-dark-200 hover:bg-dark-800 rounded-lg transition-all"
                >
                  <X className="w-6 h-6" />
                </button>
              </div>

              {/* Modal Body */}
              <div className="flex-1 overflow-y-auto custom-scrollbar p-6">
                {modalLoading ? (
                  <div className="flex items-center justify-center py-20">
                    <div className="spinner w-8 h-8" />
                  </div>
                ) : (
                  <ProblemList
                    problems={problems}
                    onProgressUpdate={handleProgressUpdate}
                  />
                )}
              </div>

              {/* Modal Footer */}
              <div className="p-6 border-t border-dark-800 bg-dark-900/50">
                <div className="flex items-center justify-between">
                  <div className="text-sm text-dark-400">
                    {problems.length} problem{problems.length !== 1 ? "s" : ""}{" "}
                    in this pattern
                  </div>
                  <button
                    onClick={handleCloseModal}
                    className="btn btn-secondary"
                  >
                    Close
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Roadmap;
