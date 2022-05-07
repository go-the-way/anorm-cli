package bundle

import (
	"fmt"
	"github.com/go-the-way/anorm-cli/config"
	"github.com/go-the-way/anorm-cli/entity"
	"strings"
)

func Tables(database string, c *config.Configuration) []*entity.Table {
	whereSql := ""
	if c.IncludeTables != nil && len(c.IncludeTables) > 0 {
		whereSql = fmt.Sprintf("AND t.`TABLE_NAME` IN('%s')", strings.Join(c.IncludeTables, "','"))
	}
	tableList := SelectTableListSelectMapper.Prepare(map[string]interface{}{
		"DBName": database,
		"Where":  whereSql,
	}).Exec().List(new(entity.Table))
	ts := make([]*entity.Table, len(tableList))
	for i, t := range tableList {
		tt := t.(*entity.Table)
		tt.Columns = make([]*entity.Column, 0)
		tt.Indexes = make([]*entity.Index, 0)
		ts[i] = tt
	}
	return ts
}

func Columns(database string) []*entity.Column {
	columnList := SelectTableColumnListSelectMapper.Prepare(database).Exec().List(new(entity.Column))
	cs := make([]*entity.Column, len(columnList))
	for i, c := range columnList {
		cc := c.(*entity.Column)
		pk := "F"
		insertIgnore := "F"
		updateIgnore := "F"
		autoIncrement := ""
		notNull := "NULL"
		defaultStr := ""
		defaultPre := ""
		defaultSuf := ""
		if cc.AutoIncrement == 1 {
			pk = "T"
			insertIgnore = "T"
			autoIncrement = " auto_increment"
		}
		if strings.Contains(cc.ColumnKey, "PRI") {
			pk = "T"
		}
		if cc.NotNull == 1 {
			notNull = "NOT NULL"
		}
		if cc.Default != "__NULL__" {
			if cc.Default != "CURRENT_TIMESTAMP" {
				defaultPre = "'"
				defaultSuf = "'"
			} else {
				insertIgnore = "T"
				updateIgnore = "T"
			}
			defaultStr = " default " + defaultPre + strings.ToLower(cc.Default) + defaultSuf
		}
		cc.OrmTag = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("orm:\"pk{%s} c{%s} ig{%s} ug{%s} def{%s}\"",
			pk, cc.Name, insertIgnore, updateIgnore, fmt.Sprintf("%s %s %s%s%s comment '%s'", cc.Name, cc.DataType, notNull, autoIncrement, defaultStr, cc.Comment)), "pk{F} ", ""), "ig{F} ", ""), "ug{F} ", "")
		cs[i] = cc
	}
	return cs
}

func Indexes(database string) []*entity.Index {
	indexList := SelectTableIndexListSelectMapper.Prepare(database).Exec().List(new(entity.Index))
	is := make([]*entity.Index, len(indexList))
	for i, c := range indexList {
		ii := c.(*entity.Index)
		is[i] = ii
	}
	return is
}

func TransformTables(tables []*entity.Table) map[string]*entity.Table {
	tableMap := make(map[string]*entity.Table, len(tables))
	for _, t := range tables {
		tableMap[t.Name] = t
	}
	return tableMap
}

func TransformColumns(columns []*entity.Column) map[string]*[]*entity.Column {
	columnMap := make(map[string]*[]*entity.Column, 0)
	for _, c := range columns {
		cs, have := columnMap[c.Table]
		if have {
			*cs = append(*cs, c)
		} else {
			csp := make([]*entity.Column, 1)
			csp[0] = c
			columnMap[c.Table] = &csp
		}
	}
	return columnMap
}

func TransformIndexes(indexes []*entity.Index) map[string]*[]*entity.Index {
	indexMap := make(map[string]*[]*entity.Index, 0)
	for _, i := range indexes {
		cs, have := indexMap[i.Table]
		if have {
			*cs = append(*cs, i)
		} else {
			csp := make([]*entity.Index, 1)
			csp[0] = i
			indexMap[i.Table] = &csp
		}
	}
	return indexMap
}

func SetTableColumns(tableMap map[string]*entity.Table, columnMap map[string]*[]*entity.Column) {
	for k, v := range tableMap {
		if cc, have := columnMap[k]; have {
			v.Columns = append(v.Columns, *cc...)
		}
	}
}

func SetTableIndexes(tableMap map[string]*entity.Table, indexMap map[string]*[]*entity.Index) {
	for k, v := range tableMap {
		if vv, have := indexMap[k]; have {
			v.Indexes = append(v.Indexes, *vv...)
		}
	}
}
