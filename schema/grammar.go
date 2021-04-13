package schema

type ActionHandler func(*Blueprint, *Action) []string

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

func GetActionCompile(compiler Compiler, action *Action) ActionHandler {
	switch action.GetType() {
	case ActionCreateTable:
		return compiler.CreateTable
	case ActionDropTable:
		return compiler.DropTable
	case ActionDropTableIfExists:
		return compiler.DropTableIfExists
	case ActionCreateIndex:
		return compiler.CreateIndex
	case ActionDropIndex:
		return compiler.DropIndex
	case ActionAddColumn:
		return compiler.AddColumn
	case ActionRenameColumn:
		return compiler.RenameColumn
	case ActionChangeColumn:
		return compiler.ChangeColumn
	case ActionDropColumn:
		return compiler.DropColumn
	}
	return nil
}
