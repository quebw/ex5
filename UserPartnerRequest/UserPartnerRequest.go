package userpartnerrequest

import (
	"encoding/json"
	"gRPC/db"
	_ "gRPC/ex5"
	__ "gRPC/ex5"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

func conn() (engine *xorm.Engine) {
	engine, err := db.Connect()
	if err == nil {
		log.Println(err)
		return
	}
	return engine
}

func UserPartnerRequest(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	var users []__.UserPartner
	message := __.UserPartnerRequest{
		UserId: "1",
		Phone:  "1234",
		Limit:  3,
	}
	count, err := engine.Where("user_id = ?", message.UserId).Or("phone = ?", message.Phone).Limit(int(message.Limit), 0).FindAndCount(&users)
	if err != nil {
		log.Println(w, "Error :", err)
	}
	for i := 0; i < int(count); i++ {
		log.Println(w, "ID: ", users[i].UserId,"Phone: ", users[i].Phone)
	}
}

func GetAllUser(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	users, count := db.ListUser(engine, []__.UserPartner{})
	for i := 0; i < int(count); i++ {
		log.Println(w,"id: %v, User_id: %v, Phone: %v\n",users[i].Id, users[i].UserId, users[i].Phone)
	}

}


func GetUserById(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	engine := conn()
	us, err := db.ReadUser(engine,id)
	if err != nil {
		log.Println(w, err)
	}
	if id == us.UserId {
		log.Println(w,"User: %v, phone: %v",us.UserId, us.Phone)
		return
	}
	log.Println(w,"User not found")
}


func CreateUser(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	reqBody,_ := ioutil.ReadAll(request.Body)
	user1 := __.UserPartner{
		Id:          xid.New().String(),
		Created:     time.Now().UnixMilli(),
		UpdatedAt:    time.Now().UnixMilli(),
	}
	json.Unmarshal(reqBody, &user1)
	err := db.InsertTable(engine,&user1)
	if err != nil {
		log.Println(w,"Error: ",err)
		return
	}
	log.Println(w,"Success")
}


func UpdateUser(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	engine := conn()
	user_id := params["id"]
	if user_id == "" {
		log.Println(w,"User not found")
		return
	}
	reqBody, _ :=ioutil.ReadAll(request.Body)
	time := __.UserPartner{
		UpdatedAt:    time.Now().UnixMilli(),
	}
	json.Unmarshal(reqBody,&time)
	err := db.UpdateUser(engine,user_id, time)
	if err != nil {
		log.Println(w,"Error: ",err)
		return
	}
	log.Println(w,"Success")
}


func DeleteUser(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	user_id := params["id"]
	engine := conn()
	err := db.DeleteUser(engine,user_id)
	if err != nil {
		log.Println(w,"Error:",err)
		return
	}
	log.Println(w,"Success")
}