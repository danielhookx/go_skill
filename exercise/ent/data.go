package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/danielhookx/go_skill/exercise/ent/ent"
	"github.com/danielhookx/go_skill/exercise/ent/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient() *ent.Client {
	driver, err := sql.Open("sqlite3", "file:test.db?mode=rwc&_fk=1")
	if err != nil {
		panic(err)
	}

	var client *ent.Client
	client = ent.NewClient(ent.Driver(driver))

	ctx, _ := context.WithTimeout(context.Background(), 13*time.Second)
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
