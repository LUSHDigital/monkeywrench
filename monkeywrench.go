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
func (m *MonkeyWrench) CreateClient(sessionPoolConfig spanner.SessionPoolConfig) error {
	// Build the fully qualified db name.
	fqDb := fmt.Sprintf(FqDbPattern, m.Project, m.Instance, m.Db)

	// Create the client.
	spannerClient, err := spanner.NewClientWithConfig(m.Context, fqDb, spanner.ClientConfig{
		SessionPoolConfig: sessionPoolConfig,
	}, m.Opts...)
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

// Delete - Delete a row from a table by key.
//
// Params:
//     table string - The table to delete from.
//     key spanner.Key - The key to delete.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) Delete(table string, key spanner.Key) error {
	return m.DeleteMulti(table, []spanner.Key{key})
}

// DeleteMulti - Delete multiple rows from a table by key.
//
// Params:
//     table string - The table to delete from.
//     keys []spanner.Key - The list of keys to delete.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) DeleteMulti(table string, keys []spanner.Key) error {
	// Create a mutation for each value set we have.
	mutations := make([]*spanner.Mutation, 0, len(keys))
	for _, key := range keys {
		mutations = append(mutations, spanner.Delete(table, key))
	}

	// Apply the mutations.
	err := m.applyMutations(mutations)
	if err != nil {
		return err
	}

	return nil
}

// DeleteKeyRange - Delete a range of rows by key.
//
// Params:
//     table string - The table to delete rows from.
//     startKey interface{} - The starting value of the range.
//     endKey interface{} - The ending value of the range.
//     rangeKind spanner.KeyRangeKind - The kind of range (includes keys or not)
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) DeleteKeyRange(table string, startKey, endKey spanner.Key, rangeKind spanner.KeyRangeKind) error {
	// Create the mutation.
	mutation := spanner.Delete(table, spanner.KeyRange{
		Start: startKey,
		End:   endKey,
		Kind:  rangeKind,
	})

	// Apply the mutations.
	err := m.applyMutations([]*spanner.Mutation{mutation})
	if err != nil {
		return err
	}

	return nil
}

// Query - Executes a query against Cloud Spanner.
func (m *MonkeyWrench) Query(statement string, params ...map[string]interface{}) ([]*spanner.Row, error) {
	return m.QueryCtx(m.Context, statement, params...)
}

// QueryCtx is the same as Query but allows passing your own cancellable context
func (m *MonkeyWrench) QueryCtx(ctx context.Context, statement string, params ...map[string]interface{}) ([]*spanner.Row, error) {
	// Prepare the raw statement.
	stmt := spanner.NewStatement(statement)

	// Add any parameters we've been given.
	for _, param := range params {
		for key, value := range param {
			stmt.Params[key] = value
		}
	}

	// Execute the query.
	iter := m.Client.Single().Query(ctx, stmt)
	return getResultSlice(iter)
}

// Read - Read multiple rows from Cloud Spanner.
//
// Params:
//     table string - Name of the table to read rows from.
//     index string - Name of the index to use from the table.
//     keys []spanner.KeySet - Slice of keys for the rows to read. Passing an empty
//     slice will cause all rows to be returned.
//     columns []string - List of columns to read for each row.
//
// Return:
//     []*spanner.Row - A list of all rows returned by the query.
//     error - An error if it occurred.
func (m *MonkeyWrench) Read(table string, keys []spanner.KeySet, columns []string) ([]*spanner.Row, error) {
	// Default to all keys.
	var spannerKeys = spanner.AllKeys()

	// If we have some specified keys, use those instead.
	if len(keys) > 0 {
		spannerKeys = spanner.KeySets(keys...)
	}

	// Execute the query.
	iter := m.Client.Single().Read(m.Context, table, spannerKeys, columns)
	return getResultSlice(iter)
}

// ReadUsingIndex - Read multiple rows from Cloud Spanner using an index.
//
// Params:
//     table string - Name of the table to read rows from.
//     index string - Name of the index to use from the table.
//     keys []spanner.KeySet - List of keys for the rows to read. Passing an empty
//     slice will cause all rows to be returned.
//     columns []string - List of columns to read for each row.
//
// Return:
//     []*spanner.Row - A list of all rows returned by the query.
//     error - An error if it occurred.
func (m *MonkeyWrench) ReadUsingIndex(table, index string, keys []spanner.KeySet, columns []string) ([]*spanner.Row, error) {
	// Default to all keys.
	var spannerKeys = spanner.AllKeys()

	// If we have some specified keys, use those instead.
	if len(keys) > 0 {
		spannerKeys = spanner.KeySets(keys...)
	}

	// Execute the query.
	iter := m.Client.Single().ReadUsingIndex(m.Context, table, index, spannerKeys, columns)
	return getResultSlice(iter)
}

// ReadToStruct - Read a row from Spanner table to a struct.
//
// Params:
//     table string - Name of the table to read from.
//     key spanner.Key - The key for the row to read.
//     dst interface - Destination struct.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) ReadToStruct(table string, key spanner.Key, dst interface{}) error {
	// Get the value of the destination parameter.
	dstValue := reflect.Indirect(reflect.ValueOf(dst))

	// Check we were passed a valid data type.
	dataType := dstValue.Type().Kind()
	if dataType != reflect.Struct {
		return fmt.Errorf("Unsupported data type: %s", dataType.String())
	}

	// The columns to read.
	cols, err := GetColsFromStruct(dst)
	if err != nil {
		return fmt.Errorf("Could not get columns from struct. Reason: %s", err)
	}

	// Perform the read.
	rows, err := m.Read(table, []spanner.KeySet{key}, cols)
	if err != nil {
		return err
	}

	// Decode the row onto the struct.
	for _, row := range rows {
		row.ToStruct(dst)
	}

	return nil
}

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

	// Check the type of data we were passed is valid.
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
