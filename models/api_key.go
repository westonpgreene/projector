package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type APIKey struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Label     string         `json:"label"`
	KeyHash   string         `gorm:"uniqueIndex" json:"-"`
	Active    bool           `json:"active"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func GenerateKey() (plaintext string, hash string, err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return "", "", fmt.Errorf("failed to generate key: %w", err)
	}
	plaintext = "projector_" + hex.EncodeToString(buf)
	hash = HashKey(plaintext)
	return plaintext, hash, nil
}

func HashKey(plaintext string) string {
	sum := sha256.Sum256([]byte(plaintext))
	return hex.EncodeToString(sum[:])
}
