package rel

type OverNode BinaryNode

func (node *OverNode) As(literal SqlLiteralNode) *AsNode {
	return aliasPredicationAs(node, literal)
}
