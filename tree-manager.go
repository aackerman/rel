package arel

type AstManager interface {
	ToDot() string
	Visitor() *Visitor
	ToSql() string
	Where(...interface{}) *AstManager
}

type TreeManager struct {
	Engine *Engine
	Ast    NodeCreator
	BaseNodeCreator
}

func NewTreeManager(e *Engine) *TreeManager {
	return &TreeManager{
		Engine:          e,
		BaseNodeCreator: BaseNodeCreator{},
	}
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() *Visitor {
	return t.Engine.Connection().Visitor()
}
