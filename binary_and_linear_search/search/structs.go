package search

// A structure to hold data that I wanted to use later.
type Values struct {
	TypeOfSearch  string // "Binary" or "Linear"
	IndexLocation int
	TimesSearched int
	ArrayIndices  // Anonymous field (I don't have to specify the type and I can use the values in ArrayIndices as Values.High, Values.Low)
	Err           error
	Array         []int
	Target        int
}

type ArrayIndices struct {
	High int
	Low  int
}
