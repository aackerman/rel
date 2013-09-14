package arel

type AstManager interface {
	ToDot() string
	Visitor() *Visitor
	ToSql() string
	Where(...interface{}) *AstManager
}

type TreeManager struct {
	Engine *Engine
	Ast    *SqlAst
	NodeCreator
}

func NewTreeManager(e *Engine) *TreeManager {
	return &TreeManager{
		NodeCreator: NodeCreator{},
	}
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() *Visitor {
	return t.Engine.Connection().Visitor()
}
