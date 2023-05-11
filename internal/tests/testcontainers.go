package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var DB *sql.DB
var ctx context.Context
var container testcontainers.Container

func InitContainers(m *testing.M) {

	if DB == nil || isDBNotInitialized() {
		initDB()
		defer DB.Close()
		defer container.Terminate(ctx)
	}

	executeTests(m)
}

func isDBNotInitialized() bool {
	return DB.Ping() != nil
}

func executeTests(m *testing.M) {
	os.Exit(m.Run())
}

func initDB() {
	ctx = context.Background()
	var err error

	// Start the container
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:13-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": "password",
				"POSTGRES_DB":       "testdb",
			},
			WaitingFor: wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}

	container, err = testcontainers.GenericContainer(ctx, req)
	if err != nil {
		log.Fatalf("could not start container: %v", err)
	}

	// Connect to the database
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("could not get mapped port: %v", err)
	}

	connStr := fmt.Sprintf("postgres://postgres:password@localhost:%s/testdb?sslmode=disable", port.Port())
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

}
