package rel

type SQLiteVisitor struct {
	ToSqlVisitor
}

func (v SQLiteVisitor) Accept(a Visitable) string {
	return v.Visit(a)
}

func (v SQLiteVisitor) Visit(a Visitable) string {
	switch val := a.(type) {
	case *LockNode:
		return v.VisitLockNode(val)
	case *SelectStatementNode:
		return v.VisitSelectStatementNode(val)
	default:
		return v.ToSqlVisitor.Visit(a)
	}
}

// VisitLockNode is overwritten for the SQLiteVisitor
// Locks are not supported in SQLite
func (v SQLiteVisitor) VisitLockNode(node *LockNode) string {
	return ""
}

func (v SQLiteVisitor) VisitSelectStatementNode(node *SelectStatementNode) string {
	if node.Offset != nil && node.Limit == nil {
		node.Limit = &LimitNode{Expr: Sql("-1")}
	}
	return v.ToSqlVisitor.VisitSelectStatementNode(node)
}
