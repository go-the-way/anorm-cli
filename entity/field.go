package entity

type Field struct {
	Name             string
	Type             string
	OpName           string
	OpVar            string
	Comment          bool
	ColumnAnnotation bool
	ColumnDefinition string
	JSONTag          bool
	JSONTagName      string
	Column           *Column
	HaveDefault      bool
	IgnoreDefault    bool
	Default          string
	Orm              bool
}
