package schema

import (
	"fmt"
	"sync"
)

type Builder struct {
	driver  Driver
	grammar Grammar
}

const (
	DefaultConnection = "default"
)

var builderMap map[string]*Builder
var mutex sync.RWMutex

func newBuilder(grammar Grammar, driver Driver) *Builder {
	return &Builder{driver: driver, grammar: grammar}
}

func (b *Builder) Create(table string, callback func(*Blueprint)) {
	statements := b.CreateSQL(table, callback)
	b.exec(statements)
}

func (b *Builder) CreateSQL(table string, callback func(*Blueprint)) []string {
	blueprint := NewBlueprint(table)
	blueprint.Create()
	callback(blueprint)
	return b.toSQL(blueprint)
}

func (b *Builder) Table(table string, callback func(*Blueprint)) {
	statements := b.TableSQL(table, callback)
	b.exec(statements)
}

func (b *Builder) TableSQL(table string, callback func(*Blueprint)) []string {
	blueprint := NewBlueprint(table)
	blueprint.AddColumns()
	callback(blueprint)
	return b.toSQL(blueprint)
}

func (b *Builder) Drop(table string) {
	statements := b.DropSQL(table)
	b.exec(statements)
}

func (b *Builder) DropSQL(table string) []string {
	blueprint := NewBlueprint(table)
	blueprint.Drop()
	return b.toSQL(blueprint)
}

func (b *Builder) DropIfExists(table string) {
	statements := b.DropIfExistsSQL(table)
	b.exec(statements)
}

func (b *Builder) DropIfExistsSQL(table string) []string {
	blueprint := NewBlueprint(table)
	blueprint.DropIfExist()
	return b.toSQL(blueprint)
}

func (b *Builder) build(blueprint *Blueprint) {
	statements := b.toSQL(blueprint)
	b.exec(statements)
}

func (b *Builder) exec(statements []string) {
	for _, statement := range statements {
		err := b.driver.Run(statement)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (b *Builder) toSQL(blueprint *Blueprint) []string {
	return blueprint.ToSQL(b.grammar)
}

func Create(table string, callback func(*Blueprint)) {
	GetBuilder().Create(table, callback)
}

func CreateSQL(table string, callback func(*Blueprint)) []string {
	return GetBuilder().CreateSQL(table, callback)
}

func Table(table string, callback func(*Blueprint)) {
	GetBuilder().Table(table, callback)
}

func TableSQL(table string, callback func(*Blueprint)) []string {
	return GetBuilder().TableSQL(table, callback)
}

func Drop(table string) {
	GetBuilder().Drop(table)
}

func DropSQL(table string) []string {
	return GetBuilder().DropSQL(table)
}

func DropIfExists(table string) {
	GetBuilder().DropIfExists(table)
}

func DropIfExistsSQL(table string) []string {
	return GetBuilder().DropIfExistsSQL(table)
}

func Connection(connections ...string) *Builder {
	return GetBuilder(connections...)
}

func Register(connection string, grammar Grammar,driver Driver) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := builderMap[connection]; ok {
		panic(fmt.Sprintf("connection(%s) had been registered", connection))
	}
	builderMap[connection] = newBuilder(grammar, driver)
}

func GetBuilder(connections ...string) *Builder {
	var connection string
	if len(connections) == 0 {
		connection = DefaultConnection
	} else {
		connection = connections[0]
	}
	mutex.RLock()
	defer mutex.RUnlock()
	return builderMap[connection]
}

func init() {
	builderMap = make(map[string]*Builder)
}
