package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintextPassword string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return hashedPassword
}
