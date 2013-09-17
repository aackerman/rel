package arel

import (
	"log"
)

type SelectManager struct {
	TreeManager
}

func NewSelectManager(e *Engine, t *Table) *SelectManager {
	selectstmt := SelectStatement{
		Cores:  make([]SelectCore, 10),
		Limit:  0,
		Orders: make([]Order, 10),
	}
	context := selectstmt.Cores[len(selectstmt.Cores)-1]
	return &SelectManager{
		TreeManager{
			Ast:    selectstmt,
			ctx:    context,
			Engine: e,
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
