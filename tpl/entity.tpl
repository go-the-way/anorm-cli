package {{.Config.Entity.PKG}}

// @author {{.Config.Global.Author}}
// @since {{.Extra.Date}}
// @created by {{.Config.Global.CopyrightContent}}
// @repo {{.Config.Global.WebsiteContent}}

import (
	"encoding/json"

    "github.com/go-the-way/sg"
    "github.com/go-the-way/anorm"
     _ "{{.Config.Module}}/{{.Config.Config.PKG}}"
)

{{if .Config.Entity.Comment}}// {{.Entity.Name}} struct {{.Entity.Table.Comment}}{{end}}
type {{.Entity.Name}} struct {
    {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{$e.Column.OrmTag}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}} generator:"DB_PRI"`
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} {{$e.Type}} `{{$e.Column.OrmTag}}{{if $e.JSONTag}} json:"{{$e.JSONTagName}}"{{end}}`
    {{end}}
}

func (entity *{{.Entity.Name}}) Configure(c *anorm.EC) {
	c.Table = "{{.Entity.Table.Name}}"
	c.Migrate = true
	c.Commented = true
	c.Comment = "{{.Entity.Table.Comment}}"
	c.IFNotExists = true
	c.IndexDefinitions = []sg.Ge{ {{range $i, $e := .Entity.OrmIndexDefinitions}}
		{{$e}},{{end}}
	}
}

func init(){
	anorm.Register(new({{.Entity.Name}}))
}

// FieldMap entity to map named with fields
func (entity *{{.Entity.Name}}) FieldMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Entity.Ids}}
        "{{$e.Name}}": entity.{{$e.Name}},
        {{end}}{{range $i, $e := .Entity.Fields}}
        "{{$e.Name}}": entity.{{$e.Name}},
        {{end}}
    }
}

// ColumnMap entity to map named with columns
func (entity *{{.Entity.Name}}) ColumnMap() map[string]interface{} {
	return map[string]interface{}{
	    {{range $i, $e := .Entity.Ids}}
        "{{$e.Column.Name}}": entity.{{$e.Name}},
        {{end}}{{range $i, $e := .Entity.Fields}}
        "{{$e.Column.Name}}": entity.{{$e.Name}},
        {{end}}
    }
}

// JSON entity to json
func (entity *{{.Entity.Name}}) JSON() string {
	bytes, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

var {{.Entity.Name}}Columns = &struct{ {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}
}{ {{range $i, $e := .Entity.Ids}}
    {{$e.Name}} : "{{$e.Column.Name}}",
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{$e.Name}} : "{{$e.Column.Name}}",
{{end}}
}

var {{.Entity.Name}}AliasColumns = &struct{ {{range $i, $e := .Entity.Ids}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}// {{$e.Name}} {{$e.Column.Comment}}{{end}}
    {{$e.Name}} sg.C
    {{end}}
}{ {{range $i, $e := .Entity.Ids}}
    {{$e.Name}} : "t.{{$e.Column.Name}}",
    {{end}}{{range $i, $e := .Entity.Fields}}
    {{$e.Name}} : "t.{{$e.Column.Name}}",
{{end}}
}