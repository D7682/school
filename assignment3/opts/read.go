package opts

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/D7682/assignment3/stringParser"
)

// creating new types here.
type Students struct {
	List []Person
}

type Person struct {
	Name          string
	Grades        []float64
	AverageGrade  float64
	AverageLetter string
	Max           float64
	Min           float64
}

// This is the method of the Students type.
// The syntax for a method is:
// func (s *ReceiverType) NameOfMethod() {}
func (s *Students) ReadFile() {
	// os.OpenFile returns a pointer to the file and an error value.
	// The octal value 0444 gives everyone permission to read the file.
	file, err := os.OpenFile("NamesGrades.txt", os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewScanner(file)

	// the Scan() method scans every token in a file until it reaches the end of a file or an error (it returns true/false).
	for r.Scan() {
		// I split each line
		// returns a string for each line.
		line := r.Text()
		// I split each line into words
		// this returns an array of strings separated by a space.
		words := strings.Split(line, " ")

		// the "Students" struct has a values of List which is an array of the "Person" struct
		// So I added a "Person" struct in each value of the array, and then I added the values.
		s.List = append(s.List, Person{
			Name:   stringParser.FormatName(words[:2]), // words[:2] is telling the compiler to get words[0], words[1] ("Sidra", "Amartey"); stringParser.FormatName is a package I created to concatenate the first and last name.
			Grades: stringParser.Parse(words[2:]),      // words[2:] is telling it to get the values from words[3], words[4], words[5], words[6], words[7] (the grades); string.Parser.Parse is parsing from type string to type float64.
		})
	}

	// f.Close() to close the file.
	// I don't need the file anymore, because I already added all the values to the original struct via a pointer.
	// defer means that this anonymous function will be executed after everything else is done executing and this "ReadFile" method is going to close.
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
