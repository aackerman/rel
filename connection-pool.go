package arel

type ConnectionPool struct {
	config     *ConnectionConfig
	connection *Connection
}

func NewConnectionPool() *ConnectionPool {
	connection := &Connection{}
	connection.visitor = &ToSqlVisitor{
		connection: connection,
	}

	return &ConnectionPool{
		connection: connection,
	}
}

func (c *ConnectionPool) Connection() *Connection {
	return c.connection
}

func (c *ConnectionPool) TableExists(name string) bool {
	return c.Connection().TableExists(name)
}
