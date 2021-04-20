package schema

const (
	IndexPrimaryKey = uint8(iota) + 1
	IndexUnique
	IndexIndex
	IndexForeignKey
)

type Index struct {
	indexType uint8             //索引类型
	columns   []string          //索引字段
	options   map[string]string //自定义选项参数
}

// newIndex create a pointed Index object
// options: key1 value1 key2 value2
func newIndex(columns []string, indexType uint8, options ...string) *Index {
	index := &Index{columns: columns, indexType: indexType}
	if len(options) > 0 {
		for i := 0; i < len(options); i += 2 {
			if i+1 >= len(options) {
				break
			}
			index.SetOption(options[i], options[i+1])
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
	return i.GetOption("name")
}

func (i *Index) SetName(value string) {
	i.SetOption("name", value)
}

func (i *Index) GetAlgorithm() string {
	return i.GetOption("algorithm")
}

func (i *Index) SetAlgorithm(value string) {
	i.SetOption("algorithm", value)
}

func (i *Index) GetOptions() map[string]string {
	return i.options
}

func (i *Index) GetOption(key string) string {
	if i.options == nil {
		return ""
	}
	return i.options[key]
}

func (i *Index) SetOption(key string, value string) {
	i.checkOptions()
	i.options[key] = value
}

func (i *Index) checkOptions() {
	if i.options == nil {
		i.options = make(map[string]string)
	}
}
