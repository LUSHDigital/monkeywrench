package monkeywrench

import (
	"context"
	"fmt"
	"os"
)

// ExampleMonkeyWrench_Insert - Example of the Insert function.
func ExampleMonkeyWrench_Insert() {
	ctx := context.Background()

	// Create Cloud Spanner wrapper.
	mW := &MonkeyWrench{
		Context:  ctx,
		Project:  "my-awesome-project",
		Instance: "my-awesome-spanner-instance",
		Db:       "my-awesome-spanner-database",
	}

	// Create a Spanner client.
	if spannerErr := mW.CreateClient(); spannerErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner client. Reason - %+v\n", spannerErr)
		os.Exit(1)
	}

	singerCols := []string{
		"SingerId",
		"FirstName",
		"LastName",
	}

	// Insert a single row.
	if insertErr := mW.Insert("Singers", singerCols, []interface{}{1, "Joe", "Bloggs"}); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertMulti - Example of the InsertMulti function.
func ExampleMonkeyWrench_InsertMulti() {
	ctx := context.Background()

	// Create Cloud Spanner wrapper.
	mW := &MonkeyWrench{
		Context:  ctx,
		Project:  "my-awesome-project",
		Instance: "my-awesome-spanner-instance",
		Db:       "my-awesome-spanner-database",
	}

	// Create a Spanner client.
	if spannerErr := mW.CreateClient(); spannerErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner client. Reason - %+v\n", spannerErr)
		os.Exit(1)
	}

	singerCols := []string{
		"SingerId",
		"FirstName",
		"LastName",
	}

	singers := [][]interface{}{
		[]interface{}{2, "John", "Smith"},
		[]interface{}{3, "Anne", "Other"},
	}

	// Insert multiple rows.
	if insertErr := mW.InsertMulti("Singers", singerCols, singers); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdate - Example of the InsertOrUpdate function.
func ExampleMonkeyWrench_InsertOrUpdate() {
	ctx := context.Background()

	// Create Cloud Spanner wrapper.
	mW := &MonkeyWrench{
		Context:  ctx,
		Project:  "my-awesome-project",
		Instance: "my-awesome-spanner-instance",
		Db:       "my-awesome-spanner-database",
	}

	// Create a Spanner client.
	if spannerErr := mW.CreateClient(); spannerErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner client. Reason - %+v\n", spannerErr)
		os.Exit(1)
	}

	singerCols := []string{
		"SingerId",
		"FirstName",
		"LastName",
	}

	// Insert a single row.
	if insertErr := mW.InsertOrUpdate("Singers", singerCols, []interface{}{1, "J", "Bloggs"}); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateMulti - Example of the InsertOrUpdateMulti function.
func ExampleMonkeyWrench_InsertOrUpdateMulti() {
	ctx := context.Background()

	// Create Cloud Spanner wrapper.
	mW := &MonkeyWrench{
		Context:  ctx,
		Project:  "my-awesome-project",
		Instance: "my-awesome-spanner-instance",
		Db:       "my-awesome-spanner-database",
	}

	// Create a Spanner client.
	if spannerErr := mW.CreateClient(); spannerErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner client. Reason - %+v\n", spannerErr)
		os.Exit(1)
	}

	singerCols := []string{
		"SingerId",
		"FirstName",
		"LastName",
	}

	singers := [][]interface{}{
		[]interface{}{2, "J", "Smith"},
		[]interface{}{3, "A", "Other"},
	}

	// Insert multiple rows.
	if insertErr := mW.InsertOrUpdateMulti("Singers", singerCols, singers); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}
