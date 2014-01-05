package rel

type TableAliasNode struct {
	Name     SqlLiteralNode // Aliased name of the original table
	Quoted   bool           // Flag to indentify if the alias should be quoted
	Relation Visitable      // Generally a *Table, *GroupingNode; a GroupingNode can allow a SelectStatement to be aliased
	BinaryNode
}

func (t *TableAliasNode) Attr(name string) AttributeNode {
	return NewAttributeNode(t, name)
}
