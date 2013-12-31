package rel

type TableAliasNode struct {
	Name  string
	Table *Table
	BinaryNode
}

func NewTableAliasNode(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name, Table: t}
}

func (t *TableAliasNode) Attr(name string) AttributeNode {
	return NewAttributeNode(name, t.Table)
}
