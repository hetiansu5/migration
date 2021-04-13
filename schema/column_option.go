package schema

type columnOptions struct {
	length        []int
	defaultValue  string
	nullable      bool
	useCurrent    bool
	charset       string
	collate       string
	comment       string
	after         string
	autoIncrement bool
	enums         []string
}

type ColumnOption func(options *columnOptions)

func newColumnOptions() *columnOptions {
	return &columnOptions{}
}

func ColumnLength(length ...int) ColumnOption {
	return func(options *columnOptions) {
		options.length = length
	}
}

func ColumnDefault(value string) ColumnOption {
	return func(options *columnOptions) {
		options.defaultValue = value
	}
}

func ColumnNullable() ColumnOption {
	return func(options *columnOptions) {
		options.nullable = true
	}
}

func ColumnUseCurrent() ColumnOption {
	return func(options *columnOptions) {
		options.useCurrent = true
	}
}

func ColumnCharset(charset string) ColumnOption {
	return func(options *columnOptions) {
		options.charset = charset
	}
}

func ColumnCollate(collate string) ColumnOption {
	return func(options *columnOptions) {
		options.collate = collate
	}
}

func ColumnComment(comment string) ColumnOption {
	return func(options *columnOptions) {
		options.comment = comment
	}
}

func ColumnAfter(after string) ColumnOption {
	return func(options *columnOptions) {
		options.after = after
	}
}

func ColumnAutoIncrement() ColumnOption {
	return func(options *columnOptions) {
		options.autoIncrement = true
	}
}

func ColumnEnums(enums ...string) ColumnOption {
	return func(options *columnOptions) {
		options.enums = enums
	}
}
