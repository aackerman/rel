package arel

func Sql(raw string) *SqlLiteralNode {
	return NewSqlLiteralNode(raw)
}

func Star() *SqlLiteralNode {
	return Sql("*")
}
