package ports

type PasswordHasher interface {
	HashPassword(plain string) (string, error)
    VerifyPassword(plain, hash []byte) (bool, error)
}