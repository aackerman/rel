package rel

type UpdateManager struct {
	Ast UpdateStatementNode
	BaseVisitable
}

func (mgr *UpdateManager) Take(limit int) *UpdateManager {
	if limit > 0 {
		mgr.Ast.Limit = NewLimitNode(Sql(limit))
	}
	return mgr
}

func (mgr *UpdateManager) SetKey(visitable Visitable) *UpdateManager {
	mgr.Ast.Key = visitable
	return mgr
}

func (mgr *UpdateManager) Order(expressions ...Visitable) *UpdateManager {
	mgr.Ast.Orders = &expressions
	return mgr
}

func (mgr *UpdateManager) Table(relation *Table) *UpdateManager {
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

func (mgr *UpdateManager) Set(column AttributeNode, value SqlLiteralNode) *UpdateManager {
	if mgr.Ast.Values == nil {
		mgr.Ast.Values = &[]Visitable{}
	}
	*mgr.Ast.Values = append(*mgr.Ast.Values, &AssignmentNode{
		Left:  &UnqualifiedColumnNode{Expr: &column},
		Right: &value,
	})
	return mgr
}
