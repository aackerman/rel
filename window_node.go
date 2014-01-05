package rel

type WindowNode struct {
	BaseVisitable
}

type NamedWindowNode struct {
	Name SqlLiteralNode
	BaseVisitable
}
