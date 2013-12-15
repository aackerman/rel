package arel

func Sql(raw string) SqlLiteralNode {
	return NewSqlLiteralNode(raw)
}

func Star() SqlLiteralNode {
	return Sql("*")
}

type SqlLiteralNode struct {
	Str string
	*Predicator
}

func NewSqlLiteralNode(raw string) SqlLiteralNode {
	return SqlLiteralNode{
		raw,
		&Predicator{},
	}
}
