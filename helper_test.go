package monkeywrench

import (
	"fmt"
	"log"
	"testing"
)

// This is a singer.
type Singer struct {
	SingerID  int64
	FirstName string
	LastName  string
}

// The list of expected columns.
var expectedCols = []string{
	"SingerID",
	"FirstName",
	"LastName",
}

// TestGetColsFromStruct - Test the GetColsFromStruct function.
func TestGetColsFromStruct(t *testing.T) {
	// Get the columns from the struct.
	cols, err := GetColsFromStruct(&Singer{})
	if err != nil {
		log.Fatal(err)
	}

	// Check each expected column is there.
	for _, expectedCol := range expectedCols {
		if !stringInSlice(cols, expectedCol) {
			t.Error(fmt.Sprintf("Expected %s in slice. Was not found", expectedCol))
		}
	}
}

// stringInSlice - Is a given string in a slice of strings.
//
// Params:
//     slice []string - A slice of strings.
//     item string - The item to search for.
//
// Return:
//     bool - Is the item in the slice?
func stringInSlice(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

// ExampleGetColsFromStruct - Example usage of the GetColsFromStruct usage.
func ExampleGetColsFromStruct() {
	// This is a singer.
	type Singer struct {
		SingerID  int64
		FirstName string
		LastName  string
	}

	// Get the columns from the struct.
	cols, err := GetColsFromStruct(&Singer{})
	if err != nil {
		log.Fatal(err)
	}

	// Check each expected column is there.
	for _, col := range cols {
		fmt.Printf("Found column %s", col)
	}
}
