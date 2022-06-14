package main

import (
	"log"
	"net/http"
    "gRPC/UserPartnerRequest"
    "github.com/gorilla/mux"
)


func Handle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Error:",http.StatusNotFound)
        return
    }
    log.Println(w, "Pong")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/userpartner", userpartnerrequest.UserPartnerRequest)

    //ex3
    r.HandleFunc("/", Handle)

    //ex4
    r.HandleFunc("/user-partner", userpartnerrequest.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user-partner/{id}", userpartnerrequest.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/user-partner/create", userpartnerrequest.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user-partner/update/{id}", userpartnerrequest.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user-partner/delete/{id}", userpartnerrequest.DeleteUser).Methods(http.MethodDelete)

    log.Println(http.ListenAndServe(":3001", r))
}

