package grel

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

type Connection struct{}

func (e BaseEngine) Connection() *Connection {
	return e.pool.Connection()
}

func (e BaseEngine) Visitor() Visitor {
	return e.visitor
}

var DefaultEngine BaseEngine = NewDefaultEngine()

func NewDefaultEngine() BaseEngine {
	e := BaseEngine{}
	e.pool = ConnectionPool{conn: new(Connection)}
	e.visitor = ToSqlVisitor{conn: e.pool.Connection()}
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
