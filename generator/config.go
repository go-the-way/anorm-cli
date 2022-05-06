package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-the-way/anorm-cli/config"
	"github.com/go-the-way/anorm-cli/tpl"
)

var configGeneratorLogger = log.New(os.Stdout, "[ConfigGenerator]", log.LstdFlags)

type ConfigGenerator struct {
	C    *config.Configuration
	Body string
}

func (cg *ConfigGenerator) Generate() {
	cg.generateBody()
	cg.generateFile()
}

func (cg *ConfigGenerator) generateBody() {
	cg.Body = tpl.ExecuteTpl(tpl.ConfigTpl(), map[string]interface{}{
		"Config": cg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(cg.C.Global.DateLayout),
		},
	})
	if cg.C.Verbose {
		configGeneratorLogger.Println(fmt.Sprintf("[generateBody] for config[%s]", cg.C.Config.Name))
	}
}

func (cg *ConfigGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, cg.C.OutputDir)
	paths = append(paths, cg.C.Config.PKG)
	paths = append(paths, cg.C.Config.Name)
	fileName := filepath.Join(paths...) + ".go"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(cg.Body), 0700)
	if cg.C.Verbose {
		configGeneratorLogger.Println(fmt.Sprintf("[generateFile] for config[%s], saved as [%s]", cg.C.Config.Name, fileName))
	}
}
