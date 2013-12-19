package arel

func Sql(raw string) SqlLiteralNode {
	return NewSqlLiteralNode(raw)
}

func Star() SqlLiteralNode {
	return Sql("*")
}

type SqlLiteralNode struct {
	Raw string
	Predicator
	BaseNode
}

func NewSqlLiteralNode(raw string) SqlLiteralNode {
	return SqlLiteralNode{
		raw,
		Predicator{},
		BaseNode{},
	}
}
