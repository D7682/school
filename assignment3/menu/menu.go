package menu

import (
	"errors"
	"fmt"
	"log"
)

func Init() (int, error) {
	fmt.Printf("\n\n")
	fmt.Println("Grade Report Program")
	options := []string{
		"1. Display Average Grade",
		"2. Display Maximum Grade",
		"3. Display Minimum Grade",
		"4. Quit Program",
	}

	for _, val := range options {
		fmt.Printf("\n%10v\n", val)
	}

	fmt.Printf("\nEnter your choice (1-4): \n")

	var choice int
	// I don't need the first return value of the fmt.Scanln() function
	// but I use the second one in case an integer type is not entered it will end the program as well.
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}

	// this error handler will return a 4 value which will make the program exit if the value is higher than 4 or lower than 0.
	if choice > 4 || choice < 0 {
		return 4, errors.New("the value you enter has to be between 0-4; please restart the program")
	}

	// if the choice is between 0-4 then the program will return an integer and a nil value for the error meaning there were no errors.
	return choice, nil
}
