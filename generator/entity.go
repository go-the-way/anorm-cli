package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-the-way/anorm-cli/config"
	"github.com/go-the-way/anorm-cli/entity"
	"github.com/go-the-way/anorm-cli/tpl"
	"github.com/go-the-way/anorm-cli/util"
)

var entityGeneratorLogger = log.New(os.Stdout, "[EntityGenerator]", log.LstdFlags)

type EntityGenerator struct {
	C      *config.Configuration
	Table  *entity.Table
	Entity *entity.Entity
	Body   string
}

func (eg *EntityGenerator) Generate() {
	eg.generateBody()
	eg.generateFile()
}

func (eg *EntityGenerator) Init() *EntityGenerator {
	eg.Entity = &entity.Entity{
		PKG:             eg.C.Entity.PKG,
		Table:           eg.Table,
		Ids:             make([]*entity.Field, 0),
		Fields:          make([]*entity.Field, 0),
		DefaultFields:   make([]*entity.Field, 0),
		NoDefaultFields: make([]*entity.Field, 0),
		Comment:         eg.C.Entity.Comment,
		Orm:             eg.C.Entity.Orm,
	}
	eg.Entity.Name = util.ConvertString(eg.Table.Name, eg.C.Entity.TableToEntityStrategy)
	eg.Entity.FileName = util.ConvertString(eg.Table.Name, eg.C.Entity.FileNameStrategy)
	autoIncrement := false
	for _, column := range eg.Table.Columns {
		field := &entity.Field{
			Name:    util.ConvertString(column.Name, eg.C.Entity.ColumnToFieldStrategy),
			Type:    map[int]string{0: config.GoSqlNullTypes[config.MysqlToGoTypes[column.Type]], 1: config.MysqlToGoTypes[column.Type]}[column.NotNull],
			Column:  column,
			Comment: eg.C.Entity.FieldComment,
			Orm:     eg.C.Entity.Orm,
		}
		field.OpName = config.GoTypeOps[field.Type]
		field.OpVar = config.GoTypeOpVales[field.OpName]
		if column.ColumnKey == "PRI" {
			eg.Entity.HaveId = true
			eg.Entity.Ids = append(eg.Entity.Ids, field)
		} else {
			eg.Entity.Fields = append(eg.Entity.Fields, field)
		}
		if column.Default != "__NULL__" {
			field.HaveDefault = true
			field.Default = column.Default
			switch {
			case strings.HasPrefix(field.Type, "uint") ||
				strings.HasPrefix(field.Type, "int") ||
				strings.HasPrefix(field.Type, "float"):
				fv, err := strconv.ParseFloat(field.Default, 64)
				if err == nil && fv == 0 {
					field.IgnoreDefault = true
				}
			case field.Default == "CURRENT_TIMESTAMP" && field.Type == "string":
				field.Default = `time.Now().Format("2006-01-02 15:04:05")`
				eg.Entity.ImportTime = true
			case field.Default == "CURRENT_TIMESTAMP" && (field.Type == "time.Time" || field.Type == "*time.Time"):
				field.Default = `time.Now()`
				eg.Entity.ImportTime = true
			case field.Type == "string":
				field.Default = fmt.Sprintf("\"%s\"", field.Default)
				if field.Default == "\"\"" {
					field.IgnoreDefault = true
				}
			}
		}
		if !autoIncrement {
			autoIncrement = column.AutoIncrement == 1
		}
		if eg.C.Entity.FieldIdUpper {
			switch {
			case strings.LastIndex(field.Name, "Id") != -1:
				field.Name = strings.TrimSuffix(field.Name, "Id") + "ID"
			case strings.LastIndex(field.Name, "id") != -1:
				field.Name = strings.TrimSuffix(field.Name, "id") + "ID"
			}
		}
		if eg.C.Entity.JSONTag {
			field.JSONTag = true
			field.JSONTagName = util.ConvertString(column.Name, eg.C.Entity.JSONTagKeyStrategy)
		}
		if column.NotNull == 0 {
			eg.Entity.ImportSql = true
		}
	}

	for _, field := range eg.Entity.Fields {
		if field.HaveDefault {
			if !field.IgnoreDefault {
				eg.Entity.DefaultFields = append(eg.Entity.DefaultFields, field)
			}
		} else {
			eg.Entity.NoDefaultFields = append(eg.Entity.NoDefaultFields, field)
		}
	}
	if !eg.Entity.HaveId {
		panic(fmt.Sprintf("Table of [%s] required at least one PRI column", eg.Entity.Table.Name))
	}
	eg.Entity.IntId = strings.HasPrefix(eg.Entity.Ids[0].Type, "int")
	eg.Entity.IdCount = len(eg.Entity.Ids)
	eg.Entity.AutoIncrement = autoIncrement
	eg.Entity.HaveField = len(eg.Entity.Fields) > 0

	ormIndexDefinitions := make([]string, 0)
	for _, ii := range eg.Table.Indexes {
		ormIndexDefinitions = append(ormIndexDefinitions,
			fmt.Sprintf(`sg.IndexDefinition(%v, sg.P("%s"), %s)`, ii.Unique == 1, ii.Name, `sg.C("`+strings.ReplaceAll(ii.Column, ",", `"), sg.C("`)+`")`))
	}
	eg.Entity.OrmIndexDefinitions = ormIndexDefinitions
	return eg
}

func (eg *EntityGenerator) generateBody() {
	eg.Body = tpl.ExecuteTpl(tpl.EntityTpl(), map[string]interface{}{
		"Entity": eg.Entity,
		"Config": eg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(eg.C.Global.DateLayout),
		},
	})
	if eg.C.Verbose {
		entityGeneratorLogger.Println(fmt.Sprintf("[generateBody] for entity[%s]", eg.Entity.Name))
	}
}

func (eg *EntityGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, eg.C.OutputDir)
	paths = append(paths, eg.Entity.PKG)
	paths = append(paths, eg.Entity.FileName)
	fileName := filepath.Join(paths...) + ".go"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(eg.Body), 0700)
	if eg.C.Verbose {
		entityGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", eg.Entity.Name, fileName))
	}
}
