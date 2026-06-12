package models

import "time"

// Role represents different roles for level of access
type Role struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// User represents users
type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` // Скрываем хэш пароля из JSON-ответов
	RoleID       int       `json:"role_id" db:"role_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// RegisterInput Validation schema for User registration request
type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginInput Validation schema for User login requests
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse represents the response for authentication requests
type AuthResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}

// UserResponse represents the user data returned in API responses, excluding sensitive fields
type UserResponse struct {
	ID        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"User_Name"`
	Email     string    `json:"email" example:"user@example.com"`
	RoleID    int       `json:"role_id" example:"2"`
	CreatedAt time.Time `json:"created_at" example:"2026-06-12T16:24:54Z"`
}
