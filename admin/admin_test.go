package admin

import (
	"context"
	"fmt"
	"os"
)

// ExampleSpannerAdmin_CreateAdminClient - Example usage for CreateAdminClient.
func ExampleSpannerAdmin_CreateAdminClient() {
	ctx := context.Background()

	// Create the admin struct.
	spannerAdmin := &SpannerAdmin{
		Context:  ctx,
		Project:  "my-awesome-project-id",
		Instance: "my-awesome-spanner-instance",
	}

	// Create the admin client.
	err := spannerAdmin.CreateAdminClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner admin client. Reason - %+v", err)
		os.Exit(1)
	}
}

// ExampleSpannerAdmin_CreateDatabase - Example usage for CreateDatabase.
func ExampleSpannerAdmin_CreateDatabase() {
	ctx := context.Background()

	// Create the admin struct.
	spannerAdmin := &SpannerAdmin{
		Context:  ctx,
		Project:  "my-awesome-project-id",
		Instance: "my-awesome-spanner-instance",
	}

	// Create the admin client.
	err := spannerAdmin.CreateAdminClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner admin client. Reason - %+v", err)
	}

	// Some tables to create with the database.
	tablesDDL := []string{
		`CREATE TABLE Singers (
			SingerId   INT64 NOT NULL,
			FirstName  STRING(1024),
			LastName   STRING(1024),
			SingerInfo BYTES(MAX)
		) PRIMARY KEY (SingerId)`,
		`CREATE TABLE Albums (
			SingerId	INT64 NOT NULL,
			AlbumId		INT64 NOT NULL,
			AlbumTitle	STRING(MAX),
		) PRIMARY KEY (SingerId, AlbumId),
		INTERLEAVE IN PARENT Singers ON DELETE CASCADE`,
	}

	// Create the database.
	dbErr := spannerAdmin.CreateDatabase("my-awesome-spanner-database", tablesDDL)
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner database. Reason - %+v", dbErr)
	}
}

// ExampleSpannerAdmin_AlterDatabase - Example usage for AlterDatabase.
func ExampleSpannerAdmin_AlterDatabase() {
	ctx := context.Background()

	// Create the admin struct.
	spannerAdmin := &SpannerAdmin{
		Context:  ctx,
		Project:  "my-awesome-project-id",
		Instance: "my-awesome-spanner-instance",
	}

	// Create the admin client.
	err := spannerAdmin.CreateAdminClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Spanner admin client. Reason - %+v", err)
	}

	// Some indexes we want to add.
	indexDDL := []string{
		`CREATE INDEX AlbumsByAlbumTitle ON Albums(AlbumTitle)`,
		`CREATE INDEX AlbumsByAlbumTitle2 ON Albums(AlbumTitle) STORING (MarketingBudget)`,
	}

	// Alter the database.
	dbErr := spannerAdmin.AlterDatabase("my-awesome-spanner-database", indexDDL)
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to alter Spanner database. Reason - %+v", dbErr)
	}
}
