package entity

type Column struct {
	Table         string `db:"TABLE_NAME"`
	Name          string `db:"COLUMN_NAME"`
	Type          string `db:"COLUMN_TYPE"`
	DataType      string `db:"COLUMN_DATA_TYPE"`
	ColumnKey     string `db:"COLUMN_KEY"`
	Comment       string `db:"COLUMN_COMMENT"`
	Default       string `db:"COLUMN_DEFAULT"`
	AutoIncrement int    `db:"AUTO_INCREMENT" `
	NotNull       int    `db:"NOT_NULL"`
	OrmTag        string
}
