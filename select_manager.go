package arel

import (
	"log"
)

type SelectManager struct {
	Ast Node
	ctx Node
	TreeManager
}

func NewSelectManager(e *Engine, t *Table) *SelectManager {
	stmt := SelectStatement{
		cores:  make([]*SelectCore, 10),
		Orders: make([]*Order, 10),
	}

	ctx := stmt.cores[len(stmt.cores)-1]
	return &SelectManager{
		TreeManager: TreeManager{
			Ast:    stmt,
			ctx:    ctx,
			engine: e,
		},
	}
}

// Append to internally held projections
func (s *SelectManager) Project(projections ...interface{}) *SelectManager {
	for _, p := range projections {
		// For convenience we accept strings and convert them to sql literals
		switch p.(type) {
		case string:
			p = Sql(p.(string))
		case SqlLiteralNode:
		default:
			log.Fatal("Can't accept this projection type")
		}
		append(s.ctx.Projections, p)
	}
	return s
}

func (s *SelectManager) From(t *Table) *SelectManager {
	s.ctx.Source.Left = t
	return s
}

func (s *SelectManager) Projections() []*SqlLiteralNode {
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
