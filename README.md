# Rel

A SQL AST manager for Go. It:

* Simplifies the generation of complex SQL queries

It is intended to be a framework, meaning this library could fulfill the query generation for a great ORM.

## Interfaces

* `Visitable`
* `Visitor`
* `Engine`

## Important Classes

* `Table`
* `SelectManager`
* `InsertManager`
* `UpdateManager`
* `DeleteManager`

## Usage

```go
t := Table{
  Name: "users",
  Engine: DefaultEngine
}
sm := t.Project("*")
fmt.Println(sm.ToSql()) // SELECT * FROM "users"
```
