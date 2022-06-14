package db

import (
	_ "gRPC/ex5"
	__ "gRPC/ex5"
	"log"

	"github.com/go-xorm/xorm"
)

func CreateTable(engine *xorm.Engine, tb interface{}) error {
	_, err := engine.IsTableExist(tb)
	if err != nil {
		log.Println(err)
		return err
	}
	err = engine.Sync2(tb)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ReadUser(engine *xorm.Engine, id string) (*__.UserPartner, error) {
	tb := __.UserPartner{}
	_, err := engine.Where("user_id = ?",id).Get(&tb)
	if err != nil {
		log.Println(err)
	}
	return &tb,nil
}

func InsertTable(engine *xorm.Engine, data interface{}) error {
	_, err := engine.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func ListUser(engine *xorm.Engine, tb []__.UserPartner) ([]__.UserPartner, int64) {
	count,_ := engine.Count(__.UserPartner{})
	err := engine.Limit(int(count),0).Find(&tb)
	if err != nil {
		log.Println(err)	
	}
	return tb ,count
}

func DeleteUser(engine *xorm.Engine, id string) error {
	_, err := engine.Where("user_id = ?", id).Delete(__.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

func Delete(engine *xorm.Engine) error {
	id := ""
	_, err := engine.Where("user_id = ?", id).Delete(__.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(engine *xorm.Engine, id string, data interface{}) error {
	_, err := engine.Where("user_id = ?",id).Update(data)
	if err != nil {
		return err
	}
	return nil
}