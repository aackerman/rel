package arel

type ToSqlVisitor struct {
}

func ToSqlVisitorNew(c Connection) ToSqlVisitor {
	return ToSqlVisitor{}
}
