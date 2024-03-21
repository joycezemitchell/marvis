package common

import(
    "strings"
    "crypto/rand"
    "encoding/hex"
)

func CheckWordsExist(target string, msg string) bool {
    words := strings.Fields(msg)

    // Tokenize the target string into words
	tokenizedTarget := strings.Fields(target)
	tokenMap := make(map[string]bool)

	// Create a map for easy lookup
	for _, word := range tokenizedTarget {
		tokenMap[word] = true
	}

	// Check if each word exists in the tokenized target
	for _, word := range words {
		if _, exists := tokenMap[word]; !exists {
			return false
		}
	}

	return true
}


func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
