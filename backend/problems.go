package main

import "strings"

// getAllProblems returns all problems organized by pattern
func getAllProblems() []Problem {
	return []Problem{
		// Pattern 1: Converging
		{Number: "11", Title: "Container With Most Water", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/container-with-most-water/", PatternID: 1},
		{Number: "15", Title: "3Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/3sum/", PatternID: 1},
		{Number: "16", Title: "3Sum Closest", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/3sum-closest/", PatternID: 1},
		{Number: "18", Title: "4Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/4sum/", PatternID: 1},
		{Number: "167", Title: "Two Sum II - Input Array Is Sorted", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/", PatternID: 1},
		{Number: "349", Title: "Intersection of Two Arrays", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/intersection-of-two-arrays/", PatternID: 1},
		{Number: "881", Title: "Boats to Save People", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/boats-to-save-people/", PatternID: 1},
		{Number: "977", Title: "Squares of a Sorted Array", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/squares-of-a-sorted-array/", PatternID: 1},
		{Number: "259", Title: "3Sum Smaller", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/3sum-smaller/", PatternID: 1},

		// Pattern 2: Fast & Slow
		{Number: "141", Title: "Linked List Cycle", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/linked-list-cycle/", PatternID: 2},
		{Number: "202", Title: "Happy Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/happy-number/", PatternID: 2},
		{Number: "287", Title: "Find the Duplicate Number", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-duplicate-number/", PatternID: 2},
		{Number: "392", Title: "Is Subsequence", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/is-subsequence/", PatternID: 2},

		// Pattern 3: Fixed Separation
		{Number: "19", Title: "Remove Nth Node From End of List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/remove-nth-node-from-end-of-list/", PatternID: 3},
		{Number: "876", Title: "Middle of the Linked List", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/middle-of-the-linked-list/", PatternID: 3},
		{Number: "2095", Title: "Delete the Middle Node of a Linked List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/", PatternID: 3},

		// Pattern 4: In-place Array Modification
		{Number: "26", Title: "Remove Duplicates from Sorted Array", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/remove-duplicates-from-sorted-array/", PatternID: 4},
		{Number: "27", Title: "Remove Element", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/remove-element/", PatternID: 4},
		{Number: "75", Title: "Sort Colors", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/sort-colors/", PatternID: 4},
		{Number: "80", Title: "Remove Duplicates from Sorted Array II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/", PatternID: 4},
		{Number: "283", Title: "Move Zeroes", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/move-zeroes/", PatternID: 4},
		{Number: "443", Title: "String Compression", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/string-compression/", PatternID: 4},
		{Number: "905", Title: "Sort Array By Parity", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/sort-array-by-parity/", PatternID: 4},
		{Number: "2337", Title: "Move Pieces to Obtain a String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/move-pieces-to-obtain-a-string/", PatternID: 4},
		{Number: "2938", Title: "Separate Black and White Balls", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/separate-black-and-white-balls/", PatternID: 4},

		// Pattern 5: String Comparison with special characters
		{Number: "844", Title: "Backspace String Compare", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/backspace-string-compare/", PatternID: 5},
		{Number: "1598", Title: "Crawler Log Folder", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/crawler-log-folder/", PatternID: 5},
		{Number: "2390", Title: "Removing Stars From a String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/removing-stars-from-a-string/", PatternID: 5},

		// Pattern 6: Expanding From Center
		{Number: "5", Title: "Longest Palindromic Substring", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-palindromic-substring/", PatternID: 6},
		{Number: "647", Title: "Palindromic Substrings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/palindromic-substrings/", PatternID: 6},

		// Pattern 7: String Reversal
		{Number: "151", Title: "Reverse Words in a String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/reverse-words-in-a-string/", PatternID: 7},
		{Number: "344", Title: "Reverse String", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/reverse-string/", PatternID: 7},
		{Number: "345", Title: "Reverse Vowels of a String", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/reverse-vowels-of-a-string/", PatternID: 7},
		{Number: "541", Title: "Reverse String II", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/reverse-string-ii/", PatternID: 7},

		// Pattern 8: Fixed Size
		{Number: "346", Title: "Moving Average from Data Stream", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/moving-average-from-data-stream/", PatternID: 8},
		{Number: "643", Title: "Maximum Average Subarray I", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/maximum-average-subarray-i/", PatternID: 8},
		{Number: "2985", Title: "Calculate Compressed Mean", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/calculate-compressed-mean/", PatternID: 8},
		{Number: "3254", Title: "Find the Power of K-Size Subarrays I", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/", PatternID: 8},
		{Number: "3318", Title: "Find X-Sum of All K-Long Subarrays I", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/", PatternID: 8},

		// Pattern 9: Variable Size
		{Number: "3", Title: "Longest Substring Without Repeating Characters", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-substring-without-repeating-characters/", PatternID: 9},
		{Number: "76", Title: "Minimum Window Substring", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-window-substring/", PatternID: 9},
		{Number: "209", Title: "Minimum Size Subarray Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-size-subarray-sum/", PatternID: 9},
		{Number: "219", Title: "Contains Duplicate II", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/contains-duplicate-ii/", PatternID: 9},
		{Number: "424", Title: "Longest Repeating Character Replacement", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-repeating-character-replacement/", PatternID: 9},
		{Number: "713", Title: "Subarray Product Less Than K", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/subarray-product-less-than-k/", PatternID: 9},
		{Number: "904", Title: "Fruit Into Baskets", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/fruit-into-baskets/", PatternID: 9},
		{Number: "1004", Title: "Max Consecutive Ones III", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/max-consecutive-ones-iii/", PatternID: 9},
		{Number: "1438", Title: "Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/", PatternID: 9},
		{Number: "1493", Title: "Longest Subarray of 1's After Deleting One Element", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/", PatternID: 9},
		{Number: "1658", Title: "Minimum Operations to Reduce X to Zero", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/", PatternID: 9},
		{Number: "1838", Title: "Frequency of the Most Frequent Element", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/frequency-of-the-most-frequent-element/", PatternID: 9},
		{Number: "2461", Title: "Maximum Sum of Distinct Subarrays With Length K", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/", PatternID: 9},
		{Number: "2516", Title: "Take K of Each Character From Left and Right", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/", PatternID: 9},
		{Number: "2762", Title: "Continuous Subarrays", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/continuous-subarrays/", PatternID: 9},
		{Number: "2779", Title: "Maximum Beauty of an Array After Applying Operation", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/", PatternID: 9},
		{Number: "2981", Title: "Find Longest Special Substring That Occurs Thrice I", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-i/", PatternID: 9},
		{Number: "3026", Title: "Maximum Good Subarray Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-good-subarray-sum/", PatternID: 9},
		{Number: "3346", Title: "Maximum Frequency of an Element After Performing Operations I", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-i/", PatternID: 9},
		{Number: "3347", Title: "Maximum Frequency of an Element After Performing Operations II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/", PatternID: 9},

		// Pattern 10: Monotonic Queue for Max/Min
		{Number: "239", Title: "Sliding Window Maximum", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/sliding-window-maximum/", PatternID: 10},
		{Number: "862", Title: "Shortest Subarray with Sum at Least K", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/", PatternID: 10},
		{Number: "1696", Title: "Jump Game VI", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/jump-game-vi/", PatternID: 10},

		// Pattern 11: Character Frequency Matching
		{Number: "1", Title: "Two Sum", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/two-sum/", PatternID: 11},
		{Number: "438", Title: "Find All Anagrams in a String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-all-anagrams-in-a-string/", PatternID: 11},
		{Number: "567", Title: "Permutation in String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/permutation-in-string/", PatternID: 11},

		// Pattern 12: Level Order Traversal
		{Number: "102", Title: "Binary Tree Level Order Traversal", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/binary-tree-level-order-traversal/", PatternID: 12},
		{Number: "103", Title: "Binary Tree Zigzag Level Order Traversal", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/", PatternID: 12},
		{Number: "199", Title: "Binary Tree Right Side View", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/binary-tree-right-side-view/", PatternID: 12},
		{Number: "515", Title: "Find Largest Value in Each Tree Row", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-largest-value-in-each-tree-row/", PatternID: 12},
		{Number: "1161", Title: "Maximum Level Sum of a Binary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/", PatternID: 12},

		// Pattern 13: Recursive Preorder Traversal
		{Number: "100", Title: "Same Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/same-tree/", PatternID: 13},
		{Number: "101", Title: "Symmetric Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/symmetric-tree/", PatternID: 13},
		{Number: "105", Title: "Construct Binary Tree from Preorder and Inorder Traversal", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/", PatternID: 13},
		{Number: "114", Title: "Flatten Binary Tree to Linked List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/flatten-binary-tree-to-linked-list/", PatternID: 13},
		{Number: "226", Title: "Invert Binary Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/invert-binary-tree/", PatternID: 13},
		{Number: "257", Title: "Binary Tree Paths", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/binary-tree-paths/", PatternID: 13},
		{Number: "988", Title: "Smallest String Starting From Leaf", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/smallest-string-starting-from-leaf/", PatternID: 13},

		// Pattern 14: Recursive Inorder Traversal
		{Number: "94", Title: "Binary Tree Inorder Traversal", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/binary-tree-inorder-traversal/", PatternID: 14},
		{Number: "98", Title: "Validate Binary Search Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/validate-binary-search-tree/", PatternID: 14},
		{Number: "173", Title: "Binary Search Tree Iterator", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/binary-search-tree-iterator/", PatternID: 14},
		{Number: "230", Title: "Kth Smallest Element in a BST", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/kth-smallest-element-in-a-bst/", PatternID: 14},
		{Number: "501", Title: "Find Mode in Binary Search Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/find-mode-in-binary-search-tree/", PatternID: 14},
		{Number: "530", Title: "Minimum Absolute Difference in BST", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/minimum-absolute-difference-in-bst/", PatternID: 14},

		// Pattern 15: Recursive Postorder Traversal
		{Number: "104", Title: "Maximum Depth of Binary Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/maximum-depth-of-binary-tree/", PatternID: 15},
		{Number: "110", Title: "Balanced Binary Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/balanced-binary-tree/", PatternID: 15},
		{Number: "124", Title: "Binary Tree Maximum Path Sum", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/binary-tree-maximum-path-sum/", PatternID: 15},
		{Number: "145", Title: "Binary Tree Postorder Traversal", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/binary-tree-postorder-traversal/", PatternID: 15},
		{Number: "337", Title: "House Robber III", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/house-robber-iii/", PatternID: 15},
		{Number: "366", Title: "Find Leaves of Binary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-leaves-of-binary-tree/", PatternID: 15},
		{Number: "543", Title: "Diameter of Binary Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/diameter-of-binary-tree/", PatternID: 15},
		{Number: "863", Title: "All Nodes Distance K in Binary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/", PatternID: 15},
		{Number: "1110", Title: "Delete Nodes And Return Forest", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/delete-nodes-and-return-forest/", PatternID: 15},
		{Number: "2458", Title: "Height of Binary Tree After Subtree Removal Queries", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/", PatternID: 15},

		// Pattern 16: Lowest Common Ancestor
		{Number: "235", Title: "Lowest Common Ancestor of a Binary Search Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/", PatternID: 16},
		{Number: "236", Title: "Lowest Common Ancestor of a Binary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/", PatternID: 16},

		// Pattern 17: Serialization and Deserialization
		{Number: "297", Title: "Serialize and Deserialize Binary Tree", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/serialize-and-deserialize-binary-tree/", PatternID: 17},
		{Number: "572", Title: "Subtree of Another Tree", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/subtree-of-another-tree/", PatternID: 17},
		{Number: "652", Title: "Find Duplicate Subtrees", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-duplicate-subtrees/", PatternID: 17},

		// Pattern 18: DFS - Connected Components / Island Counting
		{Number: "130", Title: "Surrounded Regions", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/surrounded-regions/", PatternID: 18},
		{Number: "200", Title: "Number of Islands", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-islands/", PatternID: 18},
		{Number: "417", Title: "Pacific Atlantic Water Flow", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/pacific-atlantic-water-flow/", PatternID: 18},
		{Number: "547", Title: "Number of Provinces", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-provinces/", PatternID: 18},
		{Number: "695", Title: "Max Area of Island", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/max-area-of-island/", PatternID: 18},
		{Number: "733", Title: "Flood Fill", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/flood-fill/", PatternID: 18},
		{Number: "841", Title: "Keys and Rooms", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/keys-and-rooms/", PatternID: 18},
		{Number: "1020", Title: "Number of Enclaves", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-enclaves/", PatternID: 18},
		{Number: "1254", Title: "Number of Closed Islands", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-closed-islands/", PatternID: 18},
		{Number: "1905", Title: "Count Sub Islands", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/count-sub-islands/", PatternID: 18},
		{Number: "2101", Title: "Detonate the Maximum Bombs", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/detonate-the-maximum-bombs/", PatternID: 18},

		// Pattern 19: BFS - Connected Components / Island Counting
		{Number: "542", Title: "01 Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/01-matrix/", PatternID: 19},
		{Number: "994", Title: "Rotting Oranges", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/rotting-oranges/", PatternID: 19},
		{Number: "1091", Title: "Shortest Path in Binary Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/shortest-path-in-binary-matrix/", PatternID: 19},

		// Pattern 20: DFS - Cycle Detection
		{Number: "207", Title: "Course Schedule", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/course-schedule/", PatternID: 20},
		{Number: "210", Title: "Course Schedule II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/course-schedule-ii/", PatternID: 20},
		{Number: "802", Title: "Find Eventual Safe States", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-eventual-safe-states/", PatternID: 20},
		{Number: "1059", Title: "All Paths from Source Lead to Destination", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/all-paths-from-source-lead-to-destination/", PatternID: 20},

		// Pattern 21: BFS - Topological Sort (Kahn's Algorithm)
		{Number: "210", Title: "Course Schedule II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/course-schedule-ii/", PatternID: 21},
		{Number: "269", Title: "Alien Dictionary", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/alien-dictionary/", PatternID: 21},
		{Number: "310", Title: "Minimum Height Trees", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-height-trees/", PatternID: 21},
		{Number: "444", Title: "Sequence Reconstruction", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/sequence-reconstruction/", PatternID: 21},
		{Number: "1136", Title: "Parallel Courses", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/parallel-courses/", PatternID: 21},
		{Number: "1857", Title: "Largest Color Value in a Directed Graph", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/largest-color-value-in-a-directed-graph/", PatternID: 21},
		{Number: "2050", Title: "Parallel Courses III", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/parallel-courses-iii/", PatternID: 21},
		{Number: "2115", Title: "Find All Possible Recipes from Given Supplies", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/", PatternID: 21},
		{Number: "2392", Title: "Build a Matrix With Conditions", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/build-a-matrix-with-conditions/", PatternID: 21},

		// Pattern 22: Deep Copy / Cloning
		{Number: "133", Title: "Clone Graph", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/clone-graph/", PatternID: 22},
		{Number: "138", Title: "Copy List with Random Pointer", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/copy-list-with-random-pointer/", PatternID: 22},
		{Number: "1334", Title: "Find the City With the Smallest Number of Neighbors at a Threshold Distance", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/", PatternID: 22},
		{Number: "1490", Title: "Clone N-ary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/clone-n-ary-tree/", PatternID: 22},

		// Pattern 23: Shortest Path
		{Number: "743", Title: "Network Delay Time", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/network-delay-time/", PatternID: 23},
		{Number: "778", Title: "Swim in Rising Water", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/swim-in-rising-water/", PatternID: 23},
		{Number: "1514", Title: "Path with Maximum Probability", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/path-with-maximum-probability/", PatternID: 23},
		{Number: "1631", Title: "Path With Minimum Effort", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/path-with-minimum-effort/", PatternID: 23},
		{Number: "1976", Title: "Number of Ways to Arrive at Destination", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/", PatternID: 23},
		{Number: "2045", Title: "Second Minimum Time to Reach Destination", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/second-minimum-time-to-reach-destination/", PatternID: 23},
		{Number: "2203", Title: "Minimum Weighted Subgraph With the Required Paths", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/", PatternID: 23},
		{Number: "2290", Title: "Minimum Obstacle Removal to Reach Corner", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/", PatternID: 23},
		{Number: "2577", Title: "Minimum Time to Visit a Cell In a Grid", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/", PatternID: 23},
		{Number: "2812", Title: "Find the Safest Path in a Grid", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-safest-path-in-a-grid/", PatternID: 23},

		// Pattern 24: Shortest Path (Bellman-Ford / BFS+K)
		{Number: "787", Title: "Cheapest Flights Within K Stops", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/cheapest-flights-within-k-stops/", PatternID: 24},
		{Number: "1129", Title: "Shortest Path with Alternating Colors", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/shortest-path-with-alternating-colors/", PatternID: 24},

		// Pattern 25: Union-Find
		{Number: "200", Title: "Number of Islands", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-islands/", PatternID: 25},
		{Number: "261", Title: "Graph Valid Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/graph-valid-tree/", PatternID: 25},
		{Number: "305", Title: "Number of Islands II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/number-of-islands-ii/", PatternID: 25},
		{Number: "323", Title: "Number of Connected Components in an Undirected Graph", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/", PatternID: 25},
		{Number: "547", Title: "Number of Provinces", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-provinces/", PatternID: 25},
		{Number: "684", Title: "Redundant Connection", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/redundant-connection/", PatternID: 25},
		{Number: "721", Title: "Accounts Merge", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/accounts-merge/", PatternID: 25},
		{Number: "737", Title: "Sentence Similarity II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/sentence-similarity-ii/", PatternID: 25},
		{Number: "947", Title: "Most Stones Removed with Same Row or Column", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/", PatternID: 25},
		{Number: "952", Title: "Largest Component Size by Common Factor", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/largest-component-size-by-common-factor/", PatternID: 25},
		{Number: "959", Title: "Regions Cut By Slashes", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/regions-cut-by-slashes/", PatternID: 25},
		{Number: "1101", Title: "The Earliest Moment When Everyone Become Friends", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/", PatternID: 25},

		// Pattern 26: Strongly Connected Components (Kosaraju / Tarjan)
		{Number: "210", Title: "Course Schedule II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/course-schedule-ii/", PatternID: 26},
		{Number: "547", Title: "Number of Provinces", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/number-of-provinces/", PatternID: 26},
		{Number: "1192", Title: "Critical Connections in a Network", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/critical-connections-in-a-network/", PatternID: 26},
		{Number: "2127", Title: "Maximum Employees to Be Invited to a Meeting", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/", PatternID: 26},

		// Pattern 27: Bridges & Articulation Points (Tarjan low-link)
		{Number: "1192", Title: "Critical Connections in a Network", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/critical-connections-in-a-network/", PatternID: 27},
		{Number: "2360", Title: "Longest Cycle in a Graph", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/longest-cycle-in-a-graph/", PatternID: 27},

		// Pattern 28: Minimum Spanning Tree (Kruskal / Prim / DSU + heap)
		{Number: "1135", Title: "Connecting Cities With Minimum Cost", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/connecting-cities-with-minimum-cost/", PatternID: 28},
		{Number: "1168", Title: "Optimize Water Distribution in a Village", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/optimize-water-distribution-in-a-village/", PatternID: 28},
		{Number: "1489", Title: "Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/", PatternID: 28},
		{Number: "1584", Title: "Min Cost to Connect All Points", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/min-cost-to-connect-all-points/", PatternID: 28},

		// Pattern 29: Bidirectional BFS
		{Number: "126", Title: "Word Ladder II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/word-ladder-ii/", PatternID: 29},
		{Number: "127", Title: "Word Ladder", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/word-ladder/", PatternID: 29},
		{Number: "815", Title: "Bus Routes", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/bus-routes/", PatternID: 29},

		// Pattern 30: Fibonacci Style
		{Number: "70", Title: "Climbing Stairs", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/climbing-stairs/", PatternID: 30},
		{Number: "91", Title: "Decode Ways", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/decode-ways/", PatternID: 30},
		{Number: "198", Title: "House Robber", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/house-robber/", PatternID: 30},
		{Number: "213", Title: "House Robber II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/house-robber-ii/", PatternID: 30},
		{Number: "337", Title: "House Robber III", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/house-robber-iii/", PatternID: 30},
		{Number: "509", Title: "Fibonacci Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/fibonacci-number/", PatternID: 30},
		{Number: "740", Title: "Delete and Earn", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/delete-and-earn/", PatternID: 30},
		{Number: "746", Title: "Min Cost Climbing Stairs", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/min-cost-climbing-stairs/", PatternID: 30},

		// Pattern 31: Kadane's Algorithm for Max/Min Subarray
		{Number: "53", Title: "Maximum Subarray", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-subarray/", PatternID: 31},
		{Number: "152", Title: "Maximum Product Subarray", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-product-subarray/", PatternID: 31},
		{Number: "918", Title: "Maximum Sum Circular Subarray", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-sum-circular-subarray/", PatternID: 31},
		{Number: "1749", Title: "Maximum Absolute Sum of Any Subarray", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/", PatternID: 31},
		{Number: "2321", Title: "Maximum Score Of Spliced Array", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/maximum-score-of-spliced-array/", PatternID: 31},

		// Pattern 32: Coin Change / Unbounded Knapsack Style
		{Number: "322", Title: "Coin Change", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/coin-change/", PatternID: 32},
		{Number: "377", Title: "Combination Sum IV", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/combination-sum-iv/", PatternID: 32},
		{Number: "518", Title: "Coin Change II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/coin-change-ii/", PatternID: 32},

		// Pattern 33: 0/1 Knapsack, Subset Sum Style
		{Number: "416", Title: "Partition Equal Subset Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/partition-equal-subset-sum/", PatternID: 33},
		{Number: "494", Title: "Target Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/target-sum/", PatternID: 33},

		// Pattern 34: Word Break Style
		{Number: "139", Title: "Word Break", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/word-break/", PatternID: 34},
		{Number: "140", Title: "Word Break II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/word-break-ii/", PatternID: 34},

		// Pattern 35: Longest Common Subsequence - LCS
		{Number: "1092", Title: "Shortest Common Supersequence", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/shortest-common-supersequence/", PatternID: 35},
		{Number: "1143", Title: "Longest Common Subsequence", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-common-subsequence/", PatternID: 35},
		{Number: "1312", Title: "Minimum Insertion Steps to Make a String Palindrome", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/", PatternID: 35},

		// Pattern 36: Edit Distance / Levenshtein Distance
		{Number: "72", Title: "Edit Distance", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/edit-distance/", PatternID: 36},
		{Number: "583", Title: "Delete Operation for Two Strings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/delete-operation-for-two-strings/", PatternID: 36},
		{Number: "712", Title: "Minimum ASCII Delete Sum for Two Strings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/", PatternID: 36},

		// Pattern 37: Unique Paths on Grid
		{Number: "62", Title: "Unique Paths", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/unique-paths/", PatternID: 37},
		{Number: "63", Title: "Unique Paths II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/unique-paths-ii/", PatternID: 37},
		{Number: "64", Title: "Minimum Path Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-path-sum/", PatternID: 37},
		{Number: "120", Title: "Triangle", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/triangle/", PatternID: 37},
		{Number: "221", Title: "Maximal Square", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximal-square/", PatternID: 37},
		{Number: "931", Title: "Minimum Falling Path Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-falling-path-sum/", PatternID: 37},
		{Number: "1277", Title: "Count Square Submatrices with All Ones", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/count-square-submatrices-with-all-ones/", PatternID: 37},

		// Pattern 38: Interval DP
		{Number: "312", Title: "Burst Balloons", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/burst-balloons/", PatternID: 38},
		{Number: "546", Title: "Remove Boxes", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/remove-boxes/", PatternID: 38},

		// Pattern 39: Catalan Numbers
		{Number: "95", Title: "Unique Binary Search Trees II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/unique-binary-search-trees-ii/", PatternID: 39},
		{Number: "96", Title: "Unique Binary Search Trees", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/unique-binary-search-trees/", PatternID: 39},
		{Number: "241", Title: "Different Ways to Add Parentheses", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/different-ways-to-add-parentheses/", PatternID: 39},

		// Pattern 40: Longest Increasing Subsequence
		{Number: "300", Title: "Longest Increasing Subsequence", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-increasing-subsequence/", PatternID: 40},
		{Number: "354", Title: "Russian Doll Envelopes", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/russian-doll-envelopes/", PatternID: 40},
		{Number: "1671", Title: "Minimum Number of Removals to Make Mountain Array", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/", PatternID: 40},
		{Number: "2407", Title: "Longest Increasing Subsequence II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/longest-increasing-subsequence-ii/", PatternID: 40},

		// Pattern 41: Stock problems
		{Number: "121", Title: "Best Time to Buy and Sell Stock", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/", PatternID: 41},
		{Number: "122", Title: "Best Time to Buy and Sell Stock II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/", PatternID: 41},
		{Number: "123", Title: "Best Time to Buy and Sell Stock III", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/", PatternID: 41},
		{Number: "188", Title: "Best Time to Buy and Sell Stock IV", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/", PatternID: 41},
		{Number: "309", Title: "Best Time to Buy and Sell Stock with Cooldown", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/", PatternID: 41},

		// Pattern 42: Top K Elements
		{Number: "215", Title: "Kth Largest Element in an Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/kth-largest-element-in-an-array/", PatternID: 42},
		{Number: "347", Title: "Top K Frequent Elements", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/top-k-frequent-elements/", PatternID: 42},
		{Number: "451", Title: "Sort Characters By Frequency", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/sort-characters-by-frequency/", PatternID: 42},
		{Number: "506", Title: "Relative Ranks", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/relative-ranks/", PatternID: 42},
		{Number: "703", Title: "Kth Largest Element in a Stream", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/kth-largest-element-in-a-stream/", PatternID: 42},
		{Number: "973", Title: "K Closest Points to Origin", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/k-closest-points-to-origin/", PatternID: 42},
		{Number: "1046", Title: "Last Stone Weight", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/last-stone-weight/", PatternID: 42},
		{Number: "2558", Title: "Take Gifts From the Richest Pile", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/take-gifts-from-the-richest-pile/", PatternID: 42},

		// Pattern 43: Two Heaps for Median Finding
		{Number: "295", Title: "Find Median from Data Stream", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-median-from-data-stream/", PatternID: 43},
		{Number: "1825", Title: "Finding MK Average", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/finding-mk-average/", PatternID: 43},

		// Pattern 44: K-way Merge
		{Number: "23", Title: "Merge k Sorted Lists", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/merge-k-sorted-lists/", PatternID: 44},
		{Number: "373", Title: "Find K Pairs with Smallest Sums", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-k-pairs-with-smallest-sums/", PatternID: 44},
		{Number: "378", Title: "Kth Smallest Element in a Sorted Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/", PatternID: 44},
		{Number: "632", Title: "Smallest Range Covering Elements from K Lists", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/", PatternID: 44},

		// Pattern 45: Scheduling / Minimum Cost
		{Number: "253", Title: "Meeting Rooms II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/meeting-rooms-ii/", PatternID: 45},
		{Number: "767", Title: "Reorganize String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/reorganize-string/", PatternID: 45},
		{Number: "857", Title: "Minimum Cost to Hire K Workers", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimum-cost-to-hire-k-workers/", PatternID: 45},
		{Number: "1642", Title: "Furthest Building You Can Reach", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/furthest-building-you-can-reach/", PatternID: 45},
		{Number: "1792", Title: "Maximum Average Pass Ratio", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-average-pass-ratio/", PatternID: 45},
		{Number: "1834", Title: "Single-Threaded CPU", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/single-threaded-cpu/", PatternID: 45},
		{Number: "1942", Title: "The Number of the Smallest Unoccupied Chair", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/", PatternID: 45},
		{Number: "2402", Title: "Meeting Rooms III", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/meeting-rooms-iii/", PatternID: 45},

		// Pattern 46: Subsets (Include/Exclude)
		{Number: "17", Title: "Letter Combinations of a Phone Number", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/letter-combinations-of-a-phone-number/", PatternID: 46},
		{Number: "77", Title: "Combinations", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/combinations/", PatternID: 46},
		{Number: "78", Title: "Subsets", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/subsets/", PatternID: 46},
		{Number: "90", Title: "Subsets II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/subsets-ii/", PatternID: 46},

		// Pattern 47: Permutations
		{Number: "31", Title: "Next Permutation", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/next-permutation/", PatternID: 47},
		{Number: "46", Title: "Permutations", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/permutations/", PatternID: 47},
		{Number: "60", Title: "Permutation Sequence", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/permutation-sequence/", PatternID: 47},

		// Pattern 48: Combination Sum
		{Number: "39", Title: "Combination Sum", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/combination-sum/", PatternID: 48},
		{Number: "40", Title: "Combination Sum II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/combination-sum-ii/", PatternID: 48},

		// Pattern 49: Parentheses Generation
		{Number: "22", Title: "Generate Parentheses", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/generate-parentheses/", PatternID: 49},
		{Number: "301", Title: "Remove Invalid Parentheses", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/remove-invalid-parentheses/", PatternID: 49},

		// Pattern 50: Word Search / Path Finding in Grid
		{Number: "79", Title: "Word Search", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/word-search/", PatternID: 50},
		{Number: "212", Title: "Word Search II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/word-search-ii/", PatternID: 50},
		{Number: "2018", Title: "Check if Word Can Be Placed In Crossword", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/check-if-word-can-be-placed-in-crossword/", PatternID: 50},

		// Pattern 51: N-Queens / Constraint Satisfaction
		{Number: "37", Title: "Sudoku Solver", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/sudoku-solver/", PatternID: 51},
		{Number: "51", Title: "N-Queens", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/n-queens/", PatternID: 51},

		// Pattern 52: Palindrome Partitioning
		{Number: "131", Title: "Palindrome Partitioning", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/palindrome-partitioning/", PatternID: 52},
		{Number: "132", Title: "Palindrome Partitioning II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/palindrome-partitioning-ii/", PatternID: 52},
		{Number: "1457", Title: "Pseudo-Palindromic Paths in a Binary Tree", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/", PatternID: 52},

		// Pattern 53: Interval Merging/Scheduling
		{Number: "56", Title: "Merge Intervals", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/merge-intervals/", PatternID: 53},
		{Number: "57", Title: "Insert Interval", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/insert-interval/", PatternID: 53},
		{Number: "759", Title: "Employee Free Time", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/employee-free-time/", PatternID: 53},
		{Number: "986", Title: "Interval List Intersections", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/interval-list-intersections/", PatternID: 53},
		{Number: "2406", Title: "Divide Intervals Into Minimum Number of Groups", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/", PatternID: 53},

		// Pattern 54: Jump Game Reachability/Minimization
		{Number: "45", Title: "Jump Game II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/jump-game-ii/", PatternID: 54},
		{Number: "55", Title: "Jump Game", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/jump-game/", PatternID: 54},

		// Pattern 55: Buy/Sell Stock
		{Number: "121", Title: "Best Time to Buy and Sell Stock", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/", PatternID: 55},
		{Number: "122", Title: "Best Time to Buy and Sell Stock II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/", PatternID: 55},

		// Pattern 56: Gas Station Circuit
		{Number: "134", Title: "Gas Station", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/gas-station/", PatternID: 56},
		{Number: "2202", Title: "Maximize the Topmost Element After K Moves", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximize-the-topmost-element-after-k-moves/", PatternID: 56},

		// Pattern 57: Task Scheduling
		{Number: "621", Title: "Task Scheduler", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/task-scheduler/", PatternID: 57},
		{Number: "767", Title: "Reorganize String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/reorganize-string/", PatternID: 57},
		{Number: "1054", Title: "Distant Barcodes", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/distant-barcodes/", PatternID: 57},

		// Pattern 58: Sorting Based
		{Number: "135", Title: "Candy", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/candy/", PatternID: 58},
		{Number: "406", Title: "Queue Reconstruction by Height", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/queue-reconstruction-by-height/", PatternID: 58},
		{Number: "455", Title: "Assign Cookies", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/assign-cookies/", PatternID: 58},
		{Number: "1029", Title: "Two City Scheduling", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/two-city-scheduling/", PatternID: 58},

		// Pattern 59: On Sorted Array/List
		{Number: "35", Title: "Search Insert Position", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/search-insert-position/", PatternID: 59},
		{Number: "69", Title: "Sqrt(x)", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/sqrtx/", PatternID: 59},
		{Number: "74", Title: "Search a 2D Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/search-a-2d-matrix/", PatternID: 59},
		{Number: "278", Title: "First Bad Version", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/first-bad-version/", PatternID: 59},
		{Number: "374", Title: "Guess Number Higher or Lower", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/guess-number-higher-or-lower/", PatternID: 59},
		{Number: "540", Title: "Single Element in a Sorted Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/single-element-in-a-sorted-array/", PatternID: 59},
		{Number: "704", Title: "Binary Search", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/binary-search/", PatternID: 59},
		{Number: "1539", Title: "Kth Missing Positive Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/kth-missing-positive-number/", PatternID: 59},

		// Pattern 60: Find Min/Max in Rotated Sorted Array
		{Number: "33", Title: "Search in Rotated Sorted Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/search-in-rotated-sorted-array/", PatternID: 60},
		{Number: "81", Title: "Search in Rotated Sorted Array II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/search-in-rotated-sorted-array-ii/", PatternID: 60},
		{Number: "153", Title: "Find Minimum in Rotated Sorted Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/", PatternID: 60},
		{Number: "162", Title: "Find Peak Element", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-peak-element/", PatternID: 60},
		{Number: "852", Title: "Peak Index in a Mountain Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/peak-index-in-a-mountain-array/", PatternID: 60},
		{Number: "1095", Title: "Find in Mountain Array", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-in-mountain-array/", PatternID: 60},

		// Pattern 61: On Answer / Condition Function
		{Number: "410", Title: "Split Array Largest Sum", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/split-array-largest-sum/", PatternID: 61},
		{Number: "774", Title: "Minimize Max Distance to Gas Station", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/minimize-max-distance-to-gas-station/", PatternID: 61},
		{Number: "875", Title: "Koko Eating Bananas", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/koko-eating-bananas/", PatternID: 61},
		{Number: "1011", Title: "Capacity To Ship Packages Within D Days", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/", PatternID: 61},
		{Number: "1482", Title: "Minimum Number of Days to Make m Bouquets", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/", PatternID: 61},
		{Number: "1760", Title: "Minimum Limit of Balls in a Bag", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/", PatternID: 61},
		{Number: "2064", Title: "Minimized Maximum of Products Distributed to Any Store", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/", PatternID: 61},
		{Number: "2226", Title: "Maximum Candies Allocated to K Children", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-candies-allocated-to-k-children/", PatternID: 61},

		// Pattern 62: Find First/Last Occurrence
		{Number: "34", Title: "Find First and Last Position of Element in Sorted Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/", PatternID: 62},
		{Number: "658", Title: "Find K Closest Elements", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-k-closest-elements/", PatternID: 62},

		// Pattern 63: Median / Kth across Two Sorted Arrays
		{Number: "4", Title: "Median of Two Sorted Arrays", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/median-of-two-sorted-arrays/", PatternID: 63},
		{Number: "378", Title: "Kth Smallest Element in a Sorted Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/", PatternID: 63},
		{Number: "719", Title: "Find K-th Smallest Pair Distance", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-k-th-smallest-pair-distance/", PatternID: 63},

		// Pattern 64: Valid Parentheses Matching
		{Number: "20", Title: "Valid Parentheses", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/valid-parentheses/", PatternID: 64},
		{Number: "32", Title: "Longest Valid Parentheses", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/longest-valid-parentheses/", PatternID: 64},
		{Number: "921", Title: "Minimum Add to Make Parentheses Valid", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/", PatternID: 64},
		{Number: "1249", Title: "Minimum Remove to Make Valid Parentheses", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/", PatternID: 64},
		{Number: "1963", Title: "Minimum Number of Swaps to Make the String Balanced", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/", PatternID: 64},

		// Pattern 65: Monotonic Stack
		{Number: "402", Title: "Remove K Digits", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/remove-k-digits/", PatternID: 65},
		{Number: "496", Title: "Next Greater Element I", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/next-greater-element-i/", PatternID: 65},
		{Number: "503", Title: "Next Greater Element II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/next-greater-element-ii/", PatternID: 65},
		{Number: "739", Title: "Daily Temperatures", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/daily-temperatures/", PatternID: 65},
		{Number: "901", Title: "Online Stock Span", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/online-stock-span/", PatternID: 65},
		{Number: "907", Title: "Sum of Subarray Minimums", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/sum-of-subarray-minimums/", PatternID: 65},
		{Number: "962", Title: "Maximum Width Ramp", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/maximum-width-ramp/", PatternID: 65},
		{Number: "1475", Title: "Final Prices With a Special Discount in a Shop", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/", PatternID: 65},
		{Number: "1673", Title: "Find the Most Competitive Subsequence", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-most-competitive-subsequence/", PatternID: 65},

		// Pattern 66: Expression Evaluation
		{Number: "150", Title: "Evaluate Reverse Polish Notation", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/evaluate-reverse-polish-notation/", PatternID: 66},
		{Number: "224", Title: "Basic Calculator", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/basic-calculator/", PatternID: 66},
		{Number: "227", Title: "Basic Calculator II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/basic-calculator-ii/", PatternID: 66},
		{Number: "772", Title: "Basic Calculator III", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/basic-calculator-iii/", PatternID: 66},

		// Pattern 67: Simulation / Backtracking Helper
		{Number: "71", Title: "Simplify Path", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/simplify-path/", PatternID: 67},
		{Number: "394", Title: "Decode String", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/decode-string/", PatternID: 67},
		{Number: "735", Title: "Asteroid Collision", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/asteroid-collision/", PatternID: 67},

		// Pattern 68: Min Stack Design
		{Number: "155", Title: "Min Stack", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/min-stack/", PatternID: 68},
		{Number: "895", Title: "Maximum Frequency Stack", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/maximum-frequency-stack/", PatternID: 68},
		{Number: "901", Title: "Online Stock Span", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/online-stock-span/", PatternID: 68},

		// Pattern 69: Largest Rectangle in Histogram
		{Number: "84", Title: "Largest Rectangle in Histogram", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/largest-rectangle-in-histogram/", PatternID: 69},
		{Number: "85", Title: "Maximal Rectangle", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/maximal-rectangle/", PatternID: 69},

		// Pattern 70: Bitwise XOR - Finding Single/Missing Number
		{Number: "136", Title: "Single Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/single-number/", PatternID: 70},
		{Number: "137", Title: "Single Number II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/single-number-ii/", PatternID: 70},
		{Number: "268", Title: "Missing Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/missing-number/", PatternID: 70},
		{Number: "389", Title: "Find the Difference", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/find-the-difference/", PatternID: 70},

		// Pattern 71: Bitwise AND - Counting Set Bits (Hamming Weight)
		{Number: "191", Title: "Number of 1 Bits", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/number-of-1-bits/", PatternID: 71},
		{Number: "231", Title: "Power of Two", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/power-of-two/", PatternID: 71},
		{Number: "477", Title: "Total Hamming Distance", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/total-hamming-distance/", PatternID: 71},

		// Pattern 72: Bitwise DP - Counting Bits Optimization
		{Number: "338", Title: "Counting Bits", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/counting-bits/", PatternID: 72},
		{Number: "1442", Title: "Count Triplets That Can Form Two Arrays of Equal XOR", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/", PatternID: 72},
		{Number: "1494", Title: "Parallel Courses II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/parallel-courses-ii/", PatternID: 72},

		// Pattern 73: Bitwise Operations - Power of Two/Four Check
		{Number: "231", Title: "Power of Two", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/power-of-two/", PatternID: 73},
		{Number: "342", Title: "Power of Four", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/power-of-four/", PatternID: 73},

		// Pattern 74: In-place Reversal
		{Number: "25", Title: "Reverse Nodes in k-Group", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/reverse-nodes-in-k-group/", PatternID: 74},
		{Number: "82", Title: "Remove Duplicates from Sorted List II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/", PatternID: 74},
		{Number: "83", Title: "Remove Duplicates from Sorted List", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/remove-duplicates-from-sorted-list/", PatternID: 74},
		{Number: "92", Title: "Reverse Linked List II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/reverse-linked-list-ii/", PatternID: 74},
		{Number: "206", Title: "Reverse Linked List", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/reverse-linked-list/", PatternID: 74},
		{Number: "234", Title: "Palindrome Linked List", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/palindrome-linked-list/", PatternID: 74},

		// Pattern 75: Merging Two Sorted Lists
		{Number: "21", Title: "Merge Two Sorted Lists", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/merge-two-sorted-lists/", PatternID: 75},
		{Number: "23", Title: "Merge k Sorted Lists", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/merge-k-sorted-lists/", PatternID: 75},

		// Pattern 76: Addition of Numbers
		{Number: "2", Title: "Add Two Numbers", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/add-two-numbers/", PatternID: 76},
		{Number: "369", Title: "Plus One Linked List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/plus-one-linked-list/", PatternID: 76},

		// Pattern 77: Intersection Detection
		{Number: "160", Title: "Intersection of Two Linked Lists", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/intersection-of-two-linked-lists/", PatternID: 77},
		{Number: "599", Title: "Minimum Index Sum of Two Lists", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/minimum-index-sum-of-two-lists/", PatternID: 77},

		// Pattern 78: Reordering / Partitioning
		{Number: "24", Title: "Swap Nodes in Pairs", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/swap-nodes-in-pairs/", PatternID: 78},
		{Number: "61", Title: "Rotate List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/rotate-list/", PatternID: 78},
		{Number: "86", Title: "Partition List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/partition-list/", PatternID: 78},
		{Number: "143", Title: "Reorder List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/reorder-list/", PatternID: 78},
		{Number: "328", Title: "Odd Even Linked List", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/odd-even-linked-list/", PatternID: 78},

		// Pattern 79: In-place Rotation
		{Number: "48", Title: "Rotate Image", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/rotate-image/", PatternID: 79},
		{Number: "189", Title: "Rotate Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/rotate-array/", PatternID: 79},
		{Number: "867", Title: "Transpose Matrix", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/transpose-matrix/", PatternID: 79},

		// Pattern 80: Spiral Traversal
		{Number: "54", Title: "Spiral Matrix", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/spiral-matrix/", PatternID: 80},
		{Number: "59", Title: "Spiral Matrix II", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/spiral-matrix-ii/", PatternID: 80},
		{Number: "885", Title: "Spiral Matrix III", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/spiral-matrix-iii/", PatternID: 80},
		{Number: "2326", Title: "Spiral Matrix IV", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/spiral-matrix-iv/", PatternID: 80},

		// Pattern 81: In-place Marking
		{Number: "73", Title: "Set Matrix Zeroes", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/set-matrix-zeroes/", PatternID: 81},
		{Number: "289", Title: "Game of Life", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/game-of-life/", PatternID: 81},
		{Number: "498", Title: "Diagonal Traverse", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/diagonal-traverse/", PatternID: 81},

		// Pattern 82: Prefix/Suffix Products
		{Number: "238", Title: "Product of Array Except Self", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/product-of-array-except-self/", PatternID: 82},
		{Number: "845", Title: "Longest Mountain in Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-mountain-in-array/", PatternID: 82},
		{Number: "2483", Title: "Minimum Penalty for a Shop", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/minimum-penalty-for-a-shop/", PatternID: 82},

		// Pattern 83: Plus One
		{Number: "43", Title: "Multiply Strings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/multiply-strings/", PatternID: 83},
		{Number: "66", Title: "Plus One", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/plus-one/", PatternID: 83},
		{Number: "67", Title: "Add Binary", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/add-binary/", PatternID: 83},
		{Number: "989", Title: "Add to Array-Form of Integer", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/add-to-array-form-of-integer/", PatternID: 83},

		// Pattern 84: In-place from End
		{Number: "88", Title: "Merge Sorted Array", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/merge-sorted-array/", PatternID: 84},
		{Number: "977", Title: "Squares of a Sorted Array", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/squares-of-a-sorted-array/", PatternID: 84},

		// Pattern 85: Cyclic Sort
		{Number: "41", Title: "First Missing Positive", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/first-missing-positive/", PatternID: 85},
		{Number: "268", Title: "Missing Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/missing-number/", PatternID: 85},
		{Number: "287", Title: "Find the Duplicate Number", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-the-duplicate-number/", PatternID: 85},
		{Number: "442", Title: "Find All Duplicates in an Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/find-all-duplicates-in-an-array/", PatternID: 85},
		{Number: "448", Title: "Find All Numbers Disappeared in an Array", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/", PatternID: 85},

		// Pattern 86: Palindrome Check
		{Number: "9", Title: "Palindrome Number", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/palindrome-number/", PatternID: 86},
		{Number: "125", Title: "Valid Palindrome", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/valid-palindrome/", PatternID: 86},
		{Number: "680", Title: "Valid Palindrome II", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/valid-palindrome-ii/", PatternID: 86},

		// Pattern 87: Anagram Check
		{Number: "49", Title: "Group Anagrams", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/group-anagrams/", PatternID: 87},
		{Number: "242", Title: "Valid Anagram", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/valid-anagram/", PatternID: 87},

		// Pattern 88: Roman to Integer Conversion
		{Number: "12", Title: "Integer to Roman", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/integer-to-roman/", PatternID: 88},
		{Number: "13", Title: "Roman to Integer", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/roman-to-integer/", PatternID: 88},

		// Pattern 89: String to Integer (atoi)
		{Number: "8", Title: "String to Integer (atoi)", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/string-to-integer-atoi/", PatternID: 89},
		{Number: "65", Title: "Valid Number", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/valid-number/", PatternID: 89},

		// Pattern 90: Manual Simulation
		{Number: "43", Title: "Multiply Strings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/multiply-strings/", PatternID: 90},
		{Number: "67", Title: "Add Binary", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/add-binary/", PatternID: 90},
		{Number: "415", Title: "Add Strings", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/add-strings/", PatternID: 90},

		// Pattern 91: String Matching - Naive / KMP / Rabin-Karp
		{Number: "28", Title: "Find the Index of the First Occurrence in a String", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/", PatternID: 91},
		{Number: "214", Title: "Shortest Palindrome", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/shortest-palindrome/", PatternID: 91},
		{Number: "686", Title: "Repeated String Match", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/repeated-string-match/", PatternID: 91},
		{Number: "796", Title: "Rotate String", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/rotate-string/", PatternID: 91},
		{Number: "3008", Title: "Find Beautiful Indices in the Given Array II", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/", PatternID: 91},

		// Pattern 92: Repeated Substring Pattern Detection
		{Number: "28", Title: "Find the Index of the First Occurrence in a String", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/", PatternID: 92},
		{Number: "459", Title: "Repeated Substring Pattern", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/repeated-substring-pattern/", PatternID: 92},
		{Number: "686", Title: "Repeated String Match", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/repeated-string-match/", PatternID: 92},

		// Pattern 93: Design (General/Specific)
		{Number: "146", Title: "LRU Cache", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/lru-cache/", PatternID: 93},
		{Number: "155", Title: "Min Stack", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/min-stack/", PatternID: 93},
		{Number: "225", Title: "Implement Stack using Queues", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/implement-stack-using-queues/", PatternID: 93},
		{Number: "232", Title: "Implement Queue using Stacks", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/implement-queue-using-stacks/", PatternID: 93},
		{Number: "251", Title: "Flatten 2D Vector", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/flatten-2d-vector/", PatternID: 93},
		{Number: "271", Title: "Encode and Decode Strings", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/encode-and-decode-strings/", PatternID: 93},
		{Number: "295", Title: "Find Median from Data Stream", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/find-median-from-data-stream/", PatternID: 93},
		{Number: "341", Title: "Flatten Nested List Iterator", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/flatten-nested-list-iterator/", PatternID: 93},
		{Number: "346", Title: "Moving Average from Data Stream", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/moving-average-from-data-stream/", PatternID: 93},
		{Number: "353", Title: "Design Snake Game", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-snake-game/", PatternID: 93},
		{Number: "359", Title: "Logger Rate Limiter", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/logger-rate-limiter/", PatternID: 93},
		{Number: "362", Title: "Design Hit Counter", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-hit-counter/", PatternID: 93},
		{Number: "379", Title: "Design Phone Directory", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-phone-directory/", PatternID: 93},
		{Number: "380", Title: "Insert Delete GetRandom O(1)", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/insert-delete-getrandom-o1/", PatternID: 93},
		{Number: "432", Title: "All O`one Data Structure", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/all-oone-data-structure/", PatternID: 93},
		{Number: "460", Title: "LFU Cache", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/lfu-cache/", PatternID: 93},
		{Number: "604", Title: "Design Compressed String Iterator", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/design-compressed-string-iterator/", PatternID: 93},
		{Number: "622", Title: "Design Circular Queue", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-circular-queue/", PatternID: 93},
		{Number: "641", Title: "Design Circular Deque", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-circular-deque/", PatternID: 93},
		{Number: "642", Title: "Design Search Autocomplete System", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/design-search-autocomplete-system/", PatternID: 93},
		{Number: "706", Title: "Design HashMap", Difficulty: "Easy", LeetCodeURL: "https://leetcode.com/problems/design-hashmap/", PatternID: 93},
		{Number: "715", Title: "Range Module", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/range-module/", PatternID: 93},
		{Number: "900", Title: "RLE Iterator", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/rle-iterator/", PatternID: 93},
		{Number: "981", Title: "Time Based Key-Value Store", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/time-based-key-value-store/", PatternID: 93},
		{Number: "1146", Title: "Snapshot Array", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/snapshot-array/", PatternID: 93},
		{Number: "1348", Title: "Tweet Counts Per Frequency", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/tweet-counts-per-frequency/", PatternID: 93},
		{Number: "1352", Title: "Product of the Last K Numbers", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/product-of-the-last-k-numbers/", PatternID: 93},
		{Number: "1381", Title: "Design a Stack With Increment Operation", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-a-stack-with-increment-operation/", PatternID: 93},
		{Number: "1756", Title: "Design Most Recently Used Queue", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-most-recently-used-queue/", PatternID: 93},
		{Number: "2013", Title: "Detect Squares", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/detect-squares/", PatternID: 93},
		{Number: "2034", Title: "Stock Price Fluctuation", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/stock-price-fluctuation/", PatternID: 93},
		{Number: "2296", Title: "Design a Text Editor", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/design-a-text-editor/", PatternID: 93},
		{Number: "2336", Title: "Smallest Number in Infinite Set", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/smallest-number-in-infinite-set/", PatternID: 93},

		// Pattern 94: Tries
		{Number: "208", Title: "Implement Trie (Prefix Tree)", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/implement-trie-prefix-tree/", PatternID: 94},
		{Number: "211", Title: "Design Add and Search Words Data Structure", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/design-add-and-search-words-data-structure/", PatternID: 94},
		{Number: "425", Title: "Word Squares", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/word-squares/", PatternID: 94},
		{Number: "642", Title: "Design Search Autocomplete System", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/design-search-autocomplete-system/", PatternID: 94},
		{Number: "648", Title: "Replace Words", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/replace-words/", PatternID: 94},
		{Number: "720", Title: "Longest Word in Dictionary", Difficulty: "Medium", LeetCodeURL: "https://leetcode.com/problems/longest-word-in-dictionary/", PatternID: 94},
		{Number: "745", Title: "Prefix and Suffix Search", Difficulty: "Hard", LeetCodeURL: "https://leetcode.com/problems/prefix-and-suffix-search/", PatternID: 94},
	}
}

// Helper function to parse problem number from string like "11. Container With Most Water"
func parseProblemString(s string) (number, title string) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 2 {
		number = strings.TrimSpace(parts[0])
		title = strings.TrimSpace(parts[1])
	}
	return
}
