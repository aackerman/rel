package arel

func Sql(raw string) *SqlLiteralNode {
	return SqlLiteralNodeNew(raw)
}

func Star() *SqlLiteralNode {
	return Sql("*")
}
