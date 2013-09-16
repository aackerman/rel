package arel

type ToSqlVisitor struct {
	connection *Connection
}

const (
	WHERE    = " WHERE "
	COMMA    = ", "
	GROUP_BY = " GROUP BY "
	ORDER_BY = " ORDER BY "
	WINDOW   = " WINDOW "
	AND      = " AND "
	DISTINCT = "DISTINCT"
)

func NewToSqlVisitor(c *Connection) *ToSqlVisitor {
	return &ToSqlVisitor{connection: c}
}

func (t ToSqlVisitor) Accept(v Visitor) string {
	return ""
}

func (t ToSqlVisitor) Visit(n Node, v Visitor) string {
	deletestr := "DELETE FROM " + t.Visit(n.Relation, t)
	var wherestr string

	// compile where string
	if len(n.Wheres) > 0 {
		for i, x := range n.Wheres {
			wherestr += t.Visit(x)
			if i != len(n.Wheres)-1 {
				wherestr += AND
			}
		}
	}
	return deletestr + " " + wherestr
}
