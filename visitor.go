package arel

type Visitor interface {
	Accept(*NodeCreator) string
	Visit(*NodeCreator) string
}

type BaseVisitor struct{}

func (v *BaseVisitor) Accept(n *NodeCreator) string {
	return v.Visit(n)
}

func (v *BaseVisitor) Visit(n *NodeCreator) string {
	return "hello world!"
}
