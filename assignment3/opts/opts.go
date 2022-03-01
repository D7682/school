// Package opts holds all the functions for the menu options.
package opts

import (
	"fmt"
	"os"
)

/*
// The Bubble sort algorithm for this structure. (if needed)

var swap bool
var temp Person

for swap {
	swap = false
	for i := 0; i < len(s.List); i++ {
		if int(s.List[i].Name[0]) > int(s.List[i+1].Name[0]) {
			temp = s.List[i]
			s.List[i] = s.List[i+1]
			s.List[i+1] = temp
			swap = true
		}
	}
}
*/

func (s *Students) DisplayAverage() int {
	// This first function iterates through the list of Students and we are using it to assign the average grade based on the list of each students grade.
	for i, val := range s.List {
		// This for loop is looping through all the values in the students grades to add them up to the average grade.
		for _, v := range val.Grades {
			// we use the s.List instead of val.AverageGrade to give values to a field, because val is just a copy of the structure we are iterating through.
			s.List[i].AverageGrade += v
		}
		s.List[i].AverageGrade = (s.List[i].AverageGrade / float64(len(s.List[i].Grades)))
	}

	// This checks each Student in the list's Average Grade.
	for i, val := range s.List {
		switch {
		case val.AverageGrade >= 90.0:
			s.List[i].AverageLetter = "A"
		case val.AverageGrade >= 80.0 && val.AverageGrade < 90.0:
			s.List[i].AverageLetter = "B"
		case val.AverageGrade >= 70.0 && val.AverageGrade < 80.0:
			s.List[i].AverageLetter = "C"
		case val.AverageGrade >= 60.0 && val.AverageGrade < 70.0:
			s.List[i].AverageLetter = "D"
		default:
			s.List[i].AverageLetter = "F"
		}
	}

	// Printing out the "Grade Averages" Table.
	fmt.Println("Grade Averages")
	fmt.Printf("\n%20s %20s %20s\n", "Name", "Average", "Grade")

	for _, val := range s.List {
		// %20.1f verb tells go to display 1 decimal place on the float64(double) type with a padding of 20.
		fmt.Printf("\n%20s %20.1f %20s\n", val.Name, val.AverageGrade, val.AverageLetter)
	}

	fmt.Printf("\n")

	return 0
}

func (s *Students) DisplayMax() int {
	// This loop goes into the Student struct which runs another for loop to iterate into each value of the Grades array.
	for i := 0; i < len(s.List); i++ {
		for x := 0; x < len(s.List[i].Grades)-1; x++ {
			if s.List[i].Max < s.List[i].Grades[x] {
				s.List[i].Max = s.List[i].Grades[x]
			}
		}
	}

	// This checks each Student in the list's Average Grade.
	for i, val := range s.List {
		switch {
		case val.Max >= 90.0:
			s.List[i].AverageLetter = "A"
		case val.Max >= 80.0 && val.Max < 90.0:
			s.List[i].AverageLetter = "B"
		case val.Max >= 70.0 && val.Max < 80.0:
			s.List[i].AverageLetter = "C"
		case val.Max >= 60.0 && val.Max < 70.0:
			s.List[i].AverageLetter = "D"
		default:
			s.List[i].AverageLetter = "F"
		}
	}

	// Printing "Max Grades" Table.
	fmt.Println("Max Grades")
	fmt.Printf("\n%20s %20s %20s\n", "Name", "Max", "Grade")

	for _, val := range s.List {
		fmt.Printf("\n%20s %20.0f %20s\n", val.Name, val.Max, val.AverageLetter)
	}

	fmt.Printf("\n")
	return 0
}

func (s *Students) DisplayMin() int {

	// setting Min to a value of 200 in order to compare to the grades.
	for i := 0; i < len(s.List); i++ {
		s.List[i].Min = 200
	}

	// This loop goes into the Student struct which runs another for loop to iterate into each value of the Grades array.
	for i := 0; i < len(s.List); i++ {
		for x := 0; x < len(s.List[i].Grades)-1; x++ {
			if s.List[i].Min > s.List[i].Grades[x] {
				s.List[i].Min = s.List[i].Grades[x]
			}
		}
	}

	// This checks each Student in the list's Average Grade.
	for i, val := range s.List {
		switch {
		case val.Min >= 90.0:
			s.List[i].AverageLetter = "A"
		case val.Min >= 80.0 && val.Min < 90.0:
			s.List[i].AverageLetter = "B"
		case val.Min >= 70.0 && val.Min < 80.0:
			s.List[i].AverageLetter = "C"
		case val.Min >= 60.0 && val.Min < 70.0:
			s.List[i].AverageLetter = "D"
		default:
			s.List[i].AverageLetter = "F"
		}
	}

	// Printing "Min Grades" Table.
	fmt.Println("Min Grades")
	fmt.Printf("\n%20s %20s %20s\n", "Name", "Min", "Grade")

	for _, val := range s.List {
		fmt.Printf("\n%20s %20.0f %20s\n", val.Name, val.Min, val.AverageLetter)
	}

	fmt.Printf("\n")
	return 0
}

func (s *Students) Quit() {
	os.Exit(0)
}
