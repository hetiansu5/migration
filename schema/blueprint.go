package schema

import (
	"strings"
)

type Blueprint struct {
	*ColumnGenerator
	table     string
	engine    string
	charset   string
	collate   string
	comment   string
	temporary bool
	columns   []*Column
	actions   []*Action
}

func NewBlueprint(table string) *Blueprint {
	b := &Blueprint{table: table}
	b.ColumnGenerator = &ColumnGenerator{blueprint: b}
	return b
}

func (b *Blueprint) ToSQL(grammar Grammar) []string {
	var statements []string
	for _, action := range b.actions {
		statements = append(statements, grammar.Compile(b, action)...)
	}
	return statements
}

func (b *Blueprint) Temporary() *Blueprint {
	b.temporary = true
	return b
}

func (b *Blueprint) Engine(v string) *Blueprint {
	b.engine = v
	return b
}

func (b *Blueprint) Charset(v string) *Blueprint {
	b.charset = v
	return b
}

func (b *Blueprint) Collate(v string) *Blueprint {
	b.collate = v
	return b
}

func (b *Blueprint) Comment(v string) *Blueprint {
	b.comment = v
	return b
}

func (b *Blueprint) Create() {
	action := newAction(ActionCreateTable)
	b.addAction(action)
}

func (b *Blueprint) Drop() {
	action := newAction(ActionDropTable)
	b.addAction(action)
}

func (b *Blueprint) DropIfExist() {
	action := newAction(ActionDropTableIfExists)
	b.addAction(action)
}

// Primary create a primary key
// columns: multiple columns joined with comma
// options[0]: indexName
// options[1]: algorithm
func (b *Blueprint) Primary(columns string, options ...string) {
	b.createIndex(columns, IndexPrimaryKey, options...)
}

// Unique create a unique index
// columns: multiple columns joined with comma
// options[0]: indexName
// options[1]: algorithm
func (b *Blueprint) Unique(columns string, options ...string) {
	b.createIndex(columns, IndexUnique, options...)
}

// Index create a common index
// columns: multiple columns joined with comma
// options[0]: indexName
// options[1]: algorithm
func (b *Blueprint) Index(columns string, options ...string) {
	b.createIndex(columns, IndexIndex, options...)
}

// ForeignKey create a foreign key
// columns: multiple columns joined with comma
// options[0]: foreign table
// options[1]: foreign column
func (b *Blueprint) ForeignKey(columns string, options ...string) {
	b.createIndex(columns, IndexForeignKey, options...)
}

func (b *Blueprint) DropIndex(indexName string) {
	action := newAction(ActionDropIndex)
	action.index = newIndex([]string{}, 0, indexName)
	b.addAction(action)
}

func (b *Blueprint) AddColumns() {
	action := newAction(ActionAddColumn)
	b.addAction(action)
}

func (b *Blueprint) RenameColumn(originName string, targetName string) {
	action := newAction(ActionRenameColumn)
	action.column = newColumn(originName, 0)
	action.column2 = newColumn(targetName, 0)
	b.addAction(action)
}

func (b *Blueprint) ChangeColumn(originName string, callback func(generator *ColumnGenerator) *Column) {
	action := newAction(ActionChangeColumn)
	action.column = newColumn(originName, 0)
	action.column2 = callback(reuseColumnGenerator())
	b.addAction(action)
}

func (b *Blueprint) DropColumn(name string) {
	action := newAction(ActionDropColumn)
	action.column = newColumn(name, 0)
	b.addAction(action)
}

func (b *Blueprint) addAction(action *Action) {
	b.actions = append(b.actions, action)
}

func (b *Blueprint) createIndex(columns string, indexType uint8, options ...string) {
	action := newAction(ActionCreateIndex)
	action.index = newIndex(strings.Split(columns, ","), indexType, options...)
	b.addAction(action)
}

func (b *Blueprint) createColumn(name string, dataType uint8, ops ...ColumnOption) *Column {
	column := newColumn(name, dataType, ops...)
	b.columns = append(b.columns, column)
	return column
}

func (b *Blueprint) GetTable() string {
	return b.table
}

func (b *Blueprint) GetEngine() string {
	return b.engine
}

func (b *Blueprint) GetCharset() string {
	return b.charset
}

func (b *Blueprint) GetCollate() string {
	return b.collate
}

func (b *Blueprint) GetComment() string {
	return b.comment
}

func (b *Blueprint) GetTemporary() bool {
	return b.temporary
}

func (b *Blueprint) GetColumns() []*Column {
	return b.columns
}

func (b *Blueprint) getAction() []*Action {
	return b.actions
}
