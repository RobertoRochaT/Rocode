package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Database connection
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/rocode?sslmode=disable"
	}

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")

	// Initialize database tables
	if err = initDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	// Allow specific origins (no paths). Include your Vercel frontend and localhost dev origins.
	config.AllowOrigins = []string{
		"https://rocode-production.up.railway.app",
		"https://rocode-rose.vercel.app",
		"http://localhost:3000",
		"http://localhost:5173",
	}
	// Allow common HTTP methods including OPTIONS for preflight
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	// Optionally expose Authorization header to the browser
	config.ExposeHeaders = []string{"Authorization"}
	r.Use(cors.New(config))

	// Routes
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/register", registerHandler)
		api.POST("/login", loginHandler)

		// Protected routes
		protected := api.Group("/")
		protected.Use(authMiddleware())
		{
			protected.GET("/progress", getProgressHandler)
			protected.POST("/progress", updateProgressHandler)
			protected.GET("/stats", getStatsHandler)
		}

		// Public routes with optional auth (to include user progress if logged in)
		optionalAuth := api.Group("/")
		optionalAuth.Use(optionalAuthMiddleware())
		{
			optionalAuth.GET("/patterns", getPatternsHandler)
			optionalAuth.GET("/problems", getProblemsHandler)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDB() error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		username VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS problem_progress (
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

	CREATE INDEX IF NOT EXISTS idx_user_progress ON problem_progress(user_id);
	CREATE INDEX IF NOT EXISTS idx_pattern_progress ON problem_progress(pattern_id, user_id);
	`

	_, err := db.Exec(schema)
	return err
}
