package arel

type TableAliasNode struct {
	Name string
	BinaryNode
}

func NewTableAliasNode(t *Table, name string) TableAliasNode {
	return TableAliasNode{Name: name}
}

func (t *TableAliasNode) Attr(name string) Attribute {
	return NewAttribute(t.Table, name)
}
