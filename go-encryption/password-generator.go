package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func main() {
	password, err := generatePassword(15, 20)
	if err != nil {
		fmt.Println("Şifre üretiminde hata:", err)
		return
	}
	fmt.Println("Password:", password)
}

// generatePassword, belirtilen uzunlukta kriptografik olarak güvenli bir şifre döndürür
func generatePassword(minLength, maxLength int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength+1)))
	if err != nil {
		return "", err
	}

	finalLength := minLength + int(length.Int64())
	bytes := make([]byte, finalLength)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}
