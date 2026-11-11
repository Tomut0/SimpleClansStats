package domain

import "time"

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

type User struct {
	email           string
	username        string
	role            Role
	emailVerifiedAt *time.Time
	createdAt       time.Time
	updatedAt       time.Time
}

func NewUser(
	email, username string,
	role Role,
	emailVerifiedAt *time.Time,
	createdAt, updatedAt time.Time,
) *User {
	return &User{
		email:           email,
		username:        username,
		role:            role,
		emailVerifiedAt: emailVerifiedAt,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
}

func (u *User) Email() string    { return u.email }
func (u *User) Username() string { return u.username }
func (u *User) Role() Role       { return u.role }

type UserCredentials struct {
	UserID       int64
	PasswordHash string
}
