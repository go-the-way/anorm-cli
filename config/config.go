package config

type Configuration struct {
	Module        string
	OutputDir     string
	Verbose       bool
	IncludeTables []string
	Global        *GlobalConfiguration
	Entity        *EntityConfiguration
	MapperEnable  bool
	Config        *CfgConfiguration
}

type GlobalConfiguration struct {
	Author           string
	Date             bool
	DateLayout       string
	Copyright        bool
	CopyrightContent string
	Website          bool
	WebsiteContent   string
}

type EntityConfiguration struct {
	PKG                   string
	TableToEntityStrategy StrategyType
	ColumnToFieldStrategy StrategyType
	FileNameStrategy      StrategyType
	JSONTag               bool
	JSONTagKeyStrategy    StrategyType
	FieldIdUpper          bool
	Comment               bool
	FieldComment          bool
	NamePrefix            string
	NameSuffix            string
	Orm                   bool
}

type CfgConfiguration struct {
	PKG  string
	Name string
}
