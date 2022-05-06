package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-the-way/anorm-cli/bundle"
	cfg "github.com/go-the-way/anorm-cli/config"
	"github.com/go-the-way/anorm-cli/generator"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g", "generate"},
	Short:   "Generate anorm Go files",
	Long: `Generate anorm Go files.
Simply type anorm-cli help gen for full details.`,
	Example: `anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome"
anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome" -o "/to/path" 
anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome" -au "bigboss" -o "/to/path"`,
	Run: func(cmd *cobra.Command, args []string) {
		CFG.Verbose = verbose

		if dsn == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The DSN is required")
			return
		}

		if module == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The Module name is required")
			return
		}

		if database == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The Database name is required")
			return
		}

		if !entity {
			_, _ = fmt.Fprintln(os.Stderr, "Nothing do...")
			return
		}

		bundle.Init(dsn)
		setCFG()

		tableList := bundle.Tables(database, CFG)
		columnList := bundle.Columns(database)
		indexList := bundle.Indexes(database)
		tableMap := bundle.TransformTables(tableList)
		columnMap := bundle.TransformColumns(columnList)
		bundle.SetTableColumns(tableMap, columnMap)
		indexMap := bundle.TransformIndexes(indexList)
		bundle.SetTableIndexes(tableMap, indexMap)
		generators := make([]generator.Generator, 0)

		entityGenerators := bundle.GetEntityGenerators(CFG, tableMap)
		generators = append(generators, entityGenerators...)

		if config {
			generators = append(generators, bundle.GetCfgGenerator(CFG))
		}

		for _, g := range generators {
			g.Generate()
		}
	},
}

func init() {
	genCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "", "The output dir")
	genCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "", "The MySQL DSN")
	genCmd.PersistentFlags().StringVarP(&database, "database", "D", "", "The Database name")
	genCmd.PersistentFlags().StringVarP(&module, "module", "m", "", "The Module name")
	genCmd.PersistentFlags().StringVarP(&includeTable, "include-table", "I", "", "The include table names[table_a,table_b]")
	genCmd.PersistentFlags().StringVar(&author, "author", "bill", "The file copyright author")
	genCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Print verbose output")

	genCmd.PersistentFlags().BoolVarP(&entity, "entity", "e", true, "Generate Entity Go file")
	genCmd.PersistentFlags().StringVarP(&entityPKG, "entity-pkg", "E", "entity", "The Entity package")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityDefault, "table2entity-default", false, "The Table to Entity name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityFirstLetterUpper, "table2entity-first-letter-upper", false, "The Table to Entity name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToCamel, "table2entity-underline-to-camel", false, "The Table to Entity name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToUpper, "table2entity-underline-to-upper", true, "The Table to Entity name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldDefault, "column2field-default", false, "The column to field name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldFirstLetterUpper, "column2field-first-letter-upper", false, "The column to field name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToCamel, "column2field-underline-to-camel", false, "The column to field name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToUpper, "column2field-underline-to-upper", true, "The column to field name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityFileNameDefault, "entity-filename-default", true, "The Entity file name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityFileNameFirstLetterUpper, "entity-filename-first-letter-upper", false, "The Entity file name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityFileNameUnderlineToCamel, "entity-filename-underline-to-camel", false, "The Entity file name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityFileNameUnderlineToUpper, "entity-filename-underline-to-upper", false, "The Entity file name strategy: UnderlineToUpper")
	genCmd.PersistentFlags().BoolVar(&entityComment, "entity-comment", true, "Generate Entity comment")
	genCmd.PersistentFlags().BoolVar(&entityFieldComment, "entity-field-comment", true, "Generate Entity field comment")
	genCmd.PersistentFlags().BoolVar(&entityJSONTag, "entity-json-tag", true, "Generate Entity field JSON tag")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyDefault, "entity-json-tag-key-default", true, "The Entity JSON Tag key strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyFirstLetterUpper, "entity-json-tag-key-first-letter-upper", false, "The Entity JSON Tag key strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyUnderlineToCamel, "entity-json-tag-key-underline-to-camel", false, "The Entity JSON Tag key strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityJSONTagKeyUnderlineToUpper, "entity-json-tag-key-underline-to-upper", false, "The Entity JSON Tag key strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVarP(&config, "config", "c", true, "Generate Config")
	genCmd.PersistentFlags().StringVarP(&configPKG, "config-pkg", "C", "config", "The Config package")

	rootCmd.AddCommand(genCmd)
}

