package entity

type Index struct {
	Table  string `db:"TABLE_NAME"`
	Name   string `db:"INDEX_NAME"`
	Unique int    `db:"INDEX_UNIQUE"`
	Column string `db:"INDEX_COLUMN"`
}
