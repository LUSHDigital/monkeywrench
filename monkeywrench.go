package monkeywrench

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/api/option"

	"cloud.google.com/go/spanner"
)

const (
	// FqParentPattern - The pattern to build the fully qualified Cloud Spanner
	// parent name.
	FqParentPattern = "projects/%s/instances/%s"

	// FqDbPattern - The pattern to build the fully qualified Cloud Spanner
	// database name.
	FqDbPattern = "projects/%s/instances/%s/databases/%s"
)

// MonkeyWrench - Wrapper for Cloud Spanner.
type MonkeyWrench struct {
	Context  context.Context
	Project  string
	Instance string
	Db       string
	Opts     []option.ClientOption
	Client   *spanner.Client
}

// CreateClient - Create a new Spanner client.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) CreateClient() error {
	// Build the fully qualified db name.
	fqDb := fmt.Sprintf(FqDbPattern, m.Project, m.Instance, m.Db)

	// Create the client.
	spannerClient, err := spanner.NewClient(m.Context, fqDb, m.Opts...)
	if err != nil {
		return err
	}

	// Set the client.
	m.Client = spannerClient
	return nil
}

// Insert - Insert a row into a table.
//
// The supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     vals interface{} - The data to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) Insert(table string, cols []string, vals []interface{}) error {
	return m.applyGenericMutations(table, cols, [][]interface{}{vals}, spanner.Insert)
}

// InsertMulti - Insert multiple rows into a table.
//
// The slice of values supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData [][]interface{} - A slice of data to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertMulti(table string, cols []string, sourceData [][]interface{}) error {
	return m.applyGenericMutations(table, cols, sourceData, spanner.Insert)
}

// InsertOrUpdate - Insert or update a row into a table.
//
// The slice of values supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData [][]interface{} - The values to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdate(table string, cols []string, vals []interface{}) error {
	return m.applyGenericMutations(table, cols, [][]interface{}{vals}, spanner.InsertOrUpdate)
}

// InsertOrUpdateMulti - Insert or update multiple rows into a table.
//
// The slice of values supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData [][]interface{} - The values to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdateMulti(table string, cols []string, sourceData [][]interface{}) error {
	return m.applyGenericMutations(table, cols, sourceData, spanner.InsertOrUpdate)
}

// Update - Update a row in a table.
//
// The supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to update.
//     cols []string - The columns to update.
//     vals interface{} - The data to update.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) Update(table string, cols []string, vals []interface{}) error {
	return m.applyGenericMutations(table, cols, [][]interface{}{vals}, spanner.Update)
}

// UpdateMulti - Update multiple rows in a table.
//
// The slice of values supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to update
//     cols []string - The columns to update.
//     sourceData [][]interface{} - A slice of data to update.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) UpdateMulti(table string, cols []string, sourceData [][]interface{}) error {
	return m.applyGenericMutations(table, cols, sourceData, spanner.Update)
}

// InsertMap - Insert a row, based on a map, into a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData map[string]interface{} - The map of col => value data to
//     insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertMap(table string, sourceData map[string]interface{}) error {
	return m.applyMapMutations(table, []map[string]interface{}{sourceData}, spanner.InsertMap)
}

// InsertMapMulti - Insert multiple rows, based on maps, into a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData []map[string]interface{} - Nested map of col => value data to
//     insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertMapMulti(table string, sourceData []map[string]interface{}) error {
	return m.applyMapMutations(table, sourceData, spanner.InsertMap)
}

// InsertOrUpdateMap - Insert or update a row, based on a map, into
// a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData map[string]interface{} - The map of col => value data to
//     insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdateMap(table string, sourceData map[string]interface{}) error {
	return m.applyMapMutations(table, []map[string]interface{}{sourceData}, spanner.InsertOrUpdateMap)
}

// InsertOrUpdateMapMulti - Insert or update multiple rows, based on maps, into
// a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData []map[string]interface{} - Nested map of col => value data to
//     insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdateMapMulti(table string, sourceData []map[string]interface{}) error {
	return m.applyMapMutations(table, sourceData, spanner.InsertOrUpdateMap)
}

// UpdateMap - Update a row, based on a map, in a table.
//
// Params:
//     table string - The name of the table to update.
//     sourceData map[string]interface{} - The map of col => value data to
//     update in the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) UpdateMap(table string, sourceData map[string]interface{}) error {
	return m.applyMapMutations(table, []map[string]interface{}{sourceData}, spanner.UpdateMap)
}

// UpdateMapMulti - Update multiple rows, based on maps, in a table.
//
// Params:
//     table string - The name of the table to update.
//     sourceData []map[string]interface{} - Nested map of col => value data to
//     update in the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) UpdateMapMulti(table string, sourceData []map[string]interface{}) error {
	return m.applyMapMutations(table, sourceData, spanner.UpdateMap)
}

// InsertStruct - Insert a row, based on a struct, into a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData interface - The data struct to insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertStruct(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, []interface{}{sourceData}, spanner.InsertStruct)
}

