// A generated module for Class1 functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/class-1/internal/dagger"
)

type Class1 struct{}

func (m *Class1) DB(db *dagger.Directory) *dagger.Service {
	initSql := db.File("database.sql")

	return dag.Container().
		From("postgres:12.21").
		WithFile("/docker-entrypoint-initdb.d/database.sql", initSql).
		WithEnvVariable("POSTGRES_DB", "db").
		WithEnvVariable("POSTGRES_USER", "dbuser").
		WithEnvVariable("POSTGRES_PASSWORD", "dbpass").
		WithExposedPort(5432).
		AsService()
}

// Returns a container that echoes whatever string argument is provided
func (m *Class1) AppTest(
	ctx context.Context,
	app *dagger.Directory,
) *dagger.Service {
	initSql := app.File("./testdata/database.sql")

	db := dag.Container().
		From("postgres:12.21").
		WithFile("/docker-entrypoint-initdb.d/database.sql", initSql).
		WithEnvVariable("POSTGRES_DB", "gocourse").
		WithEnvVariable("POSTGRES_USER", "dbuser").
		WithEnvVariable("POSTGRES_PASSWORD", "dbpass").
		WithExposedPort(5432).
		AsService()

	return dag.
		Container().
		From("golang:1.23.5").
		WithServiceBinding("db", db).
		WithMountedDirectory("/app", app).
		WithWorkdir("/app").
		WithExec([]string{"go", "mod", "download"}).
		WithExec([]string{"go", "run", "main.go"}).
		WithExposedPort(8080).
		AsService()
}

// Returns a container that echoes whatever string argument is provided
func (m *Class1) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Class1) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
