package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/chethan0707/gowc/internals/interfaces"
)

func ReadTextFile(filePath string) interfaces.Result {
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

	return interfaces.Result{Lines: lineCount}
}
