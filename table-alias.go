package arel

type TableAliasNode BinaryNode

func TableAliasNodeNew(t *Table, name string) *TableAliasNode {
	return &TableAliasNode{Name: name}
}

func (t *TableAliasNode) Attr(name string) *Attribute {
	return AttributeNew(t.Table, name)
}
