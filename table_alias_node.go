package rel

type TableAliasNode struct {
	Name     SqlLiteralNode // Aliased name of the original table
	Relation Visitable      // Generally a *Table, *GroupingNode; a GroupingNode can allow a SelectStatement to be aliased
	BinaryNode
}

func (t *TableAliasNode) Attr(name string) *AttributeNode {
	return NewAttributeNode(t, name)
}

func (t *TableAliasNode) String() string {
	return t.Name.Raw
}
