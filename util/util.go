package util

import (
	"strings"

	"github.com/go-the-way/anorm-cli/config"
)

func ConvertString(str string, st config.StrategyType) string {
	switch st {
	default:
		return str
	case "None":
		return str
	case "OnlyFirstLetterUpper":
		return onlyFirstLetterUpper(str)
	case "UnderlineToCamel":
		return underlineToCamel(str)
	case "UnderlineToUpper":
		return underlineToUpper(str)
	}
}
func onlyFirstLetterUpper(str string) string {
	if len(str) == 1 {
		return strings.ToUpper(string(str[0]))
	} else if len(str) > 1 {
		return strings.ToUpper(string(str[0])) + str[1:]
	}
	return str
}

// a_b_c_d
func underlineToCamel(str string) string {
	names := strings.Split(str, "_")
	for i, name := range names {
		if i > 0 {
			if len(name) == 1 {
				names[i] = strings.ToUpper(string(name[0]))
			} else if len(name) > 1 {
				names[i] = strings.ToUpper(string(name[0])) + name[1:]
			}
		}
	}
	return strings.Join(names, "")
}

func underlineToUpper(str string) string {
	toCamelStr := underlineToCamel(str)
	return strings.ToUpper(string(toCamelStr[0])) + toCamelStr[1:]
}
