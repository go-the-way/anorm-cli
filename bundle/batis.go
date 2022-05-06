package bundle

import (
	. "github.com/billcoding/gobatis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	SelectTableListSelectMapper       *SelectMapper
	SelectTableColumnListSelectMapper *SelectMapper
	SelectTableIndexListSelectMapper  *SelectMapper
)

func Init(dsn string) {
	Default().DSN(dsn)

	Default().AddRaw(tableXML)

	SelectTableListSelectMapper = NewHelper("table", "SelectTableList").Select()
	SelectTableColumnListSelectMapper = NewHelper("table", "SelectTableColumnList").Select()
	SelectTableIndexListSelectMapper = NewHelper("table", "SelectTableIndexList").Select()
}
