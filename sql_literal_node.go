package arel

func Sql(raw string) *SqlLiteralNode {
	return NewSqlLiteralNode(raw)
}

func Star() *SqlLiteralNode {
	return Sql("*")
}

// include Arel::Expressions
// include Arel::AliasPredication
// include Arel::OrderPredications
type SqlLiteralNode struct {
	Str string
	*Predicator
}

func NewSqlLiteralNode(str string) *SqlLiteralNode {
	return &SqlLiteralNode{
		str,
		&Predicator{},
	}
}
