package arel

// include Arel::Expressions
// include Arel::AliasPredication
// include Arel::OrderPredications
type SqlLiteralNode struct {
	Str string
	*Predicator
}

func SqlLiteralNodeNew(str string) *SqlLiteralNode {
	return &SqlLiteralNode{
		str,
		&Predicator{},
	}
}
