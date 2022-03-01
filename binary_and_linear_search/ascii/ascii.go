package ascii

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Art() {
	current, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(current + "/ascii/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Printf("\n")
}
