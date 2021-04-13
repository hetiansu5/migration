package schema

import "sync"

var columnGenerator *ColumnGenerator

type ColumnGenerator struct {
	blueprint *Blueprint
}

func (g *ColumnGenerator) Increments(name string) *Column {
	return g.createColumn(name, UnsignedInteger).AutoIncrement()
}

func (g *ColumnGenerator) TinyInteger(name string) *Column {
	return g.createColumn(name, TinyInteger)
}

func (g *ColumnGenerator) SmallInteger(name string) *Column {
	return g.createColumn(name, SmallInteger)
}

func (g *ColumnGenerator) MediumInteger(name string) *Column {
	return g.createColumn(name, MediumInteger)
}

func (g *ColumnGenerator) Integer(name string) *Column {
	return g.createColumn(name, Integer)
}

func (g *ColumnGenerator) BigInteger(name string) *Column {
	return g.createColumn(name, BigInteger)
}

func (g *ColumnGenerator) UnsignedTinyInteger(name string) *Column {
	return g.createColumn(name, UnsignedTinyInteger)
}

func (g *ColumnGenerator) UnsignedSmallInteger(name string) *Column {
	return g.createColumn(name, UnsignedSmallInteger)
}

func (g *ColumnGenerator) UnsignedMediumInteger(name string) *Column {
	return g.createColumn(name, UnsignedMediumInteger)
}

func (g *ColumnGenerator) UnsignedInteger(name string) *Column {
	return g.createColumn(name, UnsignedInteger)
}

func (g *ColumnGenerator) UnsignedBigInteger(name string) *Column {
	return g.createColumn(name, UnsignedBigInteger)
}

func (g *ColumnGenerator) Char(name string) *Column {
	return g.createColumn(name, Char)
}

func (g *ColumnGenerator) String(name string) *Column {
	return g.createColumn(name, String)
}

func (g *ColumnGenerator) TinyText(name string) *Column {
	return g.createColumn(name, TinyText)
}

func (g *ColumnGenerator) Text(name string) *Column {
	return g.createColumn(name, Text)
}

func (g *ColumnGenerator) MediumText(name string) *Column {
	return g.createColumn(name, MediumText)
}

func (g *ColumnGenerator) LongText(name string) *Column {
	return g.createColumn(name, LongText)
}

func (g *ColumnGenerator) Float(name string, length ...int) *Column {
	return g.createColumn(name, Float, ColumnLength(length...))
}

func (g *ColumnGenerator) Double(name string, length ...int) *Column {
	return g.createColumn(name, Double, ColumnLength(length...))
}

func (g *ColumnGenerator) Decimal(name string, length ...int) *Column {
	return g.createColumn(name, Decimal, ColumnLength(length...))
}

func (g *ColumnGenerator) UnsignedFloat(name string, length ...int) *Column {
	return g.createColumn(name, UnsignedFloat, ColumnLength(length...))
}

func (g *ColumnGenerator) UnsignedDouble(name string, length ...int) *Column {
	return g.createColumn(name, UnsignedDouble, ColumnLength(length...))
}

func (g *ColumnGenerator) UnsignedDecimal(name string, length ...int) *Column {
	return g.createColumn(name, UnsignedDecimal, ColumnLength(length...))
}

func (g *ColumnGenerator) Boolean(name string) *Column {
	return g.createColumn(name, Boolean)
}

func (g *ColumnGenerator) Enum(name string, enums []string) *Column {
	return g.createColumn(name, Enum, ColumnEnums(enums...))
}

func (g *ColumnGenerator) Date(name string) *Column {
	return g.createColumn(name, Date)
}

func (g *ColumnGenerator) DateTime(name string) *Column {
	return g.createColumn(name, DateTime)
}

func (g *ColumnGenerator) Time(name string) *Column {
	return g.createColumn(name, Time)
}

func (g *ColumnGenerator) Timestamp(name string) *Column {
	return g.createColumn(name, Timestamp)
}

func (g *ColumnGenerator) Json(name string) *Column {
	return g.createColumn(name, Json)
}

func (g *ColumnGenerator) Binary(name string) *Column {
	return g.createColumn(name, Binary)
}

func (g *ColumnGenerator) Bit(name string) *Column {
	return g.createColumn(name, Bit)
}

func (g *ColumnGenerator) TinyBlob(name string) *Column {
	return g.createColumn(name, TinyBlob)
}

func (g *ColumnGenerator) Blob(name string) *Column {
	return g.createColumn(name, Blob)
}

func (g *ColumnGenerator) MediumBlob(name string) *Column {
	return g.createColumn(name, MediumBlob)
}

func (g *ColumnGenerator) LongBlob(name string) *Column {
	return g.createColumn(name, LongBlob)
}

func (g *ColumnGenerator) Timestamps() {
	g.Timestamp("created_at").UseCurrent().Nullable()
	g.Timestamp("updated_at").UseCurrent().Nullable()
}

func (g *ColumnGenerator) SoftDeletes() {
	g.Timestamp("deleted_at").Nullable()
}

func (g *ColumnGenerator) createColumn(name string, dataType uint8, ops ...ColumnOption) *Column {
	if g.blueprint != nil {
		return g.blueprint.createColumn(name, dataType, ops...)
	}
	return newColumn(name, dataType, ops...)
}

func reuseColumnGenerator() *ColumnGenerator {
	var once sync.Once
	once.Do(func() {
		columnGenerator = &ColumnGenerator{}
	})
	return columnGenerator
}
