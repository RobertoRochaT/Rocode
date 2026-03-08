package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Pattern struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url,omitempty"`
}

type Problem struct {
	Number      string `json:"number"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	LeetCodeURL string `json:"leetcode_url"`
	PatternID   int    `json:"pattern_id"`
	Completed   bool   `json:"completed"`
}

type Progress struct {
	ID            int        `json:"id"`
	UserID        int        `json:"user_id"`
	ProblemNumber string     `json:"problem_number"`
	PatternID     int        `json:"pattern_id"`
	Completed     bool       `json:"completed"`
	Difficulty    string     `json:"difficulty"`
	Notes         string     `json:"notes,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type UpdateProgressRequest struct {
	ProblemNumber string `json:"problem_number" binding:"required"`
	PatternID     int    `json:"pattern_id" binding:"required"`
	Completed     bool   `json:"completed"`
	Difficulty    string `json:"difficulty"`
	Notes         string `json:"notes"`
}

type Stats struct {
	TotalProblems     int            `json:"total_problems"`
	CompletedProblems int            `json:"completed_problems"`
	CompletionRate    float64        `json:"completion_rate"`
	PatternStats      []PatternStats `json:"pattern_stats"`
	RecentActivity    []Progress     `json:"recent_activity"`
}

type PatternStats struct {
	PatternID   int     `json:"pattern_id"`
	PatternName string  `json:"pattern_name"`
	Total       int     `json:"total"`
	Completed   int     `json:"completed"`
	Percentage  float64 `json:"percentage"`
}

