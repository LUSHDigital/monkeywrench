package monkeywrench

import (
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

// TODO: Add tests and example usage for GetColsFromStruct.

// GetColsFromStruct - Get the names of all fields in a struct.
//
// Params:
//     src interface{} - The struct to get field names from.
//
// Return:
//     []string - Slice of field names from the struct.
//     error - An error if it occurred.
func GetColsFromStruct(src interface{}) ([]string, error) {
	// Get the reflected structure.
	reflectStruct := reflect.ValueOf(src).Type()

	// If we have a pointer, get the addressed element type instead.
	if reflectStruct.Kind() == reflect.Ptr {
		reflectStruct = reflectStruct.Elem()
	}

	// Check we have a structure and not something else.
	if dataType := reflectStruct.Kind(); dataType != reflect.Struct {
		return nil, fmt.Errorf("Unsupported data type %s", dataType.String())
	}

	// Get the columns.
	var cols []string
	for i := 0; i < reflectStruct.NumField(); i++ {
		// Skip this field if it is marked as ignored.
		tag := string(reflectStruct.Field(i).Tag)
		if strings.Contains(tag, "spanner:\"-\"") {
			continue
		}

		cols = append(cols, reflectStruct.Field(i).Name)
	}

	return cols, nil
}

// getResultSlice - Get the results of a row iterator as a slice.
//
// Params:
//     iter *spanner.RowIterator - The iterator from a query or read.
//
// Return:
//     []*spanner.Row - A list of all rows returned by the query.
//     error - An error if it occurred.
func getResultSlice(iter *spanner.RowIterator) ([]*spanner.Row, error) {
	var results []*spanner.Row

	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return results, nil
		}
		if err != nil {
			return nil, err
		}

		results = append(results, row)
	}
}
