package arel

type Visitor struct {
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) Accept(a interface{}) {
	v.Visit(a)
}

func (v *Visitor) Visit(a interface{}) {

}
