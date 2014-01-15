package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

type BaseEngine struct {
	pool        ConnectionPool
	visitor     Visitor
	tables      []string
	primaryKeys []string
	columns     []string
}

type ConnectionPool struct {
	conn *Connection
}

func (e BaseEngine) Connection() *Connection {
	return e.pool.Connection()
}

func (e BaseEngine) Visitor() Visitor {
	return e.visitor
}

func NewEngine() *BaseEngine {
	e := new(BaseEngine)
	e.pool = ConnectionPool{conn: new(Connection)}
	e.visitor = ToSqlVisitor{Conn: e.pool.Connection()}
	return e
}

func (e BaseEngine) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (e BaseEngine) QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

func (e BaseEngine) TableExists(tableName string) bool {
	for _, name := range e.tables {
		if tableName == name {
			return true
		}
	}
	return false
}

func (p *ConnectionPool) Connection() *Connection {
	return p.conn
}

func TestRel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rel Suite")
}
