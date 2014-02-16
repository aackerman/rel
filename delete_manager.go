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

func (mgr *DeleteManager) From(table interface{}) *DeleteManager {
	switch t := table.(type) {
	case *Table:
		mgr.Ast.Relation = t
	case Table:
		mgr.Ast.Relation = &t
	case string:
		mgr.Ast.Relation = NewTable(t)
	}
	return mgr
}

func (mgr *DeleteManager) FromTable(table *Table) *DeleteManager {
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
