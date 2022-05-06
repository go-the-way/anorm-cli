package config

type StrategyType string

var (
	// Default 无，原名称copy					abc => abc
	Default = StrategyType("Default")
	// FirstLetterUpper 仅首字母大写					abc => Abc
	FirstLetterUpper = StrategyType("FirstLetterUpper")
	// UnderlineToCamel 下划线转驼峰（首字母小写）		a_b_c => aBC
	UnderlineToCamel = StrategyType("UnderlineToCamel")
	// UnderlineToUpper 下划线转大写					a_b_c => ABC
	UnderlineToUpper = StrategyType("UnderlineToUpper")
)