// getPatternsHandler returns all programming patterns
func getPatternsHandler(c *gin.Context) {
	patterns := []Pattern{
		// Two Pointer Patterns
		{ID: 1, Name: "Converging", Category: "I. Two Pointer Patterns", Description: "Two pointers moving towards each other"},
		{ID: 2, Name: "Fast & Slow", Category: "I. Two Pointer Patterns", Description: "Two pointers moving at different speeds"},
		{ID: 3, Name: "Fixed Separation", Category: "I. Two Pointer Patterns", Description: "Two pointers with fixed distance"},
		{ID: 4, Name: "In-place Array Modification", Category: "I. Two Pointer Patterns", Description: "Modify array using two pointers"},
		{ID: 5, Name: "String Comparison with special characters", Category: "I. Two Pointer Patterns", Description: "Compare strings with special handling"},
		{ID: 6, Name: "Expanding From Center", Category: "I. Two Pointer Patterns", Description: "Expand outward from center point"},
		{ID: 7, Name: "String Reversal", Category: "I. Two Pointer Patterns", Description: "Reverse strings using two pointers"},

		// Sliding Window Patterns
		{ID: 8, Name: "Fixed Size", Category: "II. Sliding Window Patterns", Description: "Fixed window size sliding"},
		{ID: 9, Name: "Variable Size", Category: "II. Sliding Window Patterns", Description: "Dynamic window size adjustment"},
		{ID: 10, Name: "Monotonic Queue for Max/Min", Category: "II. Sliding Window Patterns", Description: "Track max/min in sliding window"},
		{ID: 11, Name: "Character Frequency Matching", Category: "II. Sliding Window Patterns", Description: "Match character frequencies"},

		// Tree Traversal Patterns
		{ID: 12, Name: "Level Order Traversal", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "BFS level-by-level traversal"},
		{ID: 13, Name: "Recursive Preorder Traversal", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "Root-Left-Right traversal"},
		{ID: 14, Name: "Recursive Inorder Traversal", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "Left-Root-Right traversal"},
		{ID: 15, Name: "Recursive Postorder Traversal", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "Left-Right-Root traversal"},
		{ID: 16, Name: "Lowest Common Ancestor", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "Find LCA of two nodes"},
		{ID: 17, Name: "Serialization and Deserialization", Category: "III. Tree Traversal Patterns (DFS & BFS)", Description: "Convert tree to/from string"},

		// Graph Traversal Patterns
		{ID: 18, Name: "DFS - Connected Components / Island Counting", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find connected components using DFS"},
		{ID: 19, Name: "BFS - Connected Components / Island Counting", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find connected components using BFS"},
		{ID: 20, Name: "DFS - Cycle Detection", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Detect cycles in graph"},
		{ID: 21, Name: "BFS - Topological Sort (Kahn's Algorithm)", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Topological ordering"},
		{ID: 22, Name: "Deep Copy / Cloning", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Clone graph structures"},
		{ID: 23, Name: "Shortest Path", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find shortest path (Dijkstra)"},
		{ID: 24, Name: "Shortest Path (Bellman-Ford / BFS+K)", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Shortest path with constraints"},
		{ID: 25, Name: "Union-Find", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Disjoint set union"},
		{ID: 26, Name: "Strongly Connected Components (Kosaraju / Tarjan)", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find SCCs"},
		{ID: 27, Name: "Bridges & Articulation Points (Tarjan low-link)", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find critical connections"},
		{ID: 28, Name: "Minimum Spanning Tree (Kruskal / Prim / DSU + heap)", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "Find MST"},
		{ID: 29, Name: "Bidirectional BFS", Category: "IV. Graph Traversal Patterns (DFS & BFS)", Description: "BFS from both ends"},

		// Dynamic Programming Patterns
		{ID: 30, Name: "Fibonacci Style", Category: "V. Dynamic Programming (DP) Patterns", Description: "Classic DP with recurrence"},
		{ID: 31, Name: "Kadane's Algorithm for Max/Min Subarray", Category: "V. Dynamic Programming (DP) Patterns", Description: "Maximum subarray sum"},
		{ID: 32, Name: "Coin Change / Unbounded Knapsack Style", Category: "V. Dynamic Programming (DP) Patterns", Description: "Unbounded knapsack problems"},
		{ID: 33, Name: "0/1 Knapsack, Subset Sum Style", Category: "V. Dynamic Programming (DP) Patterns", Description: "0/1 knapsack problems"},
		{ID: 34, Name: "Word Break Style", Category: "V. Dynamic Programming (DP) Patterns", Description: "String breaking problems"},
		{ID: 35, Name: "Longest Common Subsequence - LCS", Category: "V. Dynamic Programming (DP) Patterns", Description: "LCS problems"},
		{ID: 36, Name: "Edit Distance / Levenshtein Distance", Category: "V. Dynamic Programming (DP) Patterns", Description: "String edit distance"},
		{ID: 37, Name: "Unique Paths on Grid", Category: "V. Dynamic Programming (DP) Patterns", Description: "Grid path counting"},
		{ID: 38, Name: "Interval DP", Category: "V. Dynamic Programming (DP) Patterns", Description: "DP on intervals"},
		{ID: 39, Name: "Catalan Numbers", Category: "V. Dynamic Programming (DP) Patterns", Description: "Catalan number problems"},
		{ID: 40, Name: "Longest Increasing Subsequence", Category: "V. Dynamic Programming (DP) Patterns", Description: "LIS problems"},
		{ID: 41, Name: "Stock problems", Category: "V. Dynamic Programming (DP) Patterns", Description: "Buy/sell stock variations"},

		// Heap (Priority Queue) Patterns
		{ID: 42, Name: "Top K Elements", Category: "VI. Heap (Priority Queue) Patterns", Description: "Find top K elements"},
		{ID: 43, Name: "Two Heaps for Median Finding", Category: "VI. Heap (Priority Queue) Patterns", Description: "Maintain median with heaps"},
		{ID: 44, Name: "K-way Merge", Category: "VI. Heap (Priority Queue) Patterns", Description: "Merge K sorted lists"},
		{ID: 45, Name: "Scheduling / Minimum Cost", Category: "VI. Heap (Priority Queue) Patterns", Description: "Scheduling with heaps"},

		// Backtracking Patterns
		{ID: 46, Name: "Subsets (Include/Exclude)", Category: "VII. Backtracking Patterns", Description: "Generate all subsets"},
		{ID: 47, Name: "Permutations", Category: "VII. Backtracking Patterns", Description: "Generate permutations"},
		{ID: 48, Name: "Combination Sum", Category: "VII. Backtracking Patterns", Description: "Find combinations that sum"},
		{ID: 49, Name: "Parentheses Generation", Category: "VII. Backtracking Patterns", Description: "Generate valid parentheses"},
		{ID: 50, Name: "Word Search / Path Finding in Grid", Category: "VII. Backtracking Patterns", Description: "Search words in grid"},
		{ID: 51, Name: "N-Queens / Constraint Satisfaction", Category: "VII. Backtracking Patterns", Description: "N-Queens problem"},
		{ID: 52, Name: "Palindrome Partitioning", Category: "VII. Backtracking Patterns", Description: "Partition into palindromes"},

		// Greedy Patterns
		{ID: 53, Name: "Interval Merging/Scheduling", Category: "VIII. Greedy Patterns", Description: "Merge/schedule intervals"},
		{ID: 54, Name: "Jump Game Reachability/Minimization", Category: "VIII. Greedy Patterns", Description: "Jump game problems"},
		{ID: 55, Name: "Buy/Sell Stock", Category: "VIII. Greedy Patterns", Description: "Stock trading greedy"},
		{ID: 56, Name: "Gas Station Circuit", Category: "VIII. Greedy Patterns", Description: "Circuit completion"},
		{ID: 57, Name: "Task Scheduling", Category: "VIII. Greedy Patterns", Description: "Schedule tasks greedily"},
		{ID: 58, Name: "Sorting Based", Category: "VIII. Greedy Patterns", Description: "Greedy with sorting"},

		// Binary Search Patterns
		{ID: 59, Name: "On Sorted Array/List", Category: "IX. Binary Search Patterns", Description: "Standard binary search"},
		{ID: 60, Name: "Find Min/Max in Rotated Sorted Array", Category: "IX. Binary Search Patterns", Description: "Search in rotated array"},
		{ID: 61, Name: "On Answer / Condition Function", Category: "IX. Binary Search Patterns", Description: "Binary search on answer space"},
		{ID: 62, Name: "Find First/Last Occurrence", Category: "IX. Binary Search Patterns", Description: "Find boundaries"},
		{ID: 63, Name: "Median / Kth across Two Sorted Arrays", Category: "IX. Binary Search Patterns", Description: "Kth element in merged arrays"},

		// Stack Patterns
		{ID: 64, Name: "Valid Parentheses Matching", Category: "X. Stack Patterns", Description: "Match parentheses"},
		{ID: 65, Name: "Monotonic Stack", Category: "X. Stack Patterns", Description: "Maintain monotonic property"},
		{ID: 66, Name: "Expression Evaluation", Category: "X. Stack Patterns", Description: "Evaluate expressions"},
		{ID: 67, Name: "Simulation / Backtracking Helper", Category: "X. Stack Patterns", Description: "Stack for simulation"},
		{ID: 68, Name: "Min Stack Design", Category: "X. Stack Patterns", Description: "Stack with min tracking"},
		{ID: 69, Name: "Largest Rectangle in Histogram", Category: "X. Stack Patterns", Description: "Rectangle area problems"},

		// Bit Manipulation Patterns
		{ID: 70, Name: "Bitwise XOR - Finding Single/Missing Number", Category: "XI. Bit Manipulation Patterns", Description: "XOR properties"},
		{ID: 71, Name: "Bitwise AND - Counting Set Bits (Hamming Weight)", Category: "XI. Bit Manipulation Patterns", Description: "Count bits"},
		{ID: 72, Name: "Bitwise DP - Counting Bits Optimization", Category: "XI. Bit Manipulation Patterns", Description: "DP with bitmasks"},
		{ID: 73, Name: "Bitwise Operations - Power of Two/Four Check", Category: "XI. Bit Manipulation Patterns", Description: "Power checks"},

		// Linked List Manipulation Patterns
		{ID: 74, Name: "In-place Reversal", Category: "XII. Linked List Manipulation Patterns", Description: "Reverse linked list"},
		{ID: 75, Name: "Merging Two Sorted Lists", Category: "XII. Linked List Manipulation Patterns", Description: "Merge lists"},
		{ID: 76, Name: "Addition of Numbers", Category: "XII. Linked List Manipulation Patterns", Description: "Add numbers as lists"},
		{ID: 77, Name: "Intersection Detection", Category: "XII. Linked List Manipulation Patterns", Description: "Find intersection"},
		{ID: 78, Name: "Reordering / Partitioning", Category: "XII. Linked List Manipulation Patterns", Description: "Reorder list nodes"},

		// Array/Matrix Manipulation Patterns
		{ID: 79, Name: "In-place Rotation", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Rotate array/matrix"},
		{ID: 80, Name: "Spiral Traversal", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Traverse in spiral"},
		{ID: 81, Name: "In-place Marking", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Mark elements in-place"},
		{ID: 82, Name: "Prefix/Suffix Products", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Product arrays"},
		{ID: 83, Name: "Plus One", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Increment number arrays"},
		{ID: 84, Name: "In-place from End", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Modify from end"},
		{ID: 85, Name: "Cyclic Sort", Category: "XIII. Array/Matrix Manipulation Patterns", Description: "Sort by cycling"},

		// String Manipulation Patterns
		{ID: 86, Name: "Palindrome Check", Category: "XIV. String Manipulation Patterns", Description: "Check palindromes"},
		{ID: 87, Name: "Anagram Check", Category: "XIV. String Manipulation Patterns", Description: "Check anagrams"},
		{ID: 88, Name: "Roman to Integer Conversion", Category: "XIV. String Manipulation Patterns", Description: "Roman numeral conversion"},
		{ID: 89, Name: "String to Integer (atoi)", Category: "XIV. String Manipulation Patterns", Description: "Parse integers"},
		{ID: 90, Name: "Manual Simulation", Category: "XIV. String Manipulation Patterns", Description: "Simulate operations"},
		{ID: 91, Name: "String Matching - Naive / KMP / Rabin-Karp", Category: "XIV. String Manipulation Patterns", Description: "Pattern matching"},
		{ID: 92, Name: "Repeated Substring Pattern Detection", Category: "XIV. String Manipulation Patterns", Description: "Detect patterns"},

		// Design Patterns
		{ID: 93, Name: "Design (General/Specific)", Category: "XV. Design Patterns", Description: "Data structure design"},
		{ID: 94, Name: "Tries", Category: "XV. Design Patterns", Description: "Trie data structure"},
	}

	c.JSON(http.StatusOK, patterns)
}

// getProblemsHandler returns problems, optionally filtered by pattern
func getProblemsHandler(c *gin.Context) {
	patternID := c.Query("pattern_id")
	userID, exists := c.Get("userID")

	problems := getAllProblems()

	// Filter by pattern if specified
	if patternID != "" {
		var filtered []Problem
		for _, p := range problems {
			if p.PatternID == parseIntOrZero(patternID) {
				filtered = append(filtered, p)
			}
		}
		problems = filtered
	}

	// Add completion status if user is authenticated
	if exists {
		uid := userID.(int)
		for i := range problems {
			var completed bool
			err := db.QueryRow(
				"SELECT completed FROM problem_progress WHERE user_id = $1 AND problem_number = $2",
				uid, problems[i].Number,
			).Scan(&completed)
			if err == nil {
				problems[i].Completed = completed
			}
		}
	}

	c.JSON(http.StatusOK, problems)
}

// getProgressHandler returns user's progress
func getProgressHandler(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := userID.(int)

	rows, err := db.Query(`
		SELECT id, user_id, problem_number, pattern_id, completed, difficulty,
		       notes, completed_at, created_at, updated_at
		FROM problem_progress
		WHERE user_id = $1
		ORDER BY updated_at DESC
	`, uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch progress"})
		return
	}
	defer rows.Close()

	var progress []Progress
	for rows.Next() {
		var p Progress
		err := rows.Scan(
			&p.ID, &p.UserID, &p.ProblemNumber, &p.PatternID, &p.Completed,
			&p.Difficulty, &p.Notes, &p.CompletedAt, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			continue
		}
		progress = append(progress, p)
	}

	c.JSON(http.StatusOK, progress)
}

// updateProgressHandler updates or creates progress for a problem
func updateProgressHandler(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := userID.(int)

	var req UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var completedAt *time.Time
	if req.Completed {
		now := time.Now()
		completedAt = &now
	}

	var id int
	err := db.QueryRow(`
		INSERT INTO problem_progress (user_id, problem_number, pattern_id, completed, difficulty, notes, completed_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP)
		ON CONFLICT (user_id, problem_number)
		DO UPDATE SET
			completed = $4,
			difficulty = $5,
			notes = $6,
			completed_at = $7,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id
	`, uid, req.ProblemNumber, req.PatternID, req.Completed, req.Difficulty, req.Notes, completedAt).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update progress"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Progress updated successfully"})
}

// getStatsHandler returns user statistics
func getStatsHandler(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := userID.(int)

	var stats Stats
	var completed int

	// Get total completed
	err := db.QueryRow(
		"SELECT COUNT(*) FROM problem_progress WHERE user_id = $1 AND completed = true",
		uid,
	).Scan(&completed)

	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
		return
	}

	stats.CompletedProblems = completed
	stats.TotalProblems = getTotalProblemsCount()
	if stats.TotalProblems > 0 {
		stats.CompletionRate = float64(completed) / float64(stats.TotalProblems) * 100
	}

	// Get recent activity
	rows, err := db.Query(`
		SELECT id, user_id, problem_number, pattern_id, completed, difficulty, notes, completed_at, created_at, updated_at
		FROM problem_progress
		WHERE user_id = $1
		ORDER BY updated_at DESC
		LIMIT 10
	`, uid)

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var p Progress
			rows.Scan(&p.ID, &p.UserID, &p.ProblemNumber, &p.PatternID, &p.Completed, &p.Difficulty, &p.Notes, &p.CompletedAt, &p.CreatedAt, &p.UpdatedAt)
			stats.RecentActivity = append(stats.RecentActivity, p)
		}
	}

	c.JSON(http.StatusOK, stats)
}

func parseIntOrZero(s string) int {
	var result int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		}
	}
	return result
}

func getTotalProblemsCount() int {
	return len(getAllProblems())
}
