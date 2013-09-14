package arel

type Visitor interface {
	Accept(*SqlAst) string
	Visit(*SqlAst) string
}

type BaseVisitor struct {
	funcs map[ArelNode]func(*Table, *Attribute)
}

func NewVisitor() *BaseVisitor {
	return &BaseVisitor{}
}

func (v *BaseVisitor) dispatch() string {
	return ""
}

func (v *BaseVisitor) Accept(s *SqlAst) string {
	return v.Visit(s)
}

func (v *BaseVisitor) Visit(t *SqlAst) string {
	return v.dispatch()
}
