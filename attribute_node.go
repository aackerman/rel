package rel

type AttributeNode struct {
	Name  string
	Table *Table
	BaseNode
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
