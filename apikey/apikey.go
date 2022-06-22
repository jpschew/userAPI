package apikey

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// generateTokenSHA256 returns a unique token based on the provided input
// using SHA256 to hash a random 32 bytes of data
func generateTokenSHA256(input []byte) string {

	//hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Hash to store:", string(hash))

	hasher := sha256.New()
	hasher.Write(input)

	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateAPIKey generates a new API Key using 32 bytes of random data with the SHA256 hashing algorithm.
func GenerateAPIKey() string {
	// set the length to 32
	keyLen := 32
	// declare b as slice of 32 bytes of data
	apiKey := make([]byte, keyLen)
	// randomly assign 32 bytes of data to b
	_, err := rand.Read(apiKey)
	if err != nil {
		return fmt.Sprintln("error:", err)
	}

	return generateTokenSHA256(apiKey)
}
