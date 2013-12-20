package arel

type ToSqlVisitor struct {
	conn *Connection
	BaseVisitor
}

const (
	WHERE    = " WHERE "
	SPACE    = " "
	COMMA    = ", "
	GROUP_BY = " GROUP BY "
	ORDER_BY = " ORDER BY "
	WINDOW   = " WINDOW "
	AND      = " AND "
	DISTINCT = "DISTINCT"
)

func NewToSqlVisitor(c *Connection) ToSqlVisitor {
	return ToSqlVisitor{conn: c, BaseVisitor: BaseVisitor{}}
}
