package main

import (
	"UTS_1122029/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/roomsGetAll", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/roomsGetDetail", controller.GetDetailRooms).Methods("GET")
	router.HandleFunc("/insertRooms", controller.InsertRoom).Methods("POST")
	router.HandleFunc("/leaveRooms/{id}", controller.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
