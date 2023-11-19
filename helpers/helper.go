package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadTextFile(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if info.IsDir() {
		fmt.Println("Error: ", filePath, " is a directory")
		os.Exit(1)
	}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return fmt.Sprintf("%8d", lineCount)
}
