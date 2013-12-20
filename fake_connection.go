package arel

type Connector struct {
	ConnectionPool *ConnectionPool
}

func (c *Connector) Connection() *Connection {
	return c.ConnectionPool.Connection()
}

type Connection struct {
	Visitor     Visitor
	Tables      []string
	primaryKeys []string
	columns     []string
}

func (c *Connection) TableExists(tableName string) bool {
	for _, name := range c.Tables {
		if tableName == name {
			return true
		}
	}
	return false
}

func QuoteTableName(name string) string {
	return "\"" + name + "\""
}

func QuoteColumnName(name string) string {
	return "\"" + name + "\""
}

type ConnectionPool struct {
	conn *Connection
}

func NewConnectionPool() *ConnectionPool {
	conn := Connection{}
	conn.Visitor = ToSqlVisitor{conn: &conn}
	return &ConnectionPool{conn: &conn}
}

func (c *ConnectionPool) Connection() *Connection {
	return c.conn
}

func (c *ConnectionPool) TableExists(name string) bool {
	return c.Connection().TableExists(name)
}
