package arel

type ToSqlVisitor struct {
	conn *Connection
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

func NewToSqlVisitor(c *Connection) ToSqlVisitor {
	return ToSqlVisitor{conn: c}
}

func (t ToSqlVisitor) Accept(v Visitor) string {
	return ""
}

func (t ToSqlVisitor) Visit(n AstNode, v Visitor) string {
	deleteStmt := "DELETE FROM " + t.Visit(n.Relation, t)
	var whereStmt string

	if len(n.Wheres) > 0 {
		for i, x := range n.Wheres {
			whereStmt += t.Visit(x)
			if i != len(n.Wheres)-1 {
				whereStmt += AND
			}
		}
	}
	return deleteStmt + " " + whereStmt
}
