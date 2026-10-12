package domain

import "time"

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

type User struct {
	ID        int64
	Email     string
	Username  string
	Role      Role
	EmailVerifiedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCredentials struct {
	UserID       int64
	PasswordHash string
}
