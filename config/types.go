package config

var MysqlToGoTypes = map[string]string{
	"bit":        "bool",
	"tinyint":    "byte",
	"smallint":   "int8",
	"mediumint":  "int16",
	"int":        "int",
	"bigint":     "int64",
	"float":      "float32",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "time.Time",
	"time":       "string",
	"year":       "int8",
	"datetime":   "string",
	"timestamp":  "int64",
	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"tinyblob":   "byte[]",
	"mediumblob": "byte[]",
	"blob":       "byte[]",
	"longblob":   "byte[]",
}

var GoSqlNullTypes = map[string]string{
	"bool":    "sql.NullBool",
	"byte":    "sql.NullByte",
	"int8":    "sql.NullByte",
	"int16":   "sql.NullInt16",
	"int":     "sql.NullInt32",
	"int64":   "sql.NullInt64",
	"float32": "sql.NullFloat64",
	"float64": "sql.NullFloat64",
	"string":  "sql.NullString",
}

var GoTypeOps = map[string]string{
	"byte":    ">",
	"int8":    ">",
	"int16":   ">",
	"int32":   ">",
	"int":     ">",
	"int64":   ">",
	"float32": ">",
	"float64": ">",
	"string":  "!=",
}

var GoTypeOpVales = map[string]string{
	">":  "0",
	"!=": "\"\"",
}
