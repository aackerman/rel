package rel

type UpdateManager struct {
	Engine Engine
	Ast    *UpdateStatementNode
	BaseVisitable
}

func NewUpdateManager(engine Engine) *UpdateManager {
	return &UpdateManager{
		Engine: engine,
		Ast:    NewUpdateStatementNode(),
	}
}

func (mgr *UpdateManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func (mgr *UpdateManager) Take(limit int) *UpdateManager {
	if limit > 0 {
		mgr.Ast.Limit = NewLimitNode(Sql(limit))
	}
	return mgr
}

func (mgr *UpdateManager) SetKey(node *AttributeNode) *UpdateManager {
	mgr.Ast.Key = node
	return mgr
}

func (mgr *UpdateManager) Order(expressions ...Visitable) *UpdateManager {
	mgr.Ast.Orders = &expressions
	return mgr
}

func (mgr *UpdateManager) SetTable(relation *Table) *UpdateManager {
	mgr.Ast.Relation = relation
	return mgr
}

func (mgr *UpdateManager) Where(visitable Visitable) *UpdateManager {
	if mgr.Ast.Wheres == nil {
		mgr.Ast.Wheres = &[]Visitable{}
	}
	*mgr.Ast.Wheres = append(*mgr.Ast.Wheres, visitable)
	return mgr
}

func (mgr *UpdateManager) Set(field *AttributeNode, value *BindParamNode) *UpdateManager {
	if mgr.Ast.Values == nil {
		mgr.Ast.Values = &[]Visitable{}
	}
	*mgr.Ast.Values = append(*mgr.Ast.Values, &AssignmentNode{
		Left:  &UnqualifiedColumnNode{Expr: field},
		Right: value,
	})
	return mgr
}
