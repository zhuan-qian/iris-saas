package common

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"os"
	"strings"
	"sync"
	"xorm.io/core"
	"xorm.io/xorm"
)

var (
	onceDb         sync.Once
	instanceDb     *xorm.Engine
	onceCasbin     sync.Once
	instanceCasbin *casbin.Enforcer
)

//自定义xorm的名称规则mapper
type CamelMapper struct {
}

func (m CamelMapper) Obj2Table(o string) string {
	return m.convert(o)
}

func (m CamelMapper) Table2Obj(t string) string {
	return m.convert(t)
}

func (m CamelMapper) convert(a string) string {
	first := strings.ToLower(string(a[0]))
	if len(a) > 1 {
		a = first + a[1:]
	}
	return a
}

func GetCasbin() *casbin.Enforcer {
	onceCasbin.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		db_config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
		adapter, err := xormadapter.NewAdapter("mysql", db_config, true)
		if err != nil {
			panic(err)
		}
		instanceCasbin, err = casbin.NewEnforcer("system/rbac_model.conf", adapter)
		if err != nil {
			panic(err)
		}
		err = instanceCasbin.LoadPolicy()
		if err != nil {
			panic(err)
		}
		instanceCasbin.AddFunction("printer", Printer)
		instanceCasbin.AddFunction("isAdminInDomain", IsAdminInDomain)
	})
	return instanceCasbin
}

func GetDB() *xorm.Engine {
	onceDb.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		charset := os.Getenv("DB_CHARSET")
		db_config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, host, port, database, charset)
		db, err := xorm.NewEngine("mysql", db_config)
		if err != nil {
			panic("数据库连接失败")
		}
		db.SetMapper(CamelMapper{})
		db.SetTableMapper(core.SnakeMapper{})
		instanceDb = db
	})
	return instanceDb
}

func Transaction() (session *xorm.Session, err error) {
	session = GetDB().NewSession()
	err = session.Begin()
	return
}
