package dbmodels

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var dbConnectErr error

func init() {
	engine, dbConnectErr = xorm.NewEngine("mysql", "root:@/p2p?charset=utf8")
	engine.SetMapper(core.SnakeMapper{})
	engine.ShowSQL(true)
}
