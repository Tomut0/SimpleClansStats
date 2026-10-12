package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type ArgonHasher struct {
	timeCost    uint32
	memoryCost  uint32
	parallelism uint8
}

func NewArgonHasher(timeCost, memoryCost uint32, parallelism uint8) *ArgonHasher {
	return &ArgonHasher{
		timeCost:    timeCost,
		memoryCost:  memoryCost,
		parallelism: parallelism,
	}
}

func generateSalt(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

func (h *ArgonHasher) HashPassword(password string) (string, error) {
	salt := generateSalt(16)

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		h.timeCost,
		h.memoryCost,
		h.parallelism,
		32,
	)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.memoryCost,
		h.timeCost,
		h.parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

func (h *ArgonHasher) VerifyPassword(password, hash []byte) (bool, error) {
	// 1. Parse the encoded hash string
	// The expected format is: $argon2id$v=19$m=65536,t=1,p=4$<salt>$<hash>
	parts := strings.Split(string(hash), "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// Verify algorithm identifier
	if parts[1] != "argon2id" {
		return false, errors.New("incompatible variant")
	}

	// Extract version (not strictly needed for argon2.IDKey, but good to validate)
	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, errors.New("invalid version format")
	}

	// Extract parameters (memory, iterations, parallelism)
	var memory, iterations uint32
	var parallelism uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism); err != nil {
		return false, errors.New("invalid parameter format")
	}

	// Decode the salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, errors.New("invalid salt encoding")
	}

	// Decode the original hash
	originalHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, errors.New("invalid hash encoding")
	}

	// Determine the key length from the decoded hash
	keyLen := uint32(len(originalHash))

	// 2. Re-hash the provided password with the extracted parameters
	computedHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLen)

	// 3. Compare the hashes in constant time
	// subtle.ConstantTimeCompare returns 1 if the byte slices are equal, 0 otherwise.
	if subtle.ConstantTimeCompare(originalHash, computedHash) == 1 {
		return true, nil
	}

	return false, nil
}
