package admin

import (
	"context"
	"fmt"

	"github.com/LUSHDigital/monkeywrench"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

// SpannerAdmin - Wrapper for Cloud Spanner admin.
type SpannerAdmin struct {
	Context     context.Context
	Project     string
	Instance    string
	Opts        []option.ClientOption
	AdminClient *database.DatabaseAdminClient
}

// CreateAdminClient - Create a new Cloud Spanner admin client.
//
// Return:
//     error - An error if it occurred.
func (a *SpannerAdmin) CreateAdminClient() error {
	adminClient, err := database.NewDatabaseAdminClient(a.Context, a.Opts...)
	if err != nil {
		return err
	}

	// Set the client.
	a.AdminClient = adminClient
	return nil
}

// CreateDatabase - Create a new Cloud Spanner database.
//
// Params:
//     db string - Name of the Cloud Spanner database to create.
//     ddl []string - Data Definition Language statements to apply to the newly
//     created database.
//
// Return:
//     error - An error if it occurred.
func (a *SpannerAdmin) CreateDatabase(db string, ddl []string) error {
	fmt.Println("Creating Cloud Spanner database.")

	op, err := a.AdminClient.CreateDatabase(a.Context, &adminpb.CreateDatabaseRequest{
		Parent:          fmt.Sprintf(monkeywrench.FqParentPattern, a.Project, a.Instance),
		CreateStatement: "CREATE DATABASE `" + db + "`",
		ExtraStatements: ddl,
	})

	// Handle any errors.
	if grpc.Code(err) == codes.AlreadyExists {
		fmt.Printf("Database (%s) already exists.\n", db)
		return nil
	} else if err != nil {
		return err
	}

	// Wait for the database to be created.
	if _, err := op.Wait(a.Context); err == nil {
		fmt.Printf("Created database (%s).\n", db)
	}

	return nil
}

// AlterDatabase - Alter a Cloud Spanner database.
//
// Can be used with any valid DDL to add/alter/drop tables or add/alter/drop
// indexes.
//
// Params:
//     db string - Name of the Cloud Spanner database to create.
//     ddl []string - Data Definition Language statements to alter a database.
//
// Return:
//     error - An error if it occurred.
func (a *SpannerAdmin) AlterDatabase(db string, ddl []string) error {
	fmt.Println("Altering Cloud Spanner database.")
	op, err := a.AdminClient.UpdateDatabaseDdl(a.Context, &adminpb.UpdateDatabaseDdlRequest{
		Database:   fmt.Sprintf(monkeywrench.FqDbPattern, a.Project, a.Instance, db),
		Statements: ddl,
	})
	if err != nil {
		return err
	}

	// Wait for the database to be altered.
	if err := op.Wait(a.Context); err == nil {
		fmt.Printf("Altered database (%s).\n", db)
	}

	return nil
}
