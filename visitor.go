package arel

type Visitor interface {
	Accept(Visitor) string
	Visit(Visitor) string
}
