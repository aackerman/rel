package rel

type BindParamNode struct {
	Raw string
	BaseVisitable
}

func NewBindParamNode(raw string) *BindParamNode {
	return &BindParamNode{Raw: raw}
}

func (node *BindParamNode) String() string {
	if node == nil {
		return "NULL"
	} else {
		return node.Raw
	}
}
