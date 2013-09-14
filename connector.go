package arel

type Connector struct {
	ConnectionPool *ConnectionPool
}

func (c *Connector) Connection() *Connection {
	return c.ConnectionPool.Connection()
}
