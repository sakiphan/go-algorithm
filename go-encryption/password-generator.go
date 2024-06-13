package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	password, err := generatePassword(10)
	if err != nil {
		fmt.Println("Şifre üretiminde hata:", err)
		return
	}
	fmt.Println("Password:", password)
}

// generatePassword, belirtilen uzunlukta kriptografik olarak güvenli bir şifre döndürür
func generatePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}
