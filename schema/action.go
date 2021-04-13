package schema

const (
	ActionCreateTable = uint8(iota) + 1
	ActionDropTable
	ActionDropTableIfExists
	ActionCreateIndex
	ActionDropIndex
	ActionAddColumn
	ActionRenameColumn
	ActionChangeColumn
	ActionDropColumn
)

type Action struct {
	actionType uint8
	column     *Column
	column2    *Column
	index      *Index
}

func newAction(actionType uint8) *Action {
	return &Action{
		actionType: actionType,
	}
}

func (a *Action) GetType() uint8 {
	return a.actionType
}

func (a *Action) GetColumn() *Column {
	return a.column
}

func (a *Action) GetColumn2() *Column {
	return a.column2
}

func (a *Action) GetIndex() *Index {
	return a.index
}
