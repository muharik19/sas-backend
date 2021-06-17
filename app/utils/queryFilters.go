package utils

import (
	"fmt"
	"strings"
)

var logicOperator string = "and"

// SetLogicOperator ...
func SetLogicOperator(logicOperatorArg string) {
	logicOperator = logicOperatorArg
}

// GenerateFilter ...
func GenerateFilter(colName string, operator string, value string, dataType string) string {
	var whereCond string

	switch operator {
	case "like":
		val := `'%` + value + `%'`
		whereCond = fmt.Sprintf(` %s  %s like %s`, logicOperator, colName, val)
	case "not_like":
		val := `'%` + value + `%'`
		whereCond = fmt.Sprintf(` %s  %s not like %s`, logicOperator, colName, val)
	case "less_than":
		val := fmt.Sprintf(`%s`, value)
		if dataType == "date" {
			val = fmt.Sprintf(`'%s'`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s < %s`, logicOperator, colName, val)
	case "less_than_equal":
		val := fmt.Sprintf(`%s`, value)
		if dataType == "date" {
			val = fmt.Sprintf(`'%s'`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s <= %s`, logicOperator, colName, val)
	case "greater_than":
		val := fmt.Sprintf(`%s`, value)
		if dataType == "date" {
			val = fmt.Sprintf(`'%s'`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s > %s`, logicOperator, colName, val)
	case "greater_than_equal":
		val := fmt.Sprintf(`%s`, value)
		if dataType == "date" {
			val = fmt.Sprintf(`'%s'`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s >= %s`, logicOperator, colName, val)
	case "not_eq":
		val := fmt.Sprintf(`'%s'`, value)
		if dataType == "int" || dataType == "float" {
			val = fmt.Sprintf(`%s`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s <> %s`, logicOperator, colName, val)
	case "between":
		val := strings.Split(value, ",")
		whereCond = fmt.Sprintf(` %s %s between '%s' and '%s'`, logicOperator, colName, val[0], val[1])
	case "in":
		val := strings.Split(value, ",")
		totalData := len(val)
		if totalData > 0 {
			inValue := ""
			i := 1
			commaSep := ","
			for _, valueItem := range val {
				if i >= totalData {
					commaSep = ""
				}
				if dataType == "int" || dataType == "float" {
					inValue += fmt.Sprintf(`%s%s`, valueItem, commaSep)
				} else {
					inValue += fmt.Sprintf(`'%s'%s`, valueItem, commaSep)
				}
				i++
			}
			whereCond = fmt.Sprintf(` %s %s in (%s)`, logicOperator, colName, inValue)
		}
	default:
		val := fmt.Sprintf(`'%s'`, value)
		if dataType == "int" || dataType == "float" {
			val = fmt.Sprintf(`%s`, value)
		}
		whereCond = fmt.Sprintf(` %s  %s = %s`, logicOperator, colName, val)
	}

	logicOperator = "and"
	return whereCond
}
