package mysql

import (
	"fmt"
	"github.com/hetiansu5/migration/schema"
	"strings"
)

type TypeText func(column *schema.Column) string

var typeTextMap = map[uint8]TypeText{
	schema.TinyInteger:           tinyIntegerText,
	schema.SmallInteger:          smallIntegerText,
	schema.MediumInteger:         mediumIntegerText,
	schema.Integer:               integerText,
	schema.BigInteger:            bigIntegerText,
	schema.UnsignedTinyInteger:   unsignedTinyIntegerText,
	schema.UnsignedSmallInteger:  unsignedSmallIntegerText,
	schema.UnsignedMediumInteger: unsignedMediumIntegerText,
	schema.UnsignedInteger:       unsignedIntegerText,
	schema.UnsignedBigInteger:    unsignedBigIntegerText,
	schema.Char:                  charText,
	schema.String:                stringText,
	schema.TinyText:              tinyTextText,
	schema.Text:                  textText,
	schema.MediumText:            mediumTextText,
	schema.LongText:              longTextText,
	schema.Float:                 floatText,
	schema.Double:                doubleText,
	schema.Decimal:               decimalText,
	schema.UnsignedFloat:         unsignedFloatText,
	schema.UnsignedDouble:        unsignedDoubleText,
	schema.UnsignedDecimal:       unsignedDecimalText,
	schema.Boolean:               booleanText,
	schema.Enum:                  enumText,
	schema.Date:                  dateText,
	schema.DateTime:              dateTimeText,
	schema.Time:                  timeText,
	schema.Timestamp:             timestampText,
	schema.Json:                  jsonText,
	schema.Binary:                binaryText,
	schema.Bit:                   bitText,
	schema.TinyBlob:              tinyBlobText,
	schema.Blob:                  blobText,
	schema.MediumBlob:            mediumBlobText,
	schema.LongBlob:              longBlobText,
}

func tinyIntegerText(column *schema.Column) string {
	return getColumnLengthText("tinyint", false, column.GetLength())
}

func smallIntegerText(column *schema.Column) string {
	return getColumnLengthText("smallint", false, column.GetLength())
}

func mediumIntegerText(column *schema.Column) string {
	return getColumnLengthText("mediumint", false, column.GetLength())
}

func integerText(column *schema.Column) string {
	return getColumnLengthText("int", false, column.GetLength())
}

func bigIntegerText(column *schema.Column) string {
	return getColumnLengthText("bigint", false, column.GetLength())
}

func unsignedTinyIntegerText(column *schema.Column) string {
	return getColumnLengthText("tinyint", true, column.GetLength())
}

func unsignedSmallIntegerText(column *schema.Column) string {
	return getColumnLengthText("smallint", true, column.GetLength())
}

func unsignedMediumIntegerText(column *schema.Column) string {
	return getColumnLengthText("mediumint", true, column.GetLength())
}

func unsignedIntegerText(column *schema.Column) string {
	return getColumnLengthText("int", true, column.GetLength())
}

func unsignedBigIntegerText(column *schema.Column) string {
	return getColumnLengthText("bigint", true, column.GetLength())
}

func charText(column *schema.Column) string {
	return fmt.Sprintf("char(%d)", column.GetLength())
}

func stringText(column *schema.Column) string {
	length := column.GetLength()
	if length <= 0 {
		length = 255
	}
	return fmt.Sprintf("varchar(%d)", length)
}

func tinyTextText(column *schema.Column) string {
	return "tinytext"
}

func textText(column *schema.Column) string {
	return "text"
}

func mediumTextText(column *schema.Column) string {
	return "mediumtext"
}

func longTextText(column *schema.Column) string {
	return "longtext"
}

func floatText(column *schema.Column) string {
	return getColumnLengthText("float", false, column.GetLengths()...)
}

func doubleText(column *schema.Column) string {
	return getColumnLengthText("double", false, column.GetLengths()...)
}

func decimalText(column *schema.Column) string {
	return getColumnLengthText("decimal", false, column.GetLengths()...)
}

func unsignedFloatText(column *schema.Column) string {
	return getColumnLengthText("float", true, column.GetLengths()...)
}

func unsignedDoubleText(column *schema.Column) string {
	return getColumnLengthText("double", true, column.GetLengths()...)
}

func unsignedDecimalText(column *schema.Column) string {
	return getColumnLengthText("decimal", true, column.GetLengths()...)
}

func booleanText(column *schema.Column) string {
	return "boolean"
}

func enumText(column *schema.Column) string {
	return fmt.Sprintf("enum('%s')", strings.Join(column.GetEnums(), "', '"))
}

func dateText(column *schema.Column) string {
	return "date"
}

func dateTimeText(column *schema.Column) string {
	return "datetime"
}

func timeText(column *schema.Column) string {
	return "time"
}

func timestampText(column *schema.Column) string {
	return "timestamp"
}

func jsonText(column *schema.Column) string {
	return "json"
}

func binaryText(column *schema.Column) string {
	return getColumnLengthText("binary", false, column.GetLength())
}

func bitText(column *schema.Column) string {
	return getColumnLengthText("bit", false, column.GetLength())
}

func tinyBlobText(column *schema.Column) string {
	return "tinyblob"
}

func blobText(column *schema.Column) string {
	return "blob"
}

func mediumBlobText(column *schema.Column) string {
	return "mediumblob"
}

func longBlobText(column *schema.Column) string {
	return "longblob"
}

func getColumnTypeText(column *schema.Column) string {
	if f, ok := typeTextMap[column.GetDataType()]; ok {
		return f(column)
	}
	panic("undefined column type text")
}

func getColumnLengthText(typeText string, unsigned bool, lengths ...int) string {
	var str string
	if len(lengths) == 1 {
		str = fmt.Sprintf("%s(%d)", typeText, lengths[0])
	} else if len(lengths) >= 2 {
		str = fmt.Sprintf("%s(%d, %d)", typeText, lengths[0], lengths[1])
	} else {
		str = typeText
	}
	if !unsigned {
		return str
	}
	return str + " unsigned"
}
