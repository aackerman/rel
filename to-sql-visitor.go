package arel

type ToSqlVisitor struct {
	connection *Connection
	BaseVisitor
}

func NewToSqlVisitor(c *Connection) *ToSqlVisitor {
	return &ToSqlVisitor{
		connection:  c,
		BaseVisitor: BaseVisitor{},
	}
}
