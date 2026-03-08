# 🚀 RoCode - LeetCode Pattern Roadmap

A beautiful, interactive platform to master coding interview patterns systematically. Built with React, Go, and PostgreSQL.

## ✨ Features

- 🎯 **94 Coding Patterns** - Comprehensive coverage of all essential patterns
- 📊 **Visual Roadmap** - Interactive pattern cards with progress tracking
- ✅ **Progress Tracking** - Mark problems as complete and add personal notes
- 📈 **Statistics Dashboard** - Monitor your learning journey with detailed analytics
- 🔗 **Direct LeetCode Integration** - Click to open problems directly on LeetCode
- 🎨 **Beautiful UI** - Modern, dark-themed interface with smooth animations
- 🔐 **User Authentication** - Secure JWT-based authentication system
- 📱 **Responsive Design** - Works perfectly on desktop, tablet, and mobile

## 🏗️ Tech Stack

### Frontend
- **React 18** - UI framework
- **React Router** - Client-side routing
- **TailwindCSS** - Utility-first styling
- **Vite** - Build tool and dev server
- **Axios** - HTTP client
- **Lucide React** - Beautiful icons

### Backend
- **Go 1.21+** - High-performance backend
- **Gin** - Web framework
- **PostgreSQL** - Reliable database
- **JWT** - Secure authentication
- **Bcrypt** - Password hashing

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose (recommended)
- OR manually: Node.js 18+, Go 1.21+, PostgreSQL 15+

### Option 1: Docker (Recommended)

1. Clone the repository:
```bash
git clone <your-repo-url>
cd rocode
```

2. Start all services:
```bash
docker-compose up -d
```

3. Access the application:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

That's it! The database will be automatically initialized.

### Option 2: Manual Setup

#### Backend Setup

1. Navigate to backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Set environment variables:
```bash
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/rocode?sslmode=disable"
export JWT_SECRET="your-secret-key-change-in-production"
export PORT="8080"
```

4. Run the backend:
```bash
go run .
```

The backend will start on http://localhost:8080

#### Frontend Setup

1. Navigate to frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Create `.env` file:
```bash
VITE_API_URL=http://localhost:8080/api
```

4. Run the development server:
```bash
npm run dev
```

The frontend will start on http://localhost:3000

## 📖 Usage Guide

### Getting Started

1. **Create an Account**
   - Click "Get Started" on the home page
   - Fill in your email, username, and password
   - You'll be automatically logged in

2. **Explore the Roadmap**
   - Browse through 94 patterns organized by category
   - Use filters to find specific pattern types
   - Search for patterns by name or description

3. **Work on Problems**
   - Click on any pattern card to see its problems
   - Click the external link icon to open problems on LeetCode
   - Mark problems as complete by clicking the checkbox
   - Add notes to remember your approach or insights

4. **Track Your Progress**
   - Visit the Dashboard to see your statistics
   - View completion rates by pattern category
   - See your recent activity and achievements

## 🎯 Pattern Categories

The platform covers 15 major categories:

1. **Two Pointer Patterns** (7 patterns)
   - Converging, Fast & Slow, Fixed Separation, etc.

2. **Sliding Window Patterns** (4 patterns)
   - Fixed Size, Variable Size, Monotonic Queue, etc.

3. **Tree Traversal Patterns** (6 patterns)
   - Level Order, Preorder, Inorder, Postorder, etc.

4. **Graph Traversal Patterns** (12 patterns)
   - DFS, BFS, Union-Find, MST, Shortest Path, etc.

5. **Dynamic Programming** (12 patterns)
   - Fibonacci, Kadane's, LCS, Knapsack, etc.

6. **Heap (Priority Queue)** (4 patterns)
   - Top K Elements, Two Heaps, K-way Merge, etc.

7. **Backtracking** (7 patterns)
   - Subsets, Permutations, N-Queens, etc.

8. **Greedy** (6 patterns)
   - Interval Scheduling, Jump Game, etc.

