package arel

type SelectManager struct {
	Engine *Engine
	Ast    SelectStatement
	Ctx    SelectCoreNode
	BaseNode
}

func NewSelectManager(t *Table) SelectManager {
	stmt := SelectStatement{
		Cores:  make([]SelectCoreNode, 10),
		Orders: make([]OrderNode, 10),
	}
	stmt.Cores = append(stmt.Cores, CreateSelectCoreNode())
	ctx := stmt.Cores[len(stmt.Cores)-1]
	return SelectManager{
		Engine:   &t.Engine,
		Ast:      stmt,
		Ctx:      ctx,
		BaseNode: CreateBaseNode(),
	}
}

func (s *SelectManager) ToSql() string {
	return s.Engine.Connection().Visitor.Accept(s.Ast)
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
		s.Ctx.Projections = append(s.Ctx.Projections, projection)
	}
	return s
}

// func (s *SelectManager) From(t *Table) *SelectManager {
// 	s.ctx.Source.Left = t
// 	return s
// }

// func (s *SelectManager) Projections() []AstNode {
// 	return s.ctx.Projections
// }

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
	return s
}

func (s *SelectManager) Having(things ...interface{}) *SelectManager {
	return s
}
