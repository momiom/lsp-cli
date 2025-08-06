// Package main provides the entry point for the test application
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/example/testproject/user"
	"github.com/example/testproject/utils"
)

// Global constants
const (
	// AppVersion represents the current version of the application
	AppVersion = "1.0.0"
	// DefaultTimeout is the default timeout for operations
	DefaultTimeout = 30 * time.Second
)

// Global variables
var (
	// logger is the global logger instance
	logger *log.Logger
	// startTime records when the application started
	startTime time.Time
)

// Config represents the application configuration
type Config struct {
	Host     string
	Port     int
	Debug    bool
	Timeout  time.Duration
	Features []string
}

// init initializes the application
func init() {
	startTime = time.Now()
	logger = log.New(log.Writer(), "[MAIN] ", log.LstdFlags)
}

// main is the entry point of the application
func main() {
	fmt.Printf("Test Application v%s\n", AppVersion)
	
	// Create configuration
	config := &Config{
		Host:    "localhost",
		Port:    8080,
		Debug:   true,
		Timeout: DefaultTimeout,
		Features: []string{"auth", "api", "metrics"},
	}
	
	// Create user service
	userService := user.NewUserService()
	
	// Create sample users
	admin := &user.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@example.com",
		Role:     user.RoleAdmin,
		Status:   user.StatusActive,
	}
	
	// Add user
	if err := userService.CreateUser(admin); err != nil {
		logger.Fatalf("Failed to create admin user: %v", err)
	}
	
	// Test utilities
	result := utils.Calculate(10, 20, utils.OperationAdd)
	fmt.Printf("Calculation result: %v\n", result)
	
	// Run server
	runServer(config)
}

// runServer starts the application server
func runServer(config *Config) error {
	logger.Printf("Starting server on %s:%d", config.Host, config.Port)
	
	// Server implementation would go here
	
	return nil
}

// validateConfig validates the configuration
func validateConfig(config *Config) error {
	if config.Host == "" {
		return fmt.Errorf("host cannot be empty")
	}
	if config.Port <= 0 {
		return fmt.Errorf("port must be positive")
	}
	return nil
}