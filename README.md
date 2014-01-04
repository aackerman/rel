# Rel

A SQL AST manager for Go. It:

* Simplifies the generation of complex SQL queries

It is intended to be a framework, meaning this library could fulfill the query generation for a great ORM.

Rel allows you to programatically create dynamic SQL with a clear and simple interface.

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
* `UnionManager`

## Usage

```go
t := Table{
  Name: "users",
  Engine: DefaultEngine
}
sm := t.Project("*")
fmt.Println(sm.ToSql()) // SELECT * FROM "users"
```
