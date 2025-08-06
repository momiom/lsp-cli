// Package user provides user management functionality
package user

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Role represents a user role
type Role int

// Role constants
const (
	RoleGuest Role = iota
	RoleUser
	RoleModerator
	RoleAdmin
)

// Status represents a user status
type Status string

// Status constants
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBanned   Status = "banned"
	StatusPending  Status = "pending"
)

// Errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidRole       = errors.New("invalid role")
)

// User represents a system user
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Hidden in JSON
	Role      Role      `json:"role"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Profile information
	Profile *UserProfile `json:"profile,omitempty"`
}

// UserProfile contains additional user information
type UserProfile struct {
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Bio         string            `json:"bio"`
	AvatarURL   string            `json:"avatar_url"`
	Preferences map[string]string `json:"preferences"`
}

// IsAdmin checks if the user has admin role
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// FullName returns the user's full name
func (u *User) FullName() string {
	if u.Profile != nil {
		return fmt.Sprintf("%s %s", u.Profile.FirstName, u.Profile.LastName)
	}
	return u.Username
}

// Validate validates the user data
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	return nil
}

// UserService provides user management operations
type UserService struct {
	mu    sync.RWMutex
	users map[int64]*User
	idSeq int64
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		users: make(map[int64]*User),
		idSeq: 0,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Check if user exists
	for _, u := range s.users {
		if u.Email == user.Email {
			return ErrUserAlreadyExists
		}
	}
	
	// Assign ID and timestamps
	s.idSeq++
	user.ID = s.idSeq
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	
	s.users[user.ID] = user
	return nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int64) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	user, exists := s.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	
	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int64, updates *User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	user, exists := s.users[id]
	if !exists {
		return ErrUserNotFound
	}
	
	// Update fields
	if updates.Username != "" {
		user.Username = updates.Username
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.Status != "" {
		user.Status = updates.Status
	}
	
	user.UpdatedAt = time.Now()
	return nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.users[id]; !exists {
		return ErrUserNotFound
	}
	
	delete(s.users, id)
	return nil
}

// ListUsers returns all users
func (s *UserService) ListUsers() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	
	return users
}

// UserRepository defines the interface for user persistence
type UserRepository interface {
	Create(user *User) error
	Get(id int64) (*User, error)
	Update(user *User) error
	Delete(id int64) error
	FindByEmail(email string) (*User, error)
	List(limit, offset int) ([]*User, error)
}