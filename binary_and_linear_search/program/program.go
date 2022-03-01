package program

import (
	"time"

	"github.com/D7682/binary_and_linear_search/ascii"
	"github.com/D7682/binary_and_linear_search/colors"
	"github.com/D7682/binary_and_linear_search/input"
	"github.com/D7682/binary_and_linear_search/search"
)

func Init() {
	ascii.Art()
	array, target, searchType := input.Read()

	data := search.Values{TypeOfSearch: searchType, Target: target, Array: array}
	data.Algorithm()

	colors.Notification.Printf("\nPress \"ctrl + c\" to exit.")
	time.Sleep(2 * time.Minute) // if everything works correctly the app will close itself in 2 minutes. In case someone wants to have enough time to read the data.
}
