package entity

type Entity struct {
	PKG                 string
	Name                string
	FileName            string
	Comment             bool
	Table               *Table
	Fields              []*Field
	DefaultFields       []*Field
	NoDefaultFields     []*Field
	HaveId              bool
	Ids                 []*Field
	ImportTime          bool
	ImportSql           bool
	Orm                 bool
	IntId               bool
	IdCount             int
	AutoIncrement       bool
	HaveField           bool
	OrmIndexDefinitions []string
}
