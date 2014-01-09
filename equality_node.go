package rel

type EqualityNode struct {
	Left  Visitable
	Right Visitable
	BaseVisitable
}

func NewEqualityNode(left Visitable, right Visitable) EqualityNode {
	return EqualityNode{
		Left:  left,
		Right: right,
	}
}

type InNode struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}

type NotInNode struct {
	Left  Visitable
	Right []Visitable
	BaseVisitable
}
