package grel

type SelectManager struct {
	Engine Engine
	Ast    SelectStatementNode
	Ctx    *SelectCoreNode
	BaseNode
}

func NewSelectManager(t *Table) SelectManager {
	stmt := SelectStatementNode{Cores: make([]*SelectCoreNode, 0)}
	core := NewSelectCoreNode()
	stmt.Cores = append(stmt.Cores, &core)
	ctx := stmt.Cores[len(stmt.Cores)-1]
	manager := SelectManager{
		Engine: t.Engine,
		Ast:    stmt,
		Ctx:    ctx,
	}
	manager.From(t)
	return manager
}

func (s *SelectManager) ToSql() string {
	return s.Engine.Visitor().Accept(s.Ast)
}

func (s *SelectManager) Project(projections ...AstNode) *SelectManager {
	for _, projection := range projections {
		if s.Ctx.Projections == nil {
			nodeslice := make([]AstNode, 0)
			s.Ctx.Projections = &nodeslice
		}

		if s.Ctx.Projections != nil {
			*s.Ctx.Projections = append(*s.Ctx.Projections, projection)
		}
	}
	return s
}

func (s *SelectManager) From(t *Table) *SelectManager {
	s.Ctx.Source.Left = t
	return s
}

func (s *SelectManager) Join(a ...AstNode) *SelectManager {
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
			orders := make([]AstNode, 0)
			s.Ast.Orders = &orders
		}
		for _, expr := range exprs {
			order := NewSqlLiteralNode(expr)
			*s.Ast.Orders = append(*s.Ast.Orders, order)
		}
	}
	return s
}

func (s *SelectManager) Where(n AstNode) *SelectManager {
	if s.Ctx.Wheres == nil {
		wheres := make([]AstNode, 0)
		s.Ctx.Wheres = &wheres
	}
	*s.Ctx.Wheres = append(*s.Ctx.Wheres, n)
	return s
}

func (s *SelectManager) Group(columns ...AstNode) *SelectManager {
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

func (s *SelectManager) Skip(i int) *SelectManager {
	offset := NewOffsetNode(Sql(i))
	s.Ast.Offset = &offset
	return s
}

func (s *SelectManager) Offset(i int) *SelectManager {
	return s.Skip(i)
}

func (s *SelectManager) Having(a ...AstNode) *SelectManager {
	var b AstNode

	// use the first Node if there is only one
	// else create and And node
	if len(a) == 1 {
		b = a[0]
	} else {
		b = s.NewAndNode(a...)
	}

	// pass in an AstNode
	having := NewHavingNode(b)
	s.Ctx.Having = &having
	return s
}
