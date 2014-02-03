package rel

type DeleteManager struct {
	Engine Engine
	Ast    *DeleteStatementNode
	BaseVisitable
}

func NewDeleteManager(engine Engine) *DeleteManager {
	return &DeleteManager{
		Engine: engine,
		Ast:    NewDeleteStatementNode(),
	}
}

func (mgr *DeleteManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func (mgr *DeleteManager) From(table *Table) *DeleteManager {
	mgr.Ast.Relation = table
	return mgr
}

func (mgr *DeleteManager) Where(visitable Visitable) *DeleteManager {
	if mgr.Ast.Wheres == nil {
		mgr.Ast.Wheres = &[]Visitable{}
	}
	*mgr.Ast.Wheres = append(*mgr.Ast.Wheres, visitable)
	return mgr
}
