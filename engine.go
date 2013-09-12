package arel

type Engine interface {
	QuoteTableName()
	QuoteColumnName()
	Quote()
	Columns()
	TableExists(string) bool
	Tables()
	Visitor()
	Execute(string) []string
}

type Connection interface {
	PrimaryKey(string) string
	TableExists(string) bool
	Columns(string) []string
	QuoteTableName(string) string
	QuoteColumnName(string) string
	Quote()
}

type ConnectionPool interface {
	TableExists(string) bool
}
