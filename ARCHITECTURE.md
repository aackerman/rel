# Rel Architecture

The `Table` is the most central component to a SQL statement. From a `Table` we can create several different types of statements.

```
* SELECT
* UPDATE
* DELETE
* INSERT
* UNION
* UNION ALL
* EXCEPT
* INTERSECT
```

A `Table` is used as a starting point for any query. Neraly all of the `Table` instance methods return a manager type. A `SelectManager`, an `UpdateManager`, `InsertManager`, `DeleteManager`, or a `MultiStatementManager`. Each with a specific purpose to add visitable nodes to the structure of the query.

The first few managers mentioned handle the type of statements they are named for. A `MultistatementManager` can be used to create `UNION`, `UNION ALL`, `EXCEPT`, and `INTERSECT` querys which join two or more querys together to form a larger query.
