package rel

import (
	"fmt"
	"testing"
)

func TestSelectManagerJoinSources(t *testing.T) {
	t.Skip("TestSelectManagerJoinSources not implemented")
}

func TestSelectManagerSkip(t *testing.T) {
	table := NewTable("users")
	manager := table.From(table)
	sql := manager.Skip(10).ToSql()
	if sql != "SELECT FROM \"users\" OFFSET 10" {
		t.Fail()
	}
}

func TestSelectManagerClone(t *testing.T) {
	t.Skip("TestSelectManagerClone not implemented")
}

func TestSelectManagerExists(t *testing.T) {
	table := NewTable("users")
	manager := table.From(table)
	manager.Project(Sql("*"))
	m2 := NewSelectManager(DefaultEngine, nil)
	m2.Project(manager.Exists())
	sql := m2.ToSql()
	expected := fmt.Sprintf("SELECT EXISTS (%s)", manager.ToSql())
	if sql != expected {
		t.Logf("TestSelectManagerExists sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerExistsAs(t *testing.T) {
	table := NewTable("users")
	manager := table.From(table)
	manager.Project(Sql("*"))
	m2 := NewSelectManager(DefaultEngine, nil)
	m2.Project(manager.Exists().As(Sql("foo")))
	sql := m2.ToSql()
	expected := fmt.Sprintf("SELECT EXISTS (%s) AS foo", manager.ToSql())
	if sql != expected {
		t.Logf("TestSelectManagerExists sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOffset(t *testing.T) {
	table := NewTable("users")
	manager := table.From(table)
	sql := manager.Offset(10).ToSql()
	expected := "SELECT FROM \"users\" OFFSET 10"
	if sql != expected {
		t.Fail()
	}
}

func TestSelectManagerUnion(t *testing.T) {
	table := NewTable("users")
	m1 := NewSelectManager(DefaultEngine, table)
	m1.Project(Star())
	m1.Where(table.Attr("age").Lt(Sql(18)))
	m2 := NewSelectManager(DefaultEngine, table)
	m2.Project(Star())
	m2.Where(table.Attr("age").Gt(Sql(99)))
	mgr := m1.Union(m1.Ast, m2.Ast)
	sql := mgr.ToSql()
	expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 UNION SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
	if sql != expected {
		t.Logf("TestSelectManagerUnion sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerUnionAll(t *testing.T) {
	table := NewTable("users")
	m1 := NewSelectManager(DefaultEngine, table)
	m1.Project(Star())
	m1.Where(table.Attr("age").Lt(Sql(18)))
	m2 := NewSelectManager(DefaultEngine, table)
	m2.Project(Star())
	m2.Where(table.Attr("age").Gt(Sql(99)))
	mgr := m1.UnionAll(m1.Ast, m2.Ast)
	sql := mgr.ToSql()
	expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 UNION ALL SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
	if sql != expected {
		t.Logf("TestSelectManagerUnionAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerIntersect(t *testing.T) {
	table := NewTable("users")
	m1 := NewSelectManager(DefaultEngine, table)
	m1.Project(Star())
	m1.Where(table.Attr("age").Lt(Sql(18)))
	m2 := NewSelectManager(DefaultEngine, table)
	m2.Project(Star())
	m2.Where(table.Attr("age").Gt(Sql(99)))
	mgr := m1.Intersect(m1.Ast, m2.Ast)
	sql := mgr.ToSql()
	expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 18 INTERSECT SELECT * FROM \"users\" WHERE \"users\".\"age\" > 99 )"
	if sql != expected {
		t.Logf("TestSelectManagerUnionAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerExcept(t *testing.T) {
	table := NewTable("users")
	m1 := NewSelectManager(DefaultEngine, table)
	m1.Project(Star())
	m1.Where(table.Attr("age").Lt(Sql(99)))
	m2 := NewSelectManager(DefaultEngine, table)
	m2.Project(Star())
	m2.Where(table.Attr("age").Lt(Sql(50)))
	mgr := m1.Except(m1.Ast, m2.Ast)
	sql := mgr.ToSql()
	expected := "( SELECT * FROM \"users\" WHERE \"users\".\"age\" < 99 EXCEPT SELECT * FROM \"users\" WHERE \"users\".\"age\" < 50 )"
	if sql != expected {
		t.Logf("TestSelectManagerExcept sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerJoin(t *testing.T) {
	left := NewTable("users")
	right := left.Alias()
	predicate := left.Attr("id").Eq(right.Attr("id"))
	mgr := left.Select(Star()).Join(right).On(predicate)
	sql := mgr.ToSql()
	expected := "SELECT * FROM \"users\" INNER JOIN \"users\" \"users_2\" ON \"users\".\"id\" = \"users_2\".\"id\""
	if sql != expected {
		t.Logf("TestSelectManagerJoin sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOrder(t *testing.T) {
	table := NewTable("users")
	mgr := table.Select(Star())
	mgr.Order(table.Attr("id"))
	sql := mgr.ToSql()
	expected := "SELECT * FROM \"users\" ORDER BY \"users\".\"id\""
	if sql != expected {
		t.Logf("TestSelectManagerOrder sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOrderWithDirection(t *testing.T) {
	table := NewTable("users")
	mgr := table.Select(Star())
	mgr.Order(table.Attr("id").Desc())
	sql := mgr.ToSql()
	expected := "SELECT * FROM \"users\" ORDER BY \"users\".\"id\" DESC"
	if sql != expected {
		t.Logf("TestSelectManagerOrderWithDirection sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOnWithMultipleArguments(t *testing.T) {
	left := NewTable("users")
	right := left.Alias()
	predicate := left.Attr("id").Eq(right.Attr("id"))
	mgr := left.From(left)
	mgr.Join(right).On(predicate, predicate)
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" INNER JOIN \"users\" \"users_2\" ON \"users\".\"id\" = \"users_2\".\"id\" AND \"users\".\"id\" = \"users_2\".\"id\""
	if sql != expected {
		t.Logf("TestSelectManagerOnWithMultipleArguments sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerOrderWithAttributesForExpressions(t *testing.T) {
	table := NewTable("users")
	mgr := table.Select(Star())
	mgr.Order(table.Attr("id").Count().Desc())
	sql := mgr.ToSql()
	expected := "SELECT * FROM \"users\" ORDER BY COUNT(\"users\".\"id\") DESC"
	if sql != expected {
		t.Logf("TestSelectManagerOrderWithAttributesForExpressions sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerLock(t *testing.T) {
	table := NewTable("users")
	mgr := table.From(table)
	sql := mgr.Lock().ToSql()
	expected := "SELECT FROM \"users\" FOR UPDATE"
	if sql != expected {
		t.Logf("TestSelectManagerLock sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerGroup(t *testing.T) {
	users := NewTable("users")
	mgr := users.Group(users.Attr("id"))
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" GROUP BY \"users\".\"id\""
	if sql != expected {
		t.Logf("TestSelectManagerGroup sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowEmpty(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window"))
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS ()"
	if sql != expected {
		t.Logf("TestSelectManagerWindowEmpty sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithOrders(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Order(users.Attr("foo").Asc())
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ORDER BY \"users\".\"foo\" ASC)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithOrders sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithUnboundedPreceding(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Rows(&PrecedingNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS UNBOUNDED PRECEDING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithUnboundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithBoundedPreceding(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Rows(&PrecedingNode{Expr: Sql(5)})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS 5 PRECEDING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithBoundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithUnboundedFollowing(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Rows(&FollowingNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS UNBOUNDED FOLLOWING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithUnboundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithBoundedFollowing(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Rows(&FollowingNode{Expr: Sql(5)})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS 5 FOLLOWING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithUnboundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowWithCurrentRow(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Window(Sql("a_window")).Rows(&CurrentRowNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS CURRENT ROW)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowWithUnboundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameBetweenTwoDelimeters(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Frame(&BetweenNode{
		Left:  window.Rows(nil),
		Right: mgr.NewAndNode(&PrecedingNode{}, &CurrentRowNode{}),
	})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameBetweenTwoDelimeters sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeUnboundedPreceding(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Range(&PrecedingNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE UNBOUNDED PRECEDING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeUnboundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeBoundedPreceding(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Range(&PrecedingNode{Expr: Sql(5)})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE 5 PRECEDING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeBoundedPreceding sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeUnboundedFollowing(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Range(&FollowingNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE UNBOUNDED FOLLOWING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeUnboundedFollowing sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeBoundedFollowing(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Range(&FollowingNode{Expr: Sql(5)})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE 5 FOLLOWING)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeBoundedFollowing sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeCurrentRow(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Range(&CurrentRowNode{})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE CURRENT ROW)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeBoundedFollowing sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWindowFrameRangeTwoDelimeters(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	window := mgr.Window(Sql("a_window"))
	window.Frame(&BetweenNode{
		Left:  window.Range(nil),
		Right: mgr.NewAndNode(&PrecedingNode{}, &CurrentRowNode{}),
	})
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" WINDOW \"a_window\" AS (RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)"
	if sql != expected {
		t.Logf("TestSelectManagerWindowFrameRangeBoundedFollowing sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerJoinMultipleTables(t *testing.T) {
	users := NewTable("users")
	comments := NewTable("comments")
	counts := comments.From(comments).Group(comments.Attr("user_id")).Project(
		comments.Attr("user_id").As(Sql("user_id")),
		comments.Attr("user_id").Count().As(Sql("count")),
	).As(Sql("counts"))
	mgr := users.Join(counts).On(counts.Attr("user_id").Eq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\" INNER JOIN (SELECT \"comments\".\"user_id\" AS user_id, COUNT(\"comments\".\"user_id\") AS count FROM \"comments\" GROUP BY \"comments\".\"user_id\") counts ON counts.\"user_id\" = 10"
	if sql != expected {
		t.Logf("TestSelectManagerJoinMultipleTables sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerDistinct(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Distinct()
	sql := mgr.ToSql()
	expected := "SELECT DISTINCT FROM \"users\""
	if sql != expected {
		t.Logf("TestSelectManagerDistinct sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerNotDistinct(t *testing.T) {
	users := NewTable("users")
	mgr := users.From(users)
	mgr.Distinct()
	mgr.NotDistinct()
	sql := mgr.ToSql()
	expected := "SELECT FROM \"users\""
	if sql != expected {
		t.Logf("TestSelectManagerNotDistinct sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestSelectManagerWithRecursiveSupport(t *testing.T) {
	comments := NewTable("comments")
	commentsId := comments.Attr("id")
	commentsParentId := comments.Attr("parent_id")

	replies := NewTable("replies")
	repliesId := replies.Attr("id")

	recursiveTerm := NewSelectManager(DefaultEngine, nil)
	recursiveTerm.From(comments).Project(commentsId, commentsParentId).Where(commentsId.Eq(Sql(42)))

	nonRecursiveTerm := NewSelectManager(DefaultEngine, nil)
	nonRecursiveTerm.From(comments).Project(commentsId, commentsParentId).Join(replies).On(commentsParentId.Eq(repliesId))

	union := recursiveTerm.Union(recursiveTerm.Ast, nonRecursiveTerm.Ast)

	asStmt := &AsNode{Left: replies, Right: union}

	mgr := NewSelectManager(DefaultEngine, nil)
	mgr.WithRecursive(asStmt).From(replies).Project(Star())

	sql := mgr.ToSql()
	expected := "WITH RECURSIVE \"replies\" AS ( SELECT \"comments\".\"id\", \"comments\".\"parent_id\" FROM \"comments\" WHERE \"comments\".\"id\" = 42 UNION SELECT \"comments\".\"id\", \"comments\".\"parent_id\" FROM \"comments\" INNER JOIN \"replies\" ON \"comments\".\"parent_id\" = \"replies\".\"id\" ) SELECT * FROM \"replies\""
	if sql != expected {
		t.Logf("TestSelectManagerWithRecursiveSupport sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}
