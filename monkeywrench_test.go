package monkeywrench

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
)

// ExampleMonkeyWrench_Insert - Example usage for the Insert function.
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
	if err := mW.Insert("Singers", singerCols, []interface{}{1, "Joe", "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertMulti - Example usage for the InsertMulti function.
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
	if err := mW.InsertMulti("Singers", singerCols, singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdate - Example usage for the InsertOrUpdate function.
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

	// Insert/update a single row.
	if err := mW.InsertOrUpdate("Singers", singerCols, []interface{}{1, "J", "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateMulti - Example usage for the InsertOrUpdateMulti function.
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

	// Insert/update multiple rows.
	if err := mW.InsertOrUpdateMulti("Singers", singerCols, singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_Update - Example usage for the Update function.
func ExampleMonkeyWrench_Update() {
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

	// Update a single row.
	if err := mW.Update("Singers", singerCols, []interface{}{1, "J", "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_UpdateMulti - Example usage for the UpdateMulti function.
func ExampleMonkeyWrench_UpdateMulti() {
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

	// Update multiple rows.
	if err := mW.UpdateMulti("Singers", singerCols, singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertMap - Example usage for the InsertMap function.
func ExampleMonkeyWrench_InsertMap() {
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

	// Insert a single row.
	if err := mW.InsertMap("Singers", map[string]interface{}{"SingerId": 1, "FirstName": "Joe", "LastName": "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertMapMulti - Example usage for the InsertMapMulti function.
func ExampleMonkeyWrench_InsertMapMulti() {
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

	singers := []map[string]interface{}{
		map[string]interface{}{"SingerId": 2, "FirstName": "John", "LastName": "Smith"},
		map[string]interface{}{"SingerId": 3, "FirstName": "Anne", "LastName": "Other"},
	}

	// Insert multiple rows.
	if err := mW.InsertMapMulti("Singers", singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateMap - Example usage for the InsertOrUpdateMap function.
func ExampleMonkeyWrench_InsertOrUpdateMap() {
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

	// Insert/update a single row.
	if err := mW.InsertOrUpdateMap("Singers", map[string]interface{}{"SingerId": 1, "FirstName": "J", "LastName": "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateMapMulti - Example usage for the InsertOrUpdateMapMulti function.
func ExampleMonkeyWrench_InsertOrUpdateMapMulti() {
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

	singers := []map[string]interface{}{
		map[string]interface{}{"SingerId": 2, "FirstName": "J", "LastName": "Smith"},
		map[string]interface{}{"SingerId": 3, "FirstName": "A", "LastName": "Other"},
	}

	// Insert/update multiple rows.
	if err := mW.InsertOrUpdateMapMulti("Singers", singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_UpdateMap - Example usage for the UpdateMap function.
func ExampleMonkeyWrench_UpdateMap() {
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

	// Update a single row.
	if insertErr := mW.UpdateMap("Singers", map[string]interface{}{"SingerId": 1, "FirstName": "J", "LastName": "Bloggs"}); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_UpdateMapMulti - Example usage for the UpdateMapMulti function.
func ExampleMonkeyWrench_UpdateMapMulti() {
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

	singers := []map[string]interface{}{
		map[string]interface{}{"SingerId": 2, "FirstName": "J", "LastName": "Smith"},
		map[string]interface{}{"SingerId": 3, "FirstName": "A", "LastName": "Other"},
	}

	// Update multiple rows.
	if insertErr := mW.UpdateMapMulti("Singers", singers); insertErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", insertErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertStruct - Example usage for the InsertStruct function.
func ExampleMonkeyWrench_InsertStruct() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	// Insert a single row.
	if err := mW.InsertStruct("Singers", Singer{SingerID: 1, FirstName: "Joe", LastName: "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertStructMulti - Example usage for the InsertStructMulti function.
func ExampleMonkeyWrench_InsertStructMulti() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	singers := []Singer{
		Singer{SingerID: 2, FirstName: "John", LastName: "Smith"},
		Singer{SingerID: 3, FirstName: "Anne", LastName: "Other"},
	}

	// Insert multiple rows.
	if err := mW.InsertStructMulti("Singers", singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateStruct - Example usage for the InsertOrUpdateStruct function.
func ExampleMonkeyWrench_InsertOrUpdateStruct() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	// Insert/update a single row.
	if err := mW.InsertOrUpdateStruct("Singers", Singer{SingerID: 1, FirstName: "Joe", LastName: "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_InsertOrUpdateStructMulti - Example usage for the InsertOrUpdateStructMulti function.
func ExampleMonkeyWrench_InsertOrUpdateStructMulti() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	singers := []Singer{
		Singer{SingerID: 2, FirstName: "John", LastName: "Smith"},
		Singer{SingerID: 3, FirstName: "Anne", LastName: "Other"},
	}

	// Insert/update multiple rows.
	if err := mW.InsertOrUpdateStructMulti("Singers", singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_UpdateStruct - Example usage for the UpdateStruct function.
func ExampleMonkeyWrench_UpdateStruct() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	// Update a single row.
	if err := mW.UpdateStruct("Singers", Singer{SingerID: 1, FirstName: "Joe", LastName: "Bloggs"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_UpdateStructMulti - Example usage for the UpdateStructMulti function.
func ExampleMonkeyWrench_UpdateStructMulti() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	singers := []Singer{
		Singer{SingerID: 2, FirstName: "John", LastName: "Smith"},
		Singer{SingerID: 3, FirstName: "Anne", LastName: "Other"},
	}

	// Update multiple rows.
	if err := mW.UpdateStructMulti("Singers", singers); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert into Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_Delete - Example usage for the Delete function.
func ExampleMonkeyWrench_Delete() {
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

	// Delete a row.
	if deleteErr := mW.Delete("Singers", spanner.Key{2}); deleteErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete from Spanner. Reason - %+v\n", deleteErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_DeleteMulti - Example usage for the DeleteMulti function.
func ExampleMonkeyWrench_DeleteMulti() {
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

	// Delete multiple rows.
	if deleteErr := mW.DeleteMulti("Singers", []spanner.Key{spanner.Key{1}, spanner.Key{4}}); deleteErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete from Spanner. Reason - %+v\n", deleteErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_DeleteKeyRange - Example usage for the DeleteKeyRange function.
func ExampleMonkeyWrench_DeleteKeyRange() {
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

	// Delete multiple rows.
	if deleteErr := mW.DeleteKeyRange("Singers", spanner.Key{2}, spanner.Key{4}, spanner.ClosedClosed); deleteErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete from Spanner. Reason - %+v\n", deleteErr)
		os.Exit(1)
	}
}

// ExampleMonkeyWrench_Query - Example usage for the Query function.
func ExampleMonkeyWrench_Query() {
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

	// Run the query.
	results, err := mW.Query(`SELECT FirstName, LastName FROM Singers`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results.
	for _, result := range results {
		var FirstName, LastName string

		// Get each result.
		result.ColumnByName("FirstName", &FirstName)
		result.ColumnByName("LastName", &LastName)

		// What did we get?
		fmt.Printf("Found singer: %s %s\n", FirstName, LastName)
	}
}

// ExampleMonkeyWrench_Read - Example usage for the Read function.
func ExampleMonkeyWrench_Read() {
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

	// Run the query for select keys.
	results, err := mW.Read("Singers", []spanner.Key{spanner.Key{1}, spanner.Key{4}}, []string{"FirstName", "LastName"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results.
	for _, result := range results {
		var FirstName, LastName string

		// Get each result.
		result.ColumnByName("FirstName", &FirstName)
		result.ColumnByName("LastName", &LastName)

		// What did we get?
		fmt.Printf("Found singer: %s %s\n", FirstName, LastName)
	}

	// Get all the things!
	allResults, err := mW.Read("Singers", []spanner.Key{}, []string{"FirstName", "LastName"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results.
	for _, aResult := range allResults {
		var FirstName, LastName string

		// Get each result.
		aResult.ColumnByName("FirstName", &FirstName)
		aResult.ColumnByName("LastName", &LastName)

		// What did we get?
		fmt.Printf("Found singer: %s %s\n", FirstName, LastName)
	}
}

// ExampleMonkeyWrench_ReadUsingIndex - Example usage for the ReadUsingIndex function.
func ExampleMonkeyWrench_ReadUsingIndex() {
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

	// Simple index
	// Index DDL - `CREATE INDEX SingersByLastName ON Singers(LastName)`
	results, err := mW.ReadUsingIndex("Singers", "SingersByLastName", []spanner.Key{}, []string{"LastName"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner using index. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results.
	for _, result := range results {
		var LastName string

		// Get each result.
		result.ColumnByName("LastName", &LastName)

		// What did we get?
		fmt.Printf("Found last name: %s\n", LastName)
	}

	// Storing index
	// Index DDL - `CREATE INDEX SingersByLastNameWithFirstName ON Singers(LastName) STORING (FirstName)`
	storingResults, err := mW.ReadUsingIndex("Singers", "SingersByLastNameWithFirstName", []spanner.Key{}, []string{"FirstName", "LastName"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner using index. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results.
	for _, sResult := range storingResults {
		var FirstName, LastName string

		// Get each result.
		sResult.ColumnByName("FirstName", &FirstName)
		sResult.ColumnByName("LastName", &LastName)

		// What did we get?
		fmt.Printf("Found singer: %s %s\n", FirstName, LastName)
	}
}

func ExampleMonkeyWrench_Read_multistruct() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	var aSinger Singer

	// Get the columns from the struct.
	structCols, err := GetColsFromStruct(aSinger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get columns from struct. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Get all the things!
	rows, err := mW.Read("Singers", []spanner.Key{}, structCols)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner. Reason - %+v\n", err)
		os.Exit(1)
	}

	// Print the results
	for _, row := range rows {
		var singer Singer
		row.ToStruct(&singer)

		fmt.Printf("Found a singer (%d) called %s %s\n", singer.SingerID, singer.FirstName, singer.LastName)
	}
}

// ExampleMonkeyWrench_ReadToStruct - Example usage for the ReadToStruct function.
func ExampleMonkeyWrench_ReadToStruct() {
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

	// This is a singer.
	type Singer struct {
		SingerID  int64 `spanner:"SingerId"`
		FirstName string
		LastName  string
	}

	var aSinger Singer

	// Read a single result to a struct.
	structErr := mW.ReadToStruct("Singers", spanner.Key{1}, &aSinger)
	if structErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to read Spanner. Reason - %+v\n", structErr)
		os.Exit(1)
	}

	// What did we get?
	fmt.Printf("Found one singer (%d) called %s %s\n", aSinger.SingerID, aSinger.FirstName, aSinger.LastName)
}