// InsertStructMulti - Insert multiple rows, based on a struct, into a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData interface - The data struct to insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertStructMulti(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, sourceData, spanner.InsertStruct)
}

// InsertOrUpdateStruct - Insert or update a row, based on a struct, into a
// table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData interface - The data struct to insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdateStruct(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, []interface{}{sourceData}, spanner.InsertOrUpdateStruct)
}

// InsertOrUpdateStructMulti - Insert or update multiple rows, based on a
// struct, into a table.
//
// Params:
//     table string - The name of the table to insert into.
//     sourceData interface - The data struct to insert into the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertOrUpdateStructMulti(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, sourceData, spanner.InsertOrUpdateStruct)
}

// UpdateStruct - Update a row, based on a struct, into a table.
//
// Params:
//     table string - The name of the table to update.
//     sourceData interface - The data struct to update in the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) UpdateStruct(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, []interface{}{sourceData}, spanner.UpdateStruct)
}

// UpdateStructMulti - Update multiple rows, based on a struct, in a table.
//
// Params:
//     table string - The name of the table to update.
//     sourceData interface - The data struct to update in the table.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) UpdateStructMulti(table string, sourceData interface{}) error {
	return m.applyStructMutations(table, sourceData, spanner.UpdateStruct)
}

// TODO: Implement generic single delete function.

// TODO: Implement generic multiple delete function.

// TODO: Implement delete single map function.

// TODO: Implement delete multiple map function.

// TODO: Implement delete single struct function.

// TODO: Implement delete multiple struct function.

// TODO: Implement select function (columns, table, conditions, force index - optional).

// TODO: Implement generic query function.

// applyGenericMutations - Apply a set of generic mutations.
//
// This function is intended to generate and apply mutations for generic data
// based on key => value.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData [][]interface{} - The data to import.
//     generator func(table string, cols []string, vals []interface{}) *spanner.Mutation - The callback to generate mutations.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) applyGenericMutations(table string, cols []string, sourceData [][]interface{}, generator func(table string, cols []string, vals []interface{}) *spanner.Mutation) error {
	// Get the values from the passed source data.
	vals := reflect.ValueOf(sourceData)

	// Check the type of data we were passed is vald.
	dataKind := vals.Type().Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return fmt.Errorf("Unsupported type: %s", dataKind.String())
	}

	// Create a mutation for each value set we have.
	mutations := make([]*spanner.Mutation, 0, vals.Len())
	for _, value := range sourceData {
		mutations = append(mutations, generator(table, cols, value))
	}

	// Apply the mutations.
	err := m.applyMutations(mutations)
	if err != nil {
		return err
	}

	return nil
}

// applyMapMutations - Apply a set of map mutations.
//
// This function is intended to generate and apply mutations based on maps.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData interface{} - The data to import.
//     generator func(table string, data map[string]interface{}) *spanner.Mutation - The callback to generate mutations.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) applyMapMutations(table string, sourceData []map[string]interface{}, generator func(table string, data map[string]interface{}) *spanner.Mutation) error {
	// Get the values from the passed source data.
	vals := reflect.ValueOf(sourceData)

	// Check the type of data we were passed is vald.
	dataKind := vals.Type().Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return fmt.Errorf("Unsupported type: %s", dataKind.String())
	}

	// Create a mutation for each value set we have.
	mutations := make([]*spanner.Mutation, 0, vals.Len())
	for _, value := range sourceData {
		mutations = append(mutations, generator(table, value))
	}

	// Apply the mutations.
	err := m.applyMutations(mutations)
	if err != nil {
		return err
	}

	return nil
}

// applyStructMutations - Apply a set of structured mutations.
//
// This function is intended to generate and apply mutations based on structs.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     sourceData interface{} - The data to import.
//     generator func(table string, data interface{}) (*spanner.Mutation, error) - The callback to generate mutations.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) applyStructMutations(table string, sourceData interface{}, generator func(table string, data interface{}) (*spanner.Mutation, error)) error {
	// Get the values from the passed source data.
	vals := reflect.Indirect(reflect.ValueOf(sourceData))

	// Check the type of data we were passed is vald.
	dataKind := vals.Type().Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return fmt.Errorf("Unsupported type: %s", dataKind.String())
	}

	// Create a mutation for each value set we have.
	mutations := make([]*spanner.Mutation, 0, vals.Len())
	for i := 0; i < vals.Len(); i++ {
		value := vals.Index(i)
		mutation, err := generator(table, value.Interface())
		if err != nil {
			return err
		}
		mutations = append(mutations, mutation)
	}

	// Apply the mutations.
	err := m.applyMutations(mutations)
	if err != nil {
		return err
	}

	return nil
}

// applyMutations - Apply a set of mutations to Cloud Spanner
//
// Params:
//     mutations []*spanner.Mutation - The mutations to apply.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) applyMutations(mutations []*spanner.Mutation) error {
	_, err := m.Client.Apply(m.Context, mutations)
	if err != nil {
		return err
	}

	return nil
}
