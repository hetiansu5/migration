package main

import (
	"fmt"
	"github.com/hetiansu5/migration/grammar/mysql"
	"github.com/hetiansu5/migration/schema"
)

func main() {
	schema.Register(schema.DefaultConnection, mysql.Grammar{}, nil)
	createTable()
	dropTable()
	dropTableIfExists()
	createIndex()
	addColumn()
	renameColumn()
	changeColumn()
	dropColumn()
}

func createTable() {
	statements := schema.CreateSQL("users", func(table *schema.Blueprint) {
		table.Increments("id")
		table.Integer("age").Default("0").Nullable().Comment("年龄")
		table.Integer("height").Nullable()
		table.String("name").Default("").Charset("utf8mb4").Collate("utf8mb4_unicode_ci").Comment("年龄")
		table.Timestamp("created_at").UseCurrent()
		table.Enum("color", []string{"white", "red", "black"})
	})
	printStatements(statements)
}

func dropTable() {
	statements := schema.DropSQL("users")
	printStatements(statements)
}

func dropTableIfExists() {
	statements := schema.DropIfExistsSQL("users")
	printStatements(statements)
}

func createIndex() {
	statements := schema.TableSQL("users", func(table *schema.Blueprint) {
		table.Index("uid,pid", "idx_uid", "BTREE")
		table.Primary("uid")
		table.Unique("uid")
		table.ForeignKey("pid", "products", "id")
	})
	printStatements(statements)
}

func addColumn() {
	statements := schema.TableSQL("users", func(table *schema.Blueprint) {
		table.Integer("age").Default("0").Nullable().Comment("年龄")
		table.String("name").Default("").Charset("utf8mb4").Collate("utf8mb4_unicode_ci").Comment("年龄")
		table.Integer("height").Nullable()
		table.Timestamp("created_at").UseCurrent().After("name")
	})
	printStatements(statements)
}

func renameColumn() {
	statements := schema.TableSQL("users", func(table *schema.Blueprint) {
		table.RenameColumn("size", "busy")
	})
	printStatements(statements)
}

func changeColumn() {
	statements := schema.TableSQL("users", func(table *schema.Blueprint) {
		table.ChangeColumn("age", func(generator *schema.ColumnGenerator) *schema.Column {
			return generator.String("size").Default("1").Comment("尺寸")
		})
		table.ChangeColumn("height", func(generator *schema.ColumnGenerator) *schema.Column {
			return generator.BigInteger("high").Default("0").After("name")
		})
	})
	printStatements(statements)
}

func dropColumn() {
	statements := schema.TableSQL("users", func(table *schema.Blueprint) {
		table.DropColumn("size")
	})
	printStatements(statements)
}

func printStatements(statements []string) {
	for _, statement := range statements {
		fmt.Println(statement)
	}
}
