package rel

type IntersectManager struct {
	Engine Engine
	Ast    Visitable
}

func NewIntersectManager(e Engine) *IntersectManager {
	return &IntersectManager{Engine: e}
}
