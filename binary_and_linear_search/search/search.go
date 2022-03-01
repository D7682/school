package search

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/D7682/binary_and_linear_search/colors"
)

func Print(v Values) {
	if v.TypeOfSearch == "linear" {
		colors.Count.Printf("     \n\n%v\n", "---------- Linear Algorithm ----------")
		colors.Count.Printf("          index location: %v,\n          times searched: %v\n", v.IndexLocation, v.TimesSearched)
		fmt.Printf("     %v\n\n", "--- Search Time Complexity:  O(n), Space Complexity: O(n) ---")
		colors.Count.Printf("       %v\n\n", "https://www.geeksforgeeks.org/linear-search/")
	} else if v.TypeOfSearch == "binary" {
		colors.Count.Printf("     \n\n%v\n", "---------- Binary Algorithm ----------")
		colors.Count.Printf("     index: %v,\n     times searched: %v\n", v.IndexLocation, v.TimesSearched)
		fmt.Printf("     %v\n\n", "--- Search Time Complexity: O(log n), Space Complexity: O(1) ---")

		colors.Notification.Printf("     %v%v\n", "Heres the Sorted array: ", v.Array)

		colors.Count.Printf("       %v\n", "https://www.geeksforgeeks.org/binary-search/")
	}
}

// Linear Algorithm.
func (v *Values) Linear() error {
	for i, val := range v.Array {
		if val == v.Target {
			v.IndexLocation = i
			v.TimesSearched = i + 1
			v.Err = nil
			break
		} else {
			v.Err = errors.New("the value was not found! The program will close in 15 seconds")
		}
	}
	return v.Err
}

// Binary Algorithm.
func (v *Values) Binary() error {
	// method to sort the array
	sort.Ints(v.Array)
	v.High = len(v.Array) - 1
	v.Low = 0

	var midpoint float64
	// a for loop that only executes based on the number of elements of the array.
	for range v.Array {
		v.TimesSearched++
		midpoint = math.Round((float64(v.High) + float64(v.Low)) / 2.0)
		if v.Array[int(midpoint)] == v.Target {
			v.IndexLocation = int(midpoint)
			break
		} else if v.Array[int(midpoint)] > v.Target {
			v.High = int(midpoint) - 1
		} else if v.Array[int(midpoint)] < v.Target {
			v.Low = int(midpoint) + 1
		} else if v.Low > v.High {
			v.Err = errors.New("the value was not found! program will close in 15 seconds")
		}
	}
	return v.Err
}

// This is the function that routes execution to a binary or linear search.
func (v *Values) Algorithm() {
	if v.TypeOfSearch == "" {
		v.TypeOfSearch = "linear"
	}
	switch v.TypeOfSearch {
	case "binary":
		err := v.Binary()
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(15 * time.Second)
			return
		}
		Print(*v)
	case "linear":
		err := v.Linear()
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(15 * time.Second)
			return
		}
		Print(*v)
	default:
		fmt.Println("v.TypeOfSearch must be \"Binary\" or \"Linear\"")
	}
}
