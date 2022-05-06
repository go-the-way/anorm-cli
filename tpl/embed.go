package tpl

import (
	"embed"
	"strings"
	"text/template"
)

//go:embed entity.tpl config.tpl
var FS embed.FS
var entityTpl = `entity.tpl`

var configTpl = `config.tpl`
var entityTplContent = ""

var configTplContent = ""

func EntityTpl() string {
	if entityTplContent == "" {
		file, err := FS.ReadFile(entityTpl)
		if err != nil {
			panic(err)
		}
		entityTplContent = string(file)
	}
	return entityTplContent
}

func ConfigTpl() string {
	if configTplContent == "" {
		file, err := FS.ReadFile(configTpl)
		if err != nil {
			panic(err)
		}
		configTplContent = string(file)
	}
	return configTplContent
}

func ExecuteTpl(tpl string, data map[string]interface{}) string {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	var buffer strings.Builder
	err = t.Execute(&buffer, data)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
