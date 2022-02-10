package hash_provider

import "golang.org/x/crypto/bcrypt"

func Hash(hashString string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(hashString), 14)

	result := string(bytes)

	return result, err
}

func Compare(hashString string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(hashString))
	return err == nil
}
