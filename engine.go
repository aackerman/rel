package arel

type Engine interface {
	Connection
	Execute(string) []string
}

type Connection interface {
	Visitor()
	Tables()
	PrimaryKey(string) string
	TableExists(string) bool
	Columns(string) []string
	QuoteTableName(string) string
	QuoteColumnName(string) string
	Quote(interface{}) string
}

type ConnectionPool interface {
	TableExists(string) bool
}
