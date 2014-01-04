package rel

type AttributeNode struct {
	Name     string
	Relation Visitable
	BaseVisitable
}

func NewAttributeNode(v Visitable, name string) AttributeNode {
	return AttributeNode{
		Name:     name,
		Relation: v,
	}
}

func (a AttributeNode) Eq(n Visitable) EqualityNode {
	return NewEqualityNode(a, n)
}

func (a AttributeNode) Lt(i int) LessThanNode {
	return LessThanNode{Left: a, Right: Sql(i)}
}

func (a AttributeNode) Gt(i int) GreaterThanNode {
	return GreaterThanNode{Left: a, Right: Sql(i)}
}

func (a AttributeNode) In(v Visitable) Visitable {
	var ret Visitable
	switch val := v.(type) {
	case SelectManager:
		ret = &InNode{Left: a, Right: val.Ast}
	default:
		ret = &InNode{Left: a, Right: v}
	}
	return ret
}
