package arel

type ToSqlVisitor struct {
}

func NewToSqlVisitor(c Connection) *ToSqlVisitor {
	return &ToSqlVisitor{}
}
