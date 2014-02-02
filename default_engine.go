package rel

import (
	"fmt"
	"strconv"
)

var RelEngine Engine = &DefaultEngine{
	visitor: &ToSqlVisitor{Conn: DefaultConnector{}},
}

type DefaultEngine struct {
	visitor Visitor
}

func (e DefaultEngine) Visitor() Visitor {
	return e.visitor
}

type DefaultConnector struct{}

func (c DefaultConnector) Quote(thing interface{}) string {
	switch t := thing.(type) {
	case bool:
		if t {
			return "'t'"
		} else {
			return "'f'"
		}
	case int:
		return strconv.Itoa(t)
	case nil:
		return "NULL"
	case SqlLiteralNode:
		return t.Raw
	case *SqlLiteralNode:
		return t.Raw
	case *BindParamNode:
		if t != nil {
			return t.Raw
		} else {
			return "NULL"
		}
	default:
		return fmt.Sprintf("'%s'", t)
	}
}

func (c DefaultConnector) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (c DefaultConnector) QuoteColumnName(name string) string {
	return "\"" + name + "\""
}
