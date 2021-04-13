package schema

const (
	IndexPrimaryKey = uint8(iota) + 1
	IndexUnique
	IndexIndex
	IndexForeignKey
)

type Index struct {
	indexType uint8    //索引类型
	columns   []string //索引字段
	name      string   //索引名称
	algorithm string   //算法
}

// newIndex create a pointed Index object
// options[0]: indexName
// options[1]: algorithm
func newIndex(columns []string, indexType uint8, options ...string) *Index {
	index := &Index{columns: columns, indexType: indexType}
	if len(options) >= 1 {
		index.name = options[0]
		if len(options) >= 2 {
			index.algorithm = options[1]
			if len(options) >= 3 {
				panic("too long options")
			}
		}
	}
	return index
}

func (i *Index) GetIndexType() uint8 {
	return i.indexType
}

func (i *Index) GetColumn() string {
	return i.columns[0]
}

func (i *Index) GetColumns() []string {
	return i.columns
}

func (i *Index) GetName() string {
	return i.name
}

func (i *Index) GetAlgorithm() string {
	return i.algorithm
}

func (i *Index) GetForeignTable() string {
	return i.name
}

func (i *Index) GetForeignColumn() string {
	return i.algorithm
}
