package arel

// Base engine interface for handling database connections
// and operating on the database

type Engine interface {
	Visitor() Visitor
	Connection() *Connection
	QuoteTableName(string) string
	QuoteColumnName(string) string
	TableExists(string) bool
}
