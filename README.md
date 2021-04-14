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

## 自定义Grammar
不同的数据库有不同的语法，需要实现各自语法编译器的接口，Demo可参考[grammar/mysql/grammar.go](grammar/mysql/grammar.go)

```golang
type Grammar interface {
	Compile(*Blueprint, *Action) []string
}

type Compiler interface {
	CreateTable(*Blueprint, *Action) []string
	DropTable(*Blueprint, *Action) []string
	DropTableIfExists(*Blueprint, *Action) []string
	CreateIndex(*Blueprint, *Action) []string
	DropIndex(*Blueprint, *Action) []string
	AddColumn(*Blueprint, *Action) []string
	RenameColumn(*Blueprint, *Action) []string
	ChangeColumn(*Blueprint, *Action) []string
	DropColumn(*Blueprint, *Action) []string
}
```

## 自定义Driver
不同的数据库需要实现各个的驱动接口，比如通过语法编译输出的为`ALTER TABLE users ADD size int`，驱动需要能执行改语句。Demo可参考[grammar/mysql/driver.go](grammar/mysql/driver.go)

```golang
type Driver interface {
	Run(statement string) error
}
```

## License
MIT