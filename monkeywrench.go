package monkeywrench

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	"cloud.google.com/go/spanner"
)

const (
	// The pattern to build the fully qualified Cloud Spanner parent name.
	fqParentPattern = "projects/%s/instances/%s"

	// The pattern to build the fully qualified Cloud Spanner database name.
	fqDbPattern = "projects/%s/instances/%s/databases/%s"
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
	fqDb := fmt.Sprintf(fqDbPattern, m.Project, m.Instance, m.Db)

	// Create the client.
	spannerClient, err := spanner.NewClient(m.Context, fqDb, m.Opts...)
	if err != nil {
		return err
	}

	// Set the client.
	m.Client = spannerClient
	return nil
}

// TODO: Implement generic multiple insert function.

// TODO: Implement generic single insert function.

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
