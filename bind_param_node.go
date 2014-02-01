package rel

type BindParamNode struct {
	Raw string
	BaseVisitable
}

func NewBindParamNode(raw string) *BindParamNode {
	return &BindParamNode{Raw: raw}
}
