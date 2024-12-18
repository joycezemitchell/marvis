package common

import (
	"log"
	"os"
)

func SaveToTextFile(message, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
		log.Fatalf("Failed to create or open file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
		log.Fatalf("Failed to write to file: %s", err)
	}

	log.Println("File content overwritten successfully")
	return nil
}
