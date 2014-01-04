package rel

type TableAliasNode struct {
	Name  string
	Table *Table
	BinaryNode
}

func (t *TableAliasNode) Attr(name string) AttributeNode {
	return NewAttributeNode(t, name)
}
