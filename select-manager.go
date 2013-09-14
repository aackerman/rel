package arel

type SelectManager struct {
	TreeManager
}

func NewSelectManager(e *Engine, t *Table) *SelectManager {
	return &SelectManager{
		TreeManager{
			Ast:    NewSelectStatementNode(),
			Engine: e,
		},
	}
}

func (s *SelectManager) Project(things ...interface{}) *SelectManager {
	return s
}

func (s *SelectManager) Select(things ...interface{}) *SelectManager {
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
	return s
}

func (s *SelectManager) Having(things ...interface{}) *SelectManager {
	return s
}
