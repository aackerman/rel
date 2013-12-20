package arel

// An engine handles connections and string quoting
// ActiveRecord is an example of an engine used in Arel
type Engine struct {
	connector *Connector
}

func NewEngine() Engine {
	return Engine{
		connector: &Connector{
			ConnectionPool: NewConnectionPool(),
		},
	}
}

func (e *Engine) Connection() *Connection {
	return e.connector.Connection()
}
