package main

import (
	"net/http"

	"github.com/sepidario/yadetam/yadetam"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sepidario/yadetam/db"
	"github.com/sepidario/yadetam/web"
)

func main() {
	conn, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		// TODO: Log the err
		return
	}
	defer conn.Close()

	conn.AutoMigrate(&yadetam.User{})

	us := db.NewUserService(conn)

	uh := web.NewUserHandler(us)

	r := mux.NewRouter()
	r.HandleFunc("/register", uh.Register).Methods("POST")

	if err := http.ListenAndServe(":8000", r); err != nil {
		// TODO: Log the err
		return
	}
}
