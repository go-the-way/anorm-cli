package bundle

import (
	"github.com/go-the-way/anorm-cli/config"
	"github.com/go-the-way/anorm-cli/entity"
	"github.com/go-the-way/anorm-cli/generator"
)

func GetEntityGenerators(CFG *config.Configuration, tableMap map[string]*entity.Table) []generator.Generator {
	egs := make([]generator.Generator, 0)
	for _, v := range tableMap {
		eg := &generator.EntityGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Init()
		egs = append(egs, eg)
	}
	return egs
}

func GetCfgGenerator(CFG *config.Configuration) generator.Generator {
	return &generator.ConfigGenerator{C: CFG}
}
