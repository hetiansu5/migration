package mongo

import (
	"fmt"
	"github.com/hetiansu5/migration/schema"
	"strings"
)

type Grammar struct{}

func (m Grammar) CreateTable(blueprint *schema.Blueprint, action *schema.Action) []string {
	return nil
}

func (m Grammar) DropTable(blueprint *schema.Blueprint, action *schema.Action) []string {
	return []string{fmt.Sprintf("op=dropTable&table=%s", blueprint.GetTable())}
}

func (m Grammar) DropTableIfExists(blueprint *schema.Blueprint, action *schema.Action) []string {
	return []string{fmt.Sprintf("op=dropTable&table=%s&options=exists", blueprint.GetTable())}
}

func (m Grammar) CreateIndex(blueprint *schema.Blueprint, action *schema.Action) []string {
	optionString := formOptionString(action.GetIndex().GetOptions())
	if action.GetIndex().GetIndexType() == schema.IndexUnique {
		if optionString == "" {
			optionString = "unique"
		} else {
			optionString += ",unique"
		}
	}
	return []string{fmt.Sprintf("op=createIndex&table=%s&keys=%s&options=%s", blueprint.GetTable(),
		strings.Join(action.GetIndex().GetColumns(), ","), optionString)}
}

func (m Grammar) DropIndex(blueprint *schema.Blueprint, action *schema.Action) []string {
	optionString := formOptionString(action.GetIndex().GetOptions())
	return []string{fmt.Sprintf("op=dropIndex&table=%s&options=%s",
		blueprint.GetTable(), optionString)}
}

func (m Grammar) AddColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	return nil
}

func (m Grammar) RenameColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	return nil
}

func (m Grammar) ChangeColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	return nil
}

func (m Grammar) DropColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	return nil
}

func (m Grammar) Compile(blueprint *schema.Blueprint, action *schema.Action) []string {
	compile := schema.GetActionCompile(m, action)
	return compile(blueprint, action)
}

func formOptionString(options map[string]string) string {
	items := make([]string, 0, len(options))
	for k, v := range options {
		items = append(items, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(items, ",")
}
