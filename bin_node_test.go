package rel

import (
	"testing"
)

func TestBinNodeEq(t *testing.T) {
	bin1 := BinNode{Expr: Sql("zomg")}
	bin2 := BinNode{Expr: Sql("zomg")}
	if !bin1.Eq(bin2) {
		t.Fail()
	}
}

func TestBinNodeMysqlToSql(t *testing.T) {
	viz := MysqlVisitor{ToSqlVisitor{Conn: new(Connection)}}
	bin := &BinNode{Expr: Sql("zomg")}
	sql := viz.Accept(bin)
	if sql != "BINARY zomg" {
		t.Log(sql)
		t.Fail()
	}
}
