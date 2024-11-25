package auth_test

import (
	"testing"

	"github.com/otie173/odncore/internal/auth"
)

func TestAuth(t *testing.T) {
	const (
		testPassword  string = "qwerty123"
		wrongPassword string = "qwerty321"
		emptyPassword string = ""
	)
	hash, err := auth.GenerateHash(testPassword)
	if err != nil {
		t.Fatal("Failed to generate password hash: ", err)
	}
	hashedPassword := string(hash)

	t.Run("Authentication Tests", func(t *testing.T) {
		t.Run("Generate Hash", func(t *testing.T) {
			newHash, err := auth.GenerateHash(testPassword)
			if err != nil {
				t.Fatal("Failed to generate hash: ", err)
			}

			if string(newHash) == "" {
				t.Error("Hash cant be empty")
			}
		})

		t.Run("Verify Correct Password", func(t *testing.T) {
			err := auth.CheckPassword(testPassword, hashedPassword)
			if err != nil {
				t.Error("Valid password must pass check: ", err)
			}
		})

		t.Run("Verify Wrong Password", func(t *testing.T) {
			if err := auth.CheckPassword(wrongPassword, hashedPassword); err == nil {
				t.Error("Wrong password must not pass check: ")
			}
		})

		t.Run("Empty Password", func(t *testing.T) {
			if _, err := auth.GenerateHash(emptyPassword); err == nil {
				t.Error("Empty password must return error")
			}
		})
	})
}