9. **Binary Search** (5 patterns)
   - On Sorted Array, Rotated Array, etc.

10. **Stack** (6 patterns)
    - Monotonic Stack, Expression Evaluation, etc.

11. **Bit Manipulation** (4 patterns)
    - XOR, Counting Bits, Power of Two, etc.

12. **Linked List** (5 patterns)
    - Reversal, Merging, Reordering, etc.

13. **Array/Matrix** (7 patterns)
    - Rotation, Spiral Traversal, Cyclic Sort, etc.

14. **String Manipulation** (7 patterns)
    - Palindrome, Anagram, Pattern Matching, etc.

15. **Design Patterns** (2 patterns)
    - Data Structure Design, Tries

## 🔌 API Endpoints

### Authentication
- `POST /api/register` - Create new user account
- `POST /api/login` - Login and get JWT token

### Patterns & Problems
- `GET /api/patterns` - Get all coding patterns
- `GET /api/problems` - Get all problems (optional filter by pattern_id)

### Progress (Protected)
- `GET /api/progress` - Get user's progress
- `POST /api/progress` - Update problem completion status
- `GET /api/stats` - Get user statistics

## 📁 Project Structure

```
rocode/
├── backend/                 # Go backend
│   ├── main.go             # Entry point & DB setup
│   ├── auth.go             # Authentication logic
│   ├── handlers.go         # API handlers
│   ├── problems.go         # Problem data
│   ├── go.mod              # Dependencies
│   ├── Dockerfile          # Backend Docker config
│   └── README.md           # Backend documentation
│
├── frontend/               # React frontend
│   ├── src/
│   │   ├── components/     # Reusable components
│   │   │   ├── Navbar.jsx
│   │   │   ├── PatternCard.jsx
│   │   │   └── ProblemList.jsx
│   │   ├── pages/          # Page components
│   │   │   ├── Home.jsx
│   │   │   ├── Login.jsx
│   │   │   ├── Register.jsx
│   │   │   ├── Roadmap.jsx
│   │   │   └── Dashboard.jsx
│   │   ├── contexts/       # React contexts
│   │   │   └── AuthContext.jsx
│   │   ├── services/       # API services
│   │   │   └── api.js
│   │   ├── App.jsx         # Main app component
│   │   ├── main.jsx        # Entry point
│   │   └── index.css       # Global styles
│   ├── package.json        # Dependencies
│   ├── vite.config.js      # Vite configuration
│   ├── tailwind.config.js  # Tailwind configuration
│   ├── Dockerfile          # Frontend Docker config
│   └── nginx.conf          # Nginx configuration
│
├── docker-compose.yml      # Full stack setup
└── README.md               # This file
└── start.sh                # Start script
```

## 🗄️ Database Schema

### users
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### problem_progress
```sql
CREATE TABLE problem_progress (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    problem_number VARCHAR(50) NOT NULL,
    pattern_id INTEGER NOT NULL,
    completed BOOLEAN DEFAULT false,
    difficulty VARCHAR(20),
    notes TEXT,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, problem_number)
);
```

## 🛠️ Development

### Frontend Development

```bash
cd frontend
npm install
npm run dev
```

Visit http://localhost:3000

### Backend Development

```bash
cd backend
go run .
```

API available at http://localhost:8080

### Building for Production

Frontend:
```bash
cd frontend
npm run build
```

Backend:
```bash
cd backend
go build -o rocode
./rocode
```

## 🔒 Environment Variables

### Backend
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT signing
- `PORT` - Server port (default: 8080)

### Frontend
- `VITE_API_URL` - Backend API URL

## 🙏 Acknowledgments

- LeetCode for providing the problems
- NeetCode for inspiration on the pattern-based approach
- All the open-source libraries used in this project

## 📧 Contact

For questions or suggestions, please open an issue on GitHub.

---

**Happy Coding! 🚀**

Master these patterns and ace your coding interviews!
