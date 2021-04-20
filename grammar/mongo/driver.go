package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"strconv"
	"strings"
)

type Driver struct {
	db *mongo.Database
}

type Statement struct {
	op      string
	table   string
	keys    []string
	options map[string]string
}

func (s *Statement) GetOp() string {
	return s.op
}

func (s *Statement) GetTable() string {
	return s.table
}

func (s *Statement) GetKeys() []string {
	return s.keys
}

func (s *Statement) HasOption(key string) bool {
	_, ok := s.options[key]
	return ok
}

func (s *Statement) GetOption(key string) string {
	return s.options[key]
}

func NewDriver(db *mongo.Database) *Driver {
	return &Driver{
		db: db,
	}
}

func (d *Driver) Run(statement string) error {
	var err error
	sm := parseStatement(statement)
	switch sm.op {
	case "dropTable":
		err = d.db.Collection(sm.GetTable()).Drop(context.TODO())
	case "createIndex":
		docx := make(bsonx.Doc, 0, len(sm.GetKeys()))
		var i int
		for _, key := range sm.GetKeys() {
			arr := strings.SplitN(key, ":", 2)
			if len(arr) >= 2 {
				i, _ = strconv.Atoi(arr[1])
			} else {
				i = 1
			}
			docx = append(docx, bsonx.Elem{Key: arr[0], Value: bsonx.Int32(int32(i))})
		}
		indexModel := mongo.IndexModel{
			Keys: docx,
			Options: &options.IndexOptions{},
		}
		if sm.HasOption("background") {
			indexModel.Options.SetBackground(true)
		}
		if sm.HasOption("name") {
			indexModel.Options.SetName(sm.GetOption("name"))
		}
		if sm.HasOption("unique") {
			indexModel.Options.SetUnique(true)
		}
		_, err = d.db.Collection(sm.GetTable()).Indexes().CreateOne(context.TODO(), indexModel)
	case "dropIndex":
		_, err = d.db.Collection(sm.GetTable()).Indexes().DropOne(context.TODO(), sm.GetOption("name"))
	}
	return err
}

func parseStatement(statement string) *Statement {
	container := splitToMap(statement, "&", "=")
	sm := &Statement{
		op:    container["op"],
		table: container["table"],
	}
	if v, ok := container["keys"]; ok {
		sm.keys = strings.Split(v, ",")
	}
	if v, ok := container["options"]; ok {
		sm.options = parseOptions(v)
	}
	return sm
}

func parseOptions(options string) map[string]string {
	return splitToMap(options, ",", ":")
}

func splitToMap(s string, sep string, subSep string) map[string]string {
	items := strings.Split(s, sep)
	container := make(map[string]string, len(items))
	for _, item := range items {
		arr := strings.SplitN(item, subSep, 2)
		if len(arr) >= 2 {
			container[arr[0]] = arr[1]
		} else {
			container[arr[0]] = ""
		}
	}
	return container
}
