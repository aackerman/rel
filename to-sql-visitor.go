package arel

type ToSqlVisitor struct {
	connection *Connection
	BaseVisitor
}

func NewToSqlVisitor(c *Connection) *Visitor {
	return &ToSqlVisitor{
		connection: c,
		BaseVisitor{},
	}
}
