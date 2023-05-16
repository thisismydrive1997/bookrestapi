package main

import (
	"example/restfulapi/api/transport"
	"example/restfulapi/dbutil"
	"example/restfulapi/repository/repoimpl"
	"fmt"
	"net"
	"net/http"

	"gorm.io/gorm"
)

var db gorm.DB

func main() {
	db, err := dbutil.InitDb()
	if err != nil {
		fmt.Println("Failed to InitDb)")
		panic(0)
	}
	bookRepo := repoimpl.NewBookRepo(db)
	httpHandler := transport.NewBookHttpHandler(bookRepo)
	httpListener, err := net.Listen("tcp", ":2000")
	fmt.Println(http.Serve(httpListener, httpHandler))
}
