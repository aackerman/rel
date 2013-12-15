package arel

type SelectCore struct {
	Source       JoinSource
	Top          int
	Projections  []AstNode
	SetQuanifier interface{}
	Wheres       []AstNode
	Groups       []GroupingNode
	Having       interface{}
	Windows      []AstNode
	AstNode
}
