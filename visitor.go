package arel

type Visitor interface {
	Accept(*SqlStatement) string
	Visit(*SqlStatement) string
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

func (v *BaseVisitor) Accept(s *SqlStatement) string {
	return v.Visit(s)
}

func (v *BaseVisitor) Visit(t *SqlStatement) string {
	return v.dispatch()
}
