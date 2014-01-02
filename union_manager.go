package rel

import (
	"bytes"
)

// A UnionManager managers the handling of UNIONs on SELECT statements
type UnionManager struct {
	managers []*SelectManager
}

func (u *UnionManager) ToSql() string {
	var buf bytes.Buffer
	buf.WriteString("( ")
	for i, manager := range u.managers {
		buf.WriteString(manager.ToSql())
		if i != len(u.managers)-1 {
			buf.WriteString(" UNION ")
		}
	}
	buf.WriteString(" )")
	return buf.String()
}

// Union appends a SelectManager to allow for more unions
func (u *UnionManager) Union(newManagers ...*SelectManager) *UnionManager {
	for _, m := range newManagers {
		u.managers = append(u.managers, m)
	}
	return u
}

func NewUnionManager() *UnionManager {
	return &UnionManager{
		managers: make([]*SelectManager, 0),
	}
}
