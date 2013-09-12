package arel

type TableAliasNode BinaryNode

func NewTableAliasNode(t *Table, name string) *TableAliasNode {
	return &TableAliasNode{Name: name}
}

func (t *TableAliasNode) Attr(name string) *Attribute {
	return NewAttribute(t.Table, name)
}
