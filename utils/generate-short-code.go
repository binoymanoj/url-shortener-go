package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func GenerateShortCode(url string) string {
	hash := md5.Sum([]byte(url))
	hashStr := hex.EncodeToString(hash[:])
	log.Printf("hashStr: %v", hashStr)
	return hashStr[:6]
}