var (
	outputDir    = ""
	dsn          = ""
	module       = ""
	database     = ""
	includeTable = ""
	author       = ""
	verbose      = false

	entity                             = true
	entityPKG                          = "entity"
	entityTable2EntityDefault          = false
	entityTable2EntityFirstLetterUpper = false
	entityTable2EntityUnderlineToCamel = false
	entityTable2EntityUnderlineToUpper = true

	entityColumn2FieldDefault          = false
	entityColumn2FieldFirstLetterUpper = false
	entityColumn2FieldUnderlineToCamel = false
	entityColumn2FieldUnderlineToUpper = true

	entityFileNameDefault          = true
	entityFileNameFirstLetterUpper = false
	entityFileNameUnderlineToCamel = false
	entityFileNameUnderlineToUpper = false

	entityComment                    = true
	entityFieldComment               = true
	entityJSONTag                    = true
	entityJSONTagKeyDefault          = true
	entityJSONTagKeyFirstLetterUpper = false
	entityJSONTagKeyUnderlineToCamel = false
	entityJSONTagKeyUnderlineToUpper = false

	config    = true
	configPKG = "config"
)

var CFG = &cfg.Configuration{
	Module:        "",
	OutputDir:     "",
	Verbose:       false,
	IncludeTables: make([]string, 0),
	Global: &cfg.GlobalConfiguration{
		Author:           "bill",
		DateLayout:       "2006-01-02",
		CopyrightContent: "An anorm Code generator written in Go",
		WebsiteContent:   "https://github.com/go-the-way/anorm-cli",
	},
	Entity: &cfg.EntityConfiguration{
		PKG:                   "entity",
		TableToEntityStrategy: cfg.UnderlineToUpper,
		ColumnToFieldStrategy: cfg.UnderlineToUpper,
		FileNameStrategy:      cfg.Default,
		JSONTag:               true,
		JSONTagKeyStrategy:    cfg.Default,
		FieldIdUpper:          true,
		Comment:               true,
		FieldComment:          true,
		NamePrefix:            "",
		NameSuffix:            "",
	},
	Config: &cfg.CfgConfiguration{
		PKG:  "config",
		Name: "config",
	},
}

func setCFG() {
	if outputDir != "" {
		CFG.OutputDir = outputDir
	}
	if CFG.OutputDir == "" {
		exec, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		CFG.OutputDir = exec
	}
	if module != "" {
		CFG.Module = module
	}
	if includeTable != "" {
		CFG.IncludeTables = strings.Split(includeTable, ",")
	}

	if author != "" {
		CFG.Global.Author = author
	}

	{
		if entityPKG != "" {
			CFG.Entity.PKG = entityPKG
		}

		CFG.Entity.Comment = entityComment
		CFG.Entity.FieldComment = entityFieldComment
		CFG.Entity.JSONTag = entityJSONTag

		switch {
		case entityTable2EntityUnderlineToUpper:
			CFG.Entity.TableToEntityStrategy = cfg.UnderlineToUpper
		case entityTable2EntityUnderlineToCamel:
			CFG.Entity.TableToEntityStrategy = cfg.UnderlineToCamel
		case entityTable2EntityFirstLetterUpper:
			CFG.Entity.TableToEntityStrategy = cfg.FirstLetterUpper
		case entityTable2EntityDefault:
			CFG.Entity.TableToEntityStrategy = cfg.Default
		}

		switch {
		case entityColumn2FieldUnderlineToUpper:
			CFG.Entity.ColumnToFieldStrategy = cfg.UnderlineToUpper
		case entityColumn2FieldUnderlineToCamel:
			CFG.Entity.ColumnToFieldStrategy = cfg.UnderlineToCamel
		case entityColumn2FieldFirstLetterUpper:
			CFG.Entity.ColumnToFieldStrategy = cfg.FirstLetterUpper
		case entityColumn2FieldDefault:
			CFG.Entity.ColumnToFieldStrategy = cfg.Default
		}

		switch {
		case entityFileNameUnderlineToUpper:
			CFG.Entity.FileNameStrategy = cfg.UnderlineToUpper
		case entityFileNameUnderlineToCamel:
			CFG.Entity.FileNameStrategy = cfg.UnderlineToCamel
		case entityFileNameFirstLetterUpper:
			CFG.Entity.FileNameStrategy = cfg.FirstLetterUpper
		case entityFileNameDefault:
			CFG.Entity.FileNameStrategy = cfg.Default
		}

		switch {
		case entityJSONTagKeyUnderlineToUpper:
			CFG.Entity.JSONTagKeyStrategy = cfg.UnderlineToUpper
		case entityJSONTagKeyUnderlineToCamel:
			CFG.Entity.JSONTagKeyStrategy = cfg.UnderlineToCamel
		case entityJSONTagKeyFirstLetterUpper:
			CFG.Entity.JSONTagKeyStrategy = cfg.FirstLetterUpper
		case entityJSONTagKeyDefault:
			CFG.Entity.JSONTagKeyStrategy = cfg.Default
		}
	}

	{
		if configPKG != "" {
			CFG.Config.PKG = configPKG
		}
	}
}
