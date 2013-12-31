package grel

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

func (n AttributeNode) Eq(b interface{}) EqualityNode {
	return NewEqualityNode(n, Sql(b))
}
