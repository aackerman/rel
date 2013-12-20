package arel

type SelectManager struct {
	Ast SelectStatement
	ctx SelectCore
	BaseNode
	TreeManager
}

func NewSelectManager(t *Table) SelectManager {
	stmt := SelectStatement{
		cores:  make([]SelectCore, 10),
		Orders: make([]Order, 10),
	}

	stmt.cores = append(stmt.cores, SelectCore{})

	ctx := stmt.cores[len(stmt.cores)-1]
	return SelectManager{
		stmt,
		ctx,
		CreateBaseNode(),
		TreeManager{
			engine: &t.Engine,
		},
	}
}

func (s *SelectManager) Project(projections ...interface{}) *SelectManager {
	var projection SqlLiteralNode
	for _, p := range projections {
		// For convenience we accept strings and convert them to sql literals
		switch p.(type) {
		case string:
			projection = Sql(p.(string))
		default:
			projection = Sql("*")
		}
		s.ctx.Projections = append(s.ctx.Projections, projection)
	}
	return s
}

func (s *SelectManager) From(t *Table) *SelectManager {
	s.ctx.Source.Left = t
	return s
}

func (s *SelectManager) Projections() []AstNode {
	return s.ctx.Projections
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
	return s
}

func (s *SelectManager) Having(things ...interface{}) *SelectManager {
	return s
}
