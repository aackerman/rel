package rel

type SelectManager struct {
	Engine Engine
	Ast    SelectStatementNode
	Ctx    *SelectCoreNode
	BaseVisitable
}

func NewSelectManager(e Engine, t *Table) *SelectManager {
	stmt := SelectStatementNode{Cores: make([]*SelectCoreNode, 0)}
	core := NewSelectCoreNode()
	stmt.Cores = append(stmt.Cores, &core)
	ctx := stmt.Cores[len(stmt.Cores)-1]
	manager := SelectManager{
		Engine: e,
		Ast:    stmt,
		Ctx:    ctx,
	}
	// setup initial join source
	manager.From(t)
	return &manager
}

func (s *SelectManager) ToSql() string {
	return s.Engine.Visitor().Accept(s.Ast)
}

func (s *SelectManager) Project(projections ...Visitable) *SelectManager {
	for _, projection := range projections {
		if s.Ctx.Projections == nil {
			nodeslice := make([]Visitable, 0)
			s.Ctx.Projections = &nodeslice
		}

		if s.Ctx.Projections != nil {
			*s.Ctx.Projections = append(*s.Ctx.Projections, projection)
		}
	}
	return s
}

func (s *SelectManager) From(t *Table) *SelectManager {
	if t != nil {
		var v Visitable = t
		s.Ctx.Source.Left = v
	}
	return s
}

func (s *SelectManager) On(a ...Visitable) *SelectManager {
	right := s.Ctx.Source.Right

	if len(right) > 0 {
		last := right[len(right)-1]
		if l, ok := last.(JoinSource); ok {
			l.Right = append(l.Right, NewOnNode(s.collapse(a...)))
		}
	}

	return s
}

func (s *SelectManager) Join(right Visitable) *SelectManager {
	return s.InnerJoin(right)
}

func (s *SelectManager) InnerJoin(join Visitable) *SelectManager {
	s.Ctx.Source.Right = append(s.Ctx.Source.Right, InnerJoinNode{Left: join})
	return s
}

func (s *SelectManager) OuterJoin(right Visitable) *SelectManager {
	return s
}

func (s *SelectManager) StringJoin(right Visitable) *SelectManager {
	return s
}

func (s *SelectManager) Take(i int) *SelectManager {
	limit := NewLimitNode(Sql(i))
	s.Ast.Limit = &limit
	return s
}

func (s *SelectManager) Exists() ExistsNode {
	return NewExistsNode(s.Ast)
}

func (s *SelectManager) Order(exprs ...string) *SelectManager {
	if len(exprs) > 0 {
		if s.Ast.Orders == nil {
			orders := make([]Visitable, 0)
			s.Ast.Orders = &orders
		}
		for _, expr := range exprs {
			order := NewSqlLiteralNode(expr)
			*s.Ast.Orders = append(*s.Ast.Orders, order)
		}
	}
	return s
}

func (s *SelectManager) Where(n Visitable) *SelectManager {
	if s.Ctx.Wheres == nil {
		wheres := make([]Visitable, 0)
		s.Ctx.Wheres = &wheres
	}

	if expr, ok := n.(SelectManager); ok {
		*s.Ctx.Wheres = append(*s.Ctx.Wheres, expr.Ast)
	} else {
		*s.Ctx.Wheres = append(*s.Ctx.Wheres, n)
	}

	return s
}

func (s *SelectManager) Group(columns ...Visitable) *SelectManager {
	var group GroupNode
	if len(columns) > 0 {
		if s.Ctx.Groups == nil {
			groups := make([]GroupNode, 0)
			s.Ctx.Groups = &groups
		}
		for _, column := range columns {
			group = NewGroupNode(column)
			*s.Ctx.Groups = append(*s.Ctx.Groups, group)
		}
	}
	return s
}

func (s *SelectManager) Intersect(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(s.Engine).Intersect(stmt1, stmt2)
}

func (s *SelectManager) Union(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(s.Engine).Union(stmt1, stmt2)
}

func (s *SelectManager) UnionAll(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(s.Engine).UnionAll(stmt1, stmt2)
}

func (s *SelectManager) Except(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(s.Engine).Except(stmt1, stmt2)
}

func (s *SelectManager) Skip(i int) *SelectManager {
	offset := NewOffsetNode(Sql(i))
	s.Ast.Offset = &offset
	return s
}

func (s *SelectManager) Offset(i int) *SelectManager {
	return s.Skip(i)
}

func (s *SelectManager) Having(a ...Visitable) *SelectManager {
	having := NewHavingNode(s.collapse(a...))
	s.Ctx.Having = &having
	return s
}

func (s *SelectManager) collapse(a ...Visitable) Visitable {
	var v Visitable

	// use the first Node if there is only one
	// else create and And node
	if len(a) == 1 {
		v = a[0]
	} else {
		v = s.NewAndNode(a...)
	}
	return v
}
