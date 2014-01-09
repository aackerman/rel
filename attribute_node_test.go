package rel

import (
	"testing"
)

func TestAttributeNotEqSql(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" != 10"
	if sql != expected {
		t.Logf("TestAttributeNotEqSql sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 OR \"users\".\"id\" != 2)"
	if sql != expected {
		t.Logf("TestAttributeNotEqAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqNil(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEq(nil))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" IS NOT NULL"
	if sql != expected {
		t.Logf("TestAttributeNotEqNil sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeNotEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").NotEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" != 1 AND \"users\".\"id\" != 2)"
	if sql != expected {
		t.Logf("TestAttributeNotEqAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGt(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Gt(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" > 10"
	if sql != expected {
		t.Logf("TestAttributeGt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEq(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" >= 10"
	if sql != expected {
		t.Logf("TestAttributeGtEq sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 OR \"users\".\"id\" >= 2)"
	if sql != expected {
		t.Logf("TestAttributeGtEqAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" >= 1 AND \"users\".\"id\" >= 2)"
	if sql != expected {
		t.Logf("TestAttributeGtEqAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 AND \"users\".\"id\" > 2)"
	if sql != expected {
		t.Logf("TestAttributeGtAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeGtAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").GtAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" > 1 OR \"users\".\"id\" > 2)"
	if sql != expected {
		t.Logf("TestAttributeGtAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLt(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").Lt(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" < 10"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEq(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEq(Sql(10)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" <= 10"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEqAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEqAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 OR \"users\".\"id\" <= 2)"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtEqAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtEqAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" <= 1 AND \"users\".\"id\" <= 2)"
	if sql != expected {
		t.Logf("TestAttributeLt sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtAny(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtAny(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 OR \"users\".\"id\" < 2)"
	if sql != expected {
		t.Logf("TestAttributeLtAny sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeLtAll(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id"))
	mgr.Where(users.Attr("id").LtAll(Sql(1), Sql(2)))
	sql := mgr.ToSql()
	expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE (\"users\".\"id\" < 1 AND \"users\".\"id\" < 2)"
	if sql != expected {
		t.Logf("TestAttributeLtAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}

func TestAttributeCount(t *testing.T) {
	users := NewTable("users")
	mgr := users.Select(users.Attr("id").Count())
	sql := mgr.ToSql()
	expected := "SELECT COUNT(\"users\".\"id\") FROM \"users\""
	if sql != expected {
		t.Logf("TestAttributeLtAll sql: \n%s != \n%s", sql, expected)
		t.Fail()
	}
}
