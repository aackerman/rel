package arel

type SelectManager struct {
	Engine Engine
	Ast    SelectStatement
	Ctx    *SelectCoreNode
	BaseNode
}

func NewSelectManager(t *Table) SelectManager {
	stmt := SelectStatement{Cores: make([]*SelectCoreNode, 0)}
	stmt.Cores = append(stmt.Cores, CreateSelectCoreNode(t))
	ctx := stmt.Cores[len(stmt.Cores)-1]
	manager := SelectManager{
		Engine:   t.Engine,
		Ast:      stmt,
		Ctx:      ctx,
		BaseNode: CreateBaseNode(),
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

func (s *SelectManager) Join(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Take(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Order(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Where(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Group(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Skip(things ...interface{}) *SelectManager {
	s.Ast.Offset = &OffsetNode{

	}
	return s
}

func (s *SelectManager) Offset(things ...interface{}) *SelectManager {
	return s.Skip(things...)
}

func (s *SelectManager) Having(things ...interface{}) *SelectManager {
	return s
}
