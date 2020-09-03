package models

import (
	"github.com/astaxie/beego"
	"runtime"
	"xorm.io/xorm"
)

type Adapter struct {
	driverName     string
	dataSourceName string
	engine         *xorm.Engine
}

var adapter *Adapter

func InitAdapter() {
	adapter = NewAdapter("mysql", beego.AppConfig.String("dataSourceName"))
}

func NewAdapter(driverName string, dataSourceName string) *Adapter {
	a := &Adapter{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}

	a.open()
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) open() {
	if err := a.createDatabase(); err != nil {
		panic(err)
	}

	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		panic(err)
	}

	a.engine = engine
	a.createTable()
}

func (a *Adapter) createDatabase() error {
	engin, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}

	defer engin.Close()

	_, err = engin.Exec("CREATE DATABASE IF NOT EXISTS " + a.dataSourceName + " default charset utf8 COLLATE utf8_general_ci")
	return err
}

func (a *Adapter) close() {
	a.engine.Close()
	a.engine = nil
}

func finalizer(a *Adapter) {
	err := a.engine.Close()
	if err != nil {
		panic(err)
	}
}

func (a *Adapter) createTable() {
	//err := a.engine.Sync2(new(tableStruct))
	//if err != nil {
	//	panic(err)
	//}
}
