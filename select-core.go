package arel

type SelectCore struct {
  Source      JoinSource
  Top
  Projections []Node,
  Wheres      []Node,
  Groups      []GroupingNode,
  Having      interface{},
  Windows     []Node,
}

func (s SelectCore) NodeInterface() {}
