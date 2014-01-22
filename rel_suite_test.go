package rel_test

import (
	. "."
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
	"testing"
)

type BaseNewTestEngine struct {
	visitor Visitor
}

func (e BaseNewTestEngine) Visitor() Visitor {
	return e.visitor
}

func NewTestEngine() *BaseNewTestEngine {
	return &BaseNewTestEngine{
		visitor: ToSqlVisitor{Conn: BaseTestConnector{}},
	}
}

type BaseTestConnector struct{}

func (c BaseTestConnector) Quote(thing interface{}) string {
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
	default:
		return fmt.Sprintf("'%s'", t)
	}
}

func (c BaseTestConnector) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (c BaseTestConnector) QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

func TestRel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rel Suite")
}
