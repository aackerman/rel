package arel

import (
	"strconv"
)

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
		Engine:   t.Engine,
		Ast:      stmt,
		Ctx:      ctx,
		BaseNode: NewBaseNode(),
	}
	manager.From(t)
	return manager
}

func (s *SelectManager) ToSql() string {
	return s.Engine.Visitor().Accept(s.Ast)
}

func (s *SelectManager) Project(projections ...interface{}) *SelectManager {
	var projection SqlLiteralNode
	for _, p := range projections {
		switch p.(type) {
		case string:
			projection = Sql(p.(string))
		default:
			projection = Sql("*")
		}

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

func (s *SelectManager) Take(a ...AstNode) *SelectManager {
	return s
}

func (s *SelectManager) Order(a ...AstNode) *SelectManager {
	return s
}

func (s *SelectManager) Where(a ...AstNode) *SelectManager {
	return s
}

func (s *SelectManager) Group(a ...AstNode) *SelectManager {
	return s
}

func (s *SelectManager) Skip(i int) *SelectManager {
	offset := NewOffsetNode(Sql(strconv.Itoa(i)))
	s.Ast.Offset = &offset
	return s
}

func (s *SelectManager) Offset(i int) *SelectManager {
	return s.Skip(i)
}

func (s *SelectManager) Having(a ...AstNode) *SelectManager {
	return s
}
