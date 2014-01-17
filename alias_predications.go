package rel

type Aliaser interface {
	As(SqlLiteralNode) *AsNode
	Visitable
}

func aliaserAs(caller Aliaser, literal SqlLiteralNode) *AsNode {
	return &AsNode{
		Left:  caller,
		Right: literal,
	}
}
