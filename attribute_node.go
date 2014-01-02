package rel

type AttributeNode struct {
	Name  string
	Table *Table
	BaseVisitable
}

func NewAttributeNode(name string, t *Table) AttributeNode {
	return AttributeNode{
		Name:  name,
		Table: t,
	}
}

func (a AttributeNode) Eq(n SqlLiteralNode) EqualityNode {
	return NewEqualityNode(a, n)
}

func (a AttributeNode) Lt(i int) LessThanNode {
	return LessThanNode{Left: a, Right: Sql(i)}
}

func (a AttributeNode) Gt(i int) GreaterThanNode {
	return GreaterThanNode{Left: a, Right: Sql(i)}
}
