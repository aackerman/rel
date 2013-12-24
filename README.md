# Project Name

A SQL AST manager for Go. It:

* Simplifies the generation of complex SQL queries
* Adapts to various RDBMSes

It is intended to be a framework framework; that is, you can build your own ORM with it, focusing on data modeling as opposed to database compatibility and query generation.

## Interfaces

* AstNode
* Visitor
* Engine

## Important Classes

* Table
* SelectManager
* InsertManager
* UpdateManager

## Usage

```go
t := Table{
  Name: "users",
  Engine: DefaultEngine
}
sm := t.Project("*")
fmt.Println(sm.ToSql()) // SELECT * FROM "users"
```
