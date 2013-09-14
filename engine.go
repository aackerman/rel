package arel

type Engine struct {
	Connector *Connector
}

func NewEngine() *Engine {
	return &Engine{
		Connector: &Connector{
			ConnectionPool: NewConnectionPool(),
		},
	}
}

func (e *Engine) Connection() *Connection {
	return e.Connector.Connection()
}
