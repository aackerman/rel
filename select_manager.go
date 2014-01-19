package rel

import (
	"log"
)

type SelectManager struct {
	Engine Engine
	Ast    SelectStatementNode
	Ctx    *SelectCoreNode
	BaseVisitable
}

func NewSelectManager(engine Engine, table *Table) *SelectManager {
	if engine == nil {
		log.Fatal("Please register an engine before proceding")
	}
	stmt := NewSelectStatementNode()
	manager := SelectManager{
		Engine: engine,
		Ast:    stmt,
		Ctx:    stmt.Cores[len(stmt.Cores)-1],
	}
	// setup initial join source
	manager.From(table)
	return &manager
}

func (mgr *SelectManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func (mgr *SelectManager) Project(projections ...Visitable) *SelectManager {
	for _, projection := range projections {
		if mgr.Ctx.Projections == nil {
			mgr.Ctx.Projections = &[]Visitable{}
		}

		*mgr.Ctx.Projections = append(*mgr.Ctx.Projections, projection)
	}
	return mgr
}

func (mgr *SelectManager) From(table *Table) *SelectManager {
	if table != nil {
		var visitable Visitable = table
		mgr.Ctx.Source.Left = visitable
	}
	return mgr
}

func (mgr *SelectManager) As(node SqlLiteralNode) *TableAliasNode {
	return &TableAliasNode{
		Relation: &GroupingNode{Expr: []Visitable{mgr.Ast}},
		Name:     node,
	}
}

func (mgr *SelectManager) On(visitables ...Visitable) *SelectManager {
	right := mgr.Ctx.Source.Right

	if len(right) > 0 {
		last := right[len(right)-1]
		switch val := last.(type) {
		case *InnerJoinNode:
			val.Right = mgr.NewOnNode(mgr.collapse(visitables...))
		case *OuterJoinNode:
			val.Right = mgr.NewOnNode(mgr.collapse(visitables...))
		default:
			log.Fatalf("Unable to call On with input type %T", val)
		}
	}

	return mgr
}

func (mgr *SelectManager) Join(right Visitable) *SelectManager {
	return mgr.InnerJoin(right)
}

func (mgr *SelectManager) InnerJoin(visitable Visitable) *SelectManager {
	mgr.Ctx.Source.Right = append(mgr.Ctx.Source.Right, &InnerJoinNode{Left: visitable})
	return mgr
}

func (mgr *SelectManager) OuterJoin(visitable Visitable) *SelectManager {
	mgr.Ctx.Source.Right = append(mgr.Ctx.Source.Right, &OuterJoinNode{Left: visitable})
	return mgr
}

// FIXME: Allow for other types of locks
func (mgr *SelectManager) Lock() *SelectManager {
	mgr.LockForUpdate()
	return mgr
}

func (mgr *SelectManager) LockForUpdate() *SelectManager {
	mgr.Ast.Lock = NewLockNode(Sql("FOR UPDATE"))
	return mgr
}

func (mgr *SelectManager) Take(i int) *SelectManager {
	mgr.Ast.Limit = NewLimitNode(Sql(i))
	return mgr
}

func (mgr *SelectManager) Exists() *ExistsNode {
	return NewExistsNode(mgr.Ast)
}

func (mgr *SelectManager) Order(expressions ...Visitable) *SelectManager {
	if len(expressions) > 0 {
		if mgr.Ast.Orders == nil {
			mgr.Ast.Orders = &[]Visitable{}
		}
		for _, expression := range expressions {
			*mgr.Ast.Orders = append(*mgr.Ast.Orders, expression)
		}
	}
	return mgr
}

func (mgr *SelectManager) Where(n Visitable) *SelectManager {
	if mgr.Ctx.Wheres == nil {
		mgr.Ctx.Wheres = &[]Visitable{}
	}

	if expr, ok := n.(SelectManager); ok {
		*mgr.Ctx.Wheres = append(*mgr.Ctx.Wheres, expr.Ast)
	} else {
		*mgr.Ctx.Wheres = append(*mgr.Ctx.Wheres, n)
	}

	return mgr
}

func (mgr *SelectManager) Group(columns ...Visitable) *SelectManager {
	if len(columns) > 0 {
		if mgr.Ctx.Groups == nil {
			mgr.Ctx.Groups = &[]GroupNode{}
		}
		for _, column := range columns {
			*mgr.Ctx.Groups = append(*mgr.Ctx.Groups, NewGroupNode(column))
		}
	}
	return mgr
}

func (mgr *SelectManager) Intersect(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Intersect(stmt1, stmt2)
}

func (mgr *SelectManager) Union(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Union(stmt1, stmt2)
}

func (mgr *SelectManager) UnionAll(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).UnionAll(stmt1, stmt2)
}

func (mgr *SelectManager) Except(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Except(stmt1, stmt2)
}

func (mgr *SelectManager) Skip(i int) *SelectManager {
	mgr.Ast.Offset = NewOffsetNode(Sql(i))
	return mgr
}

func (mgr *SelectManager) Offset(i int) *SelectManager {
	return mgr.Skip(i)
}

func (mgr *SelectManager) Having(visitables ...Visitable) *SelectManager {
	mgr.Ctx.Having = NewHavingNode(mgr.collapse(visitables...))
	return mgr
}

func (mgr *SelectManager) Distinct() *SelectManager {
	mgr.Ctx.SetQuanifier = &DistinctNode{}
	return mgr
}

func (mgr *SelectManager) NotDistinct() *SelectManager {
	mgr.Ctx.SetQuanifier = nil
	return mgr
}

func (mgr *SelectManager) With(node Visitable) *SelectManager {
	mgr.Ast.With = &WithNode{Expr: node}
	return mgr
}

func (mgr *SelectManager) WithRecursive(node Visitable) *SelectManager {
	mgr.Ast.With = &WithRecursiveNode{Expr: node}
	return mgr
}

func (mgr *SelectManager) Window(node SqlLiteralNode) *NamedWindowNode {
	if mgr.Ctx.Windows == nil {
		mgr.Ctx.Windows = &[]Visitable{}
	}
	window := &NamedWindowNode{Name: node}
	*mgr.Ctx.Windows = append(*mgr.Ctx.Windows, window)
	return window
}

func (mgr *SelectManager) collapse(visitables ...Visitable) Visitable {
	var v Visitable

	// use the first Node if there is only one
	// else create and And node
	if len(visitables) == 1 {
		v = visitables[0]
	} else {
		v = mgr.NewAndNode(visitables...)
	}
	return v
}
