package mysql

import (
	"fmt"
	"github.com/hetiansu5/migration/schema"
	"strings"
)

type Grammar struct{}

func (m Grammar) CreateTable(blueprint *schema.Blueprint, action *schema.Action) []string {
	columns := blueprint.GetColumns()
	if len(columns) == 0 {
		return nil
	}
	temporary := ""
	if blueprint.GetTemporary() {
		temporary = "temporary"
	}
	sql := fmt.Sprintf("CREATE %s TABLE `%s` (%s) %s", temporary, blueprint.GetTable(),
		wrapCreateColumns(blueprint.GetColumns()), wrapTableOptions(blueprint))
	return []string{sql}
}

func (m Grammar) DropTable(blueprint *schema.Blueprint, action *schema.Action) []string {
	sql := fmt.Sprintf("DROP TABLE `%s`", blueprint.GetTable())
	return []string{sql}
}

func (m Grammar) DropTableIfExists(blueprint *schema.Blueprint, action *schema.Action) []string {
	sql := fmt.Sprintf("DROP TABLE if exists `%s`", blueprint.GetTable())
	return []string{sql}
}

func (m Grammar) CreateIndex(blueprint *schema.Blueprint, action *schema.Action) []string {
	index := action.GetIndex()
	var sql string
	if index.GetIndexType() == schema.IndexForeignKey {
		sql = fmt.Sprintf("ALTER TABLE `%s` ADD %s (`%s`) REFERENCES `%s`(`%s`)", blueprint.GetTable(), getIndexTypeText(index),
			index.GetColumn(), index.GetForeignTable(), index.GetForeignColumn())
	} else {
		sql = fmt.Sprintf("ALTER TABLE `%s` ADD %s %s(`%s`)", blueprint.GetTable(), getIndexTypeText(index),
			index.GetName(), strings.Join(index.GetColumns(), "`, `"))
	}
	return []string{sql}
}

func (m Grammar) DropIndex(blueprint *schema.Blueprint, action *schema.Action) []string {
	var useAlgo string
	if action.GetIndex().GetAlgorithm() != "" {
		useAlgo = fmt.Sprintf("USING %s", action.GetIndex().GetAlgorithm())
	}
	sql := fmt.Sprintf("ALTER TABLE `%s` DROP INDEX %s %s", blueprint.GetTable(), action.GetIndex().GetName(), useAlgo)
	return []string{sql}
}

func (m Grammar) AddColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	columns := blueprint.GetColumns()
	if len(columns) == 0 {
		return nil
	}
	sql := fmt.Sprintf("ALTER TABLE `%s` %s", blueprint.GetTable(), wrapColumns(blueprint.GetColumns()))
	return []string{sql}
}

func (m Grammar) RenameColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	sql := fmt.Sprintf("ALTER TABLE `%s` RENAME COLUMN `%s` TO `%s`", blueprint.GetTable(),
		action.GetColumn().GetName(), action.GetColumn2().GetName())
	return []string{sql}
}

func (m Grammar) ChangeColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	sql := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` %s", blueprint.GetTable(),
		action.GetColumn().GetName(), wrapColumn(action.GetColumn2()))
	return []string{sql}
}

func (m Grammar) DropColumn(blueprint *schema.Blueprint, action *schema.Action) []string {
	sql := fmt.Sprintf("ALTER TABLE `%s` DROP `%s`", blueprint.GetTable(), action.GetColumn().GetName())
	return []string{sql}
}

func (m Grammar) Compile(blueprint *schema.Blueprint, action *schema.Action) []string {
	compile := schema.GetActionCompile(m, action)
	return compile(blueprint, action)
}

func wrapCreateColumns(columns []*schema.Column) string {
	builder := strings.Builder{}
	for _, column := range columns {
		builder.WriteString(wrapColumn(column))
		builder.WriteString(",")
	}
	s := builder.String()
	return s[:len(s)-1]
}

func wrapTableOptions(blueprint *schema.Blueprint) string {
	var attributes []string
	if blueprint.GetEngine() != "" {
		attributes = append(attributes, fmt.Sprintf("ENGINE = %s", blueprint.GetEngine()))
	}
	if blueprint.GetCharset() != "" {
		attributes = append(attributes, fmt.Sprintf("CHARACTER SET = %s", blueprint.GetCharset()))
	}
	if blueprint.GetCollate() != "" {
		attributes = append(attributes, fmt.Sprintf("COLLATE = %s", blueprint.GetCollate()))
	}
	if blueprint.GetComment() != "" {
		attributes = append(attributes, fmt.Sprintf("COMMENT = \"%s\"", blueprint.GetComment()))
	}
	return strings.Join(attributes, ",")
}

func wrapColumns(columns []*schema.Column) string {
	builder := strings.Builder{}
	for _, column := range columns {
		builder.WriteString(fmt.Sprintf("ADD %s, ", wrapColumn(column)))
	}
	s := builder.String()
	return s[:len(s)-2]
}

func wrapColumn(column *schema.Column) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("`%s` %s", column.GetName(), getColumnTypeText(column)))
	if column.GetCharset() != "" {
		builder.WriteString(fmt.Sprintf(" CHARACTER SET %s", column.GetCharset()))
	}
	if column.GetCollate() != "" {
		builder.WriteString(fmt.Sprintf(" COLLATE %s", column.GetCollate()))
	}
	if column.GetUseCurrent() {
		builder.WriteString(" DEFAULT CURRENT_TIMESTAMP")
	} else {
		if column.GetDefault() != "" {
			builder.WriteString(fmt.Sprintf(" DEFAULT '%s'", column.GetDefault()))
		}
	}
	if column.GetNullable() {
		if !column.GetUseCurrent() && column.GetDefault() == "" {
			builder.WriteString(" DEFAULT NULL")
		} else {
			builder.WriteString(" NULL")
		}
	} else {
		builder.WriteString(" NOT NULL")
	}
	if column.GetAutoIncrement() {
		builder.WriteString(" AUTO_INCREMENT PRIMARY KEY")
	}
	if column.GetComment() != "" {
		builder.WriteString(fmt.Sprintf(" COMMENT '%s'", column.GetComment()))
	}
	if column.GetAfter() != "" {
		builder.WriteString(fmt.Sprintf(" AFTER `%s`", column.GetAfter()))
	}
	return builder.String()
}

func getIndexTypeText(index *schema.Index) string {
	switch index.GetIndexType() {
	case schema.IndexPrimaryKey:
		return "PRIMARY KEY"
	case schema.IndexUnique:
		return "UNIQUE"
	case schema.IndexForeignKey:
		return "FOREIGN KEY"
	case schema.IndexIndex:
		return "INDEX"
	}
	panic("unhandled index type")
}
