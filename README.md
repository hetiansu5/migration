## Introduce
Database schema builder for golang.

## Quick Start

```golang
package main

import (
	"database/sql"
	"fmt"
	"github.com/hetiansu5/migration/grammar/mysql"
	"github.com/hetiansu5/migration/schema"
)

func main() {
	//db connection need implemented by yourself
	var db *sql.DB

	// register grammar and db driver for default connection
	schema.Register(schema.DefaultConnection, mysql.Grammar{}, mysql.NewDriver(db))
	// create users table
	schema.Create("users", func(table *schema.Blueprint) {
		table.Increments("id")
		table.Integer("age").Default("0").Nullable().Comment("年龄")
		table.Integer("height").Nullable()
		table.String("name").Default("").Charset("utf8mb4").Collate("utf8mb4_unicode_ci").Comment("名字")
		table.Timestamp("created_at").UseCurrent()
		table.Enum("color", []string{"white", "red", "black"})
	})
	// drop users table
	schema.DropIfExists("users")

	// register mongo grammar and db driver for mongo connection
	schema.Register("mongo", mongo.Grammar{}, mongo.driver)
	// generate creating index statements
	statements := schema.Connection("mongo").TableSQL("users", func(table *schema.Blueprint) {
		table.Index("uid,pid", "idx_uid", "BTREE")
		table.Primary("uid")
		table.Unique("uid")
		table.ForeignKey("pid", "products", "id")
	})
	fmt.Println(statements)
}

```

## Example
More to see [example](example/builder.go)

## License
MIT