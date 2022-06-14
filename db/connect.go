package db

import "github.com/go-xorm/xorm"

func Connect() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "localhost@/bt")
	if err != nil {
		return engine, err
	}
	engine.ShowSQL(true)
	return engine, err
}