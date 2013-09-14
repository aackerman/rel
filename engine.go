package arel

type Engine struct {
	pool *ConnectionPool
}

func NewEngine() *Engine {
	return &Engine{
		pool: NewConnectionPool(),
	}
}

func (e *Engine) Connection() *Connection {
	return e.pool.Connection()
}
