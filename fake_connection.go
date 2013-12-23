package arel

type Engine interface {
	Visitor() Visitor
	Connection() *Connection
	QuoteTableName(string) string
	QuoteColumnName(string) string
	TableExists(string) bool
}

type BaseEngine struct {
	Pool        ConnectionPool
	visitor     Visitor
	tables      []string
	primaryKeys []string
	columns     []string
}

type ConnectionPool struct {
	conn *Connection
}

type Connection struct{}

func (e BaseEngine) Connection() *Connection {
	return e.Pool.Connection()
}

func (e BaseEngine) Visitor() Visitor {
	return e.visitor
}

var DefaultEngine BaseEngine = CreateDefaultEngine()

func CreateDefaultEngine() BaseEngine {
	e := BaseEngine{}
	e.Pool = ConnectionPool{conn: new(Connection)}
	e.visitor = ToSqlVisitor{conn: e.Pool.Connection()}
	return e
}

func (e BaseEngine) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (e BaseEngine) QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

func (c *Connection) QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func (c *Connection) QuoteColumnName(name string) string {
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
