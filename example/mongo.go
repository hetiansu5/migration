package main

import (
	"context"
	"fmt"
	mongoGrammar "github.com/hetiansu5/migration/grammar/mongo"
	"github.com/hetiansu5/migration/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect mongodb: %s", err.Error()))
		return
	}
	driver := mongoGrammar.NewDriver(client.Database("migration"))
	schema.Register(schema.DefaultConnection, mongoGrammar.Grammar{}, driver)
	dropMongoTable()
	createMongoIndex()
	dropMongoIndex()
}

func dropMongoTable() {
	schema.Drop("users")
}

func createMongoIndex() {
	schema.Table("users", func(table *schema.Blueprint) {
		table.Index("uid:-1,pid", "name", "idx_uid_pid")
		table.Unique("session_id")
	})
}

func dropMongoIndex() {
	schema.Table("users", func(table *schema.Blueprint) {
		table.DropIndex("idx_uid_pid")
	})
}
