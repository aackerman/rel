package arel

type ToSqlVisitor struct {
	connection *Connection
	*Visitor
}

func NewToSqlVisitor(c *Connection) *ToSqlVisitor {
	return &ToSqlVisitor{
		c,
		NewVisitor(),
	}
}
