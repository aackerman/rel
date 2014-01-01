package rel

type EqualityNode struct {
	Left  Visitable
	Right *Visitable
	BaseNode
}

func NewEqualityNode(left Visitable, right Visitable) EqualityNode {
	return EqualityNode{
		Left:  left,
		Right: &right,
	}
}

type InNode EqualityNode
