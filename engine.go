package arel

// An engine handles connections and string quoting
// ActiveRecord is an example of an engine used in Arel
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
