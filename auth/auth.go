package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Function to hash a password
func HashPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 10
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	// Return the hashed password as a string
	return string(salt), nil
}

// Function to verify a password
func VerifyPassword(password, hashedPassword string) bool {
	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
