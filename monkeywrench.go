package monkeywrench

import (
	"context"
	"fmt"

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
//     vals interface{} - The values to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) Insert(table string, cols []string, vals interface{}) error {
	return m.InsertMulti(table, cols, []interface{}{vals})
}

// InsertMulti - Insert multiple rows into a table.
//
// The slice of values supplied must match the names of the columns.
//
// Params:
//     table string - The name of the table to insert into.
//     cols []string - The columns to insert data into.
//     vals []interface{} - A slice of values to import.
//
// Return:
//     error - An error if it occurred.
func (m *MonkeyWrench) InsertMulti(table string, cols []string, vals []interface{}) error {
	// Get the mutation for the insert.
	mutation := spanner.Insert(table, cols, vals)

	// Do the insert.
	err := m.applyMutations([]*spanner.Mutation{mutation})
	if err != nil {
		return err
	}

	return nil
}

// TODO: Implement insert multiple map function.

// TODO: Implement insert single map function.

// TODO: Implement insert multiple struct function.

// TODO: Implement insert single struct function.

// TODO: Implement generic multiple insertOrUpdate function.

// TODO: Implement generic single insertOrUpdate function.

// TODO: Implement insertOrUpdate multiple map function.

// TODO: Implement insertOrUpdate single map function.

// TODO: Implement insertOrUpdate multiple struct function.

// TODO: Implement insertOrUpdate single struct function.

// TODO: Implement generic multiple update function.

// TODO: Implement generic single update function.

// TODO: Implement update multiple map function.

// TODO: Implement update single map function.

// TODO: Implement update multiple struct function.

// TODO: Implement update single struct function.

// TODO: Implement select function (columns, table, conditions, force index - optional).

// TODO: Implement generic query function.

// TODO: Implement generic multiple delete function.

// TODO: Implement generic single delete function.

// TODO: Implement delete multiple map function.

// TODO: Implement delete single map function.

// TODO: Implement delete multiple struct function.

// TODO: Implement delete single struct function.

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
