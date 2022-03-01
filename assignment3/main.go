package main

import (
	"log"

	"github.com/D7682/assignment3/menu"
	"github.com/D7682/assignment3/opts"
)

// int is the return value of Init
func Init() int {
	// I start by declaring a variable with the memory address of a type that I created called Students.
	s := &opts.Students{}

	// I created Methods for this struct type and one of the methods is the ReadFile method.
	// The ReadFile method will open the file and put values into the variable s of the custom type Students.
	// In Go, I don't have to specify (*s).ReadFile explicitly.
	s.ReadFile()

	// After the method above is done I execute the menu.Init() function which has 2 return values: 1) the integer for the choice of the user, 2) The error handler I provide it.
	// if err is not nil(empty) the the program will output the error and close the program with log.Fatal(err).
	choice, err := menu.Init()
	if err != nil {
		log.Fatal(err)
	}

	// if the error above is nil(empty) meaning there are no errors then choice will be assigned the number that the user input
	// and it will be compared to these cases
	// if the user chooses 1 then the s.DisplayAverage method will be executed. And a return value of type int will be assigned to the choice variable
	// same with the other methods except for s.Quit ( I did this to have a similar return value like the main function in C++ )
	// s.Quit()  will not return a value it will execute os.Exit(0) which means the program will exit with status 0 (success).
	switch choice {
	case 1:
		choice = s.DisplayAverage()
	case 2:
		choice = s.DisplayMax()
	case 3:
		choice = s.DisplayMin()
	case 4:
		s.Quit()
	}
	// this will return a status code to the main function, and it will keep running this init function if the integer returned is 0
	return choice
}

func main() {
	// While Init's return value is equal to 0 Init Will be executed again. I did this just so it would have a C++ return value style.
	for Init() == 0 {
		Init()
	}
}
