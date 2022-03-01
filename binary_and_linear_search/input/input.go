package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/D7682/binary_and_linear_search/colors"
)

// this function reads input, gets rid of spaces, and converts a string
// into an array of integers type.
func Read() ([]int, int, string) {
	colors.Notification.Println("Enter the desired values in comma separated format: ")
	colors.Notification.Println("101, 142, 147, 189, 199, 207, 222, 234, 289, 296, etc. (No comma on last value)")
	fmt.Println("Press enter when finished.")
	var values []int
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	trimSpace := strings.TrimSpace(input)

	for _, val := range strings.Split(trimSpace, ", ") {
		// convert ASCII to integer
		toInt, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, toInt)
	}

	colors.Question.Println("Enter the number you want to look for in the sorted array.")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	trimSpace = strings.TrimSpace(input)
	target, err := strconv.Atoi(trimSpace)
	if err != nil {
		log.Fatal(err)
	}

	colors.Question.Println("What type of search do you want to execute \"binary\" or \"linear\"")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	searchType := strings.TrimSpace(input)

	return values, target, searchType
}
