package schema

const (
	TinyInteger = uint8(iota) + 1
	SmallInteger
	MediumInteger
	Integer
	BigInteger
	UnsignedTinyInteger
	UnsignedSmallInteger
	UnsignedMediumInteger
	UnsignedInteger
	UnsignedBigInteger
	Char
	String
	TinyText
	Text
	MediumText
	LongText
	Float
	Double
	Decimal
	UnsignedFloat
	UnsignedDouble
	UnsignedDecimal
	Boolean
	Enum
	Date
	DateTime
	Time
	Timestamp
	Json
	Binary
	Bit
	TinyBlob
	Blob
	MediumBlob
	LongBlob
)

type Column struct {
	name     string
	dataType uint8
	options  *columnOptions
}

func newColumn(name string, dataType uint8, ops ...ColumnOption) *Column {
	c := &Column{name: name, dataType: dataType, options: newColumnOptions()}
	for _, op := range ops {
		op(c.options)
	}
	return c
}

func (c *Column) checkOptions() {
	if c.options == nil {
		c.options = newColumnOptions()
	}
}

func (c *Column) Length(length ...int) *Column {
	c.checkOptions()
	c.options.length = length
	return c
}

func (c *Column) Default(value string) *Column {
	c.checkOptions()
	c.options.defaultValue = value
	return c
}

func (c *Column) Nullable() *Column {
	c.checkOptions()
	c.options.nullable = true
	return c
}

func (c *Column) UseCurrent() *Column {
	c.checkOptions()
	c.options.useCurrent = true
	return c
}

func (c *Column) Charset(charset string) *Column {
	c.checkOptions()
	c.options.charset = charset
	return c
}

func (c *Column) Collate(collate string) *Column {
	c.checkOptions()
	c.options.collate = collate
	return c
}

func (c *Column) Comment(comment string) *Column {
	c.checkOptions()
	c.options.comment = comment
	return c
}

func (c *Column) After(after string) *Column {
	c.checkOptions()
	c.options.after = after
	return c
}

func (c *Column) AutoIncrement() *Column {
	c.checkOptions()
	c.options.autoIncrement = true
	return c
}

func (c *Column) Enums(enums ...string) *Column {
	c.options.enums = enums
	return c
}

func (c *Column) GetName() string {
	return c.name
}

func (c *Column) GetDataType() uint8 {
	return c.dataType
}

func (c *Column) GetLength() int {
	c.checkOptions()
	if len(c.options.length) > 0 {
		return c.options.length[0]
	}
	return 0
}

func (c *Column) GetLengths() []int {
	c.checkOptions()
	return c.options.length
}

func (c *Column) GetDefault() string {
	c.checkOptions()
	return c.options.defaultValue
}

func (c *Column) GetNullable() bool {
	c.checkOptions()
	return c.options.nullable
}

func (c *Column) GetUseCurrent() bool {
	c.checkOptions()
	return c.options.useCurrent
}

func (c *Column) GetCharset() string {
	c.checkOptions()
	return c.options.charset
}

func (c *Column) GetCollate() string {
	c.checkOptions()
	return c.options.collate
}

func (c *Column) GetComment() string {
	c.checkOptions()
	return c.options.comment
}

func (c *Column) GetAfter() string {
	c.checkOptions()
	return c.options.after
}

func (c *Column) GetAutoIncrement() bool {
	c.checkOptions()
	return c.options.autoIncrement
}

func (c *Column) GetEnums() []string {
	c.checkOptions()
	return c.options.enums
}
