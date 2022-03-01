package stringParser

import (
	"log"
	"strconv"
)

// Created this function to parse the numbers on the file from a type string to a type int.
func Parse(s []string) []float64 {
	var converted []float64
	for _, val := range s {
		convertedValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
		}
		converted = append(converted, convertedValue)
	}

	return converted
}

func FormatName(s []string) string {
	return s[0] + " " + s[1]
}
