# Rel [![Build Status](https://travis-ci.org/aackerman/rel.png?branch=master)](https://travis-ci.org/aackerman/rel)

A SQL AST manager for Go.

## Usage

```go
package main

import (
  "fmt"
  "rel"
)

func main() {
  sql := rel.Select(rel.Star()).From("users").ToSql()
  fmt.Println(sql) // SELECT * FROM "users"
}
```

## Where

```go
users := rel.NewTable("users")
users.Where(users.Attr("name").Eq(rel.Sql("amy")))
// SELECT * FROM "users" WHERE "users"."name" = "amy"
```

## Joins

```go
users := rel.NewTable("users")
preferences := rel.NewTable("preferences")
manager := rel.Select(rel.Star()).From(users).Join(preferences).On(preferences.Attr("user_id").Eq(users.Attr("user_id")))
fmt.Println(manager.ToSql()) // SELECT * FROM "users" INNER JOIN "preferences" ON "preferences"."user_id" = "users"."user_id"
```

## Updates

```go
users := rel.NewTable("users")
update := rel.NewUpdateManager(rel.RelEngine)
update.Table(users).Set(users.Attr("name"), rel.Sql("amy"))
fmt.Println(update.ToSql()) // UPDATE "users" SET "name" = amy
```

## Deletes

```go
users := rel.NewTable("users")
delete := rel.NewDeleteManager(rel.RelEngine)
delete.From(users).Where(users.Attr("id").Eq(rel.Sql(1)))
fmt.Println(delete.ToSql()) // DELETE FROM "users" WHERE "id" = 1
```

## Inserts

```go
users := rel.NewTable("users")
insert := rel.Insert().Into(users).Values(users.Attr("email"), Sql("a@b.com"))
fmt.Println(insert.ToSql()) // INSERT INTO "users" ("email") VALUES ('a@b.com')
```

## Orders

```go
users := rel.NewTable("users")
manager := users.Select(rel.Star()).Order(users.Attr("first_name"))
fmt.Println(manager.ToSql()) // SELECT * FROM "users" ORDER BY "users"."first_name"
```

### With Direction

```go
users := rel.NewTable("users")
manager := users.Select(rel.Star()).Order(users.Attr("first_name").Desc())
fmt.Println(manager.ToSql()) // SELECT * FROM "users" ORDER BY "users"."first_name" DESC
```

## Group By

```go
users := rel.NewTable("users")
manager := users.Select(rel.Star()).GroupBy(users.Attr("first_name"))
fmt.Println(manager.ToSql()) // SELECT * FROM "users" GROUP BY "users"."first_name"
```

## Counts

```go
users := rel.NewTable("users")
manager := users.Select(rel.Count())
fmt.Println(manager.ToSql()) // SELECT COUNT(1) FROM "users"
```

## Database Specific SQL

Nearly every RDBMS has it's own quirks and non-standard features. For the most general cases we use the `ToSqlVisitor` to handle compiling the AST to a SQL statement. It's likely that consumers will want to be more specific, for example using PostgreSQL, MySQL, or SQLite.

```go
package main

import (
  "fmt"
  "rel"
)

func main() {
  rel.RegisterDatabase("postgresql")
  users := rel.NewTable("users")
  manager := users.Select(rel.Sql("*"))
  fmt.Println(manager.ToSql()) // SELECT * FROM "users"
}
```

`rel.RegisterDatabase` is a shorthand to allow easy use of built in functionality for PostgreSQL, MySQL, or SQLite.

## Method Interfaces

Several methods in Rel only allow values that satisfy the `Visitable` interface. Rel methods will generally return `Visitable` values. In some cases methods will allow primitive types as method inputs when the type of input is predictable.

```go
users := rel.NewTable("users")
users.Having(users.Attr("id").Eq(rel.Sql(10)))
// SELECT FROM "users" HAVING "users"."id" = 10
```

Breaking down the code here, `users` is a `Table` type. `Table#Having` allows a variadic number of values that satisfy the `Visitable` interface. `users.Attr("id")` returns a pointer to an `AttributeNode` and only accepts a `string`. SQL table fields/attributes can be expressed in terms of strings so an input satifying the `Visitable` interface isn't required because it's only ever necessary to use a string. `AttributeNode#Eq` allows a single `Visitable` type as an input. We use the `Sql` method to wrap an `int` value in a `SqlLiteralNode` to satisfy the `Visitable` interface requirement of `AttributeNode#Eq`.

The type of input for `AttributeNode#Eq` is somewhat predictable ahead of time. In some cases a user may want to use an `int`, `string`, or another `AttributeNode`. That means using an interface. When an input type is unpredictable, Rel uses the `Visitable` type for input as opposed to an empty interface, and offers the `Sql` method as a way to convert primitive types to a value that will satisfy the `Visitable` interface.

## Author

| [![twitter/_aaronackerman_](http://gravatar.com/avatar/c73ff9c7e654647b2b339d9e08b52143?s=70)](http://twitter.com/_aaronackerman_ "Follow @_aaronackerman_ on Twitter") |
|---|
| [Aaron Ackerman](https://twitter.com/_aaronackerman_) |

## License

[MIT](https://github.com/aackerman/rel/blob/master/LICENSE.md)
