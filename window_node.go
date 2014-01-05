package rel

type WindowNode struct {
	BaseVisitable
}

type NamedWindowNode struct {
	Name    SqlLiteralNode
	Orders  *[]Visitable
	Framing Visitable
	BaseVisitable
}
