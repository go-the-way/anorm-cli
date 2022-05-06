package {{.Config.Config.PKG}}

// @author {{.Config.Global.Author}}
// @since {{.Extra.Date}}
// @created by {{.Config.Global.CopyrightContent}}
// @repo {{.Config.Global.WebsiteContent}}

import (
	"database/sql"
	"github.com/go-the-way/anorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN = ""
)

func init() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	anorm.DataSourcePool.Push(db)
}
