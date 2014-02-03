# Rel

A SQL AST manager for Go.

## Usage

```go
package main

import (
  "fmt"
  "rel"
)

func main() {
  users := rel.NewTable("users")
  manager := users.Select(rel.Sql("*"))
  fmt.Println(manager.ToSql()) // SELECT * FROM "users"
}
```

## Where

```go
users.Where(users.Attr("name").Eq(rel.Sql("amy")))
// SELECT * FROM "users" WHERE "users"."name" = "amy"
```

## Updates

```go
users := rel.NewTable("users")
update := rel.NewUpdateManager(RelEngine)
update.Table(users).Set(users.Attr("name"), rel.Sql("amy"))
fmt.Println(manager.ToSql()) // UPDATE "users" SET "name" = amy
```

## Deletes

```go
users := rel.NewTable("users")
delete := rel.NewDeleteManager(RelEngine)
delete.From(users).Where(users.Attr("id").Eq(rel.Sql(1)))
fmt.Println(delete.ToSql()) // DELETE FROM "users" WHERE "id" = 1
```

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
