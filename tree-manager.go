package arel

type AstManager interface {
	ToDot() string
	Visitor() *Visitor
	ToSql() string
	Where(...interface{}) *AstManager
}

type TreeManager struct {
	FactoryMethods
	Engine *Engine
	Ast    *SqlAst
}

func NewTreeManager(e *Engine) *TreeManager {
	return &TreeManager{
		FactoryMethods{},
	}
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() *Visitor {
	return t.Engine.Connection().Visitor()
}
