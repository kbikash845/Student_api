package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bikashkumar/student_api/internal/config"
)


func main(){
	fmt.Println("welcome to student api")
	// load config
cfg:= config.MustLoad()

	// database setup
	// setup router
  router:=http.NewServeMux()
  router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to student api"))
  })
	// setup server

	server:=http.Server{
		Addr:  cfg.HTTPServer.Addr,
		Handler: router,

	}

	fmt.Println("Server started")

	err:=server.ListenAndServe()
	if err != nil{
		log.Fatal("fail to start server")

	}

	fmt.Println("srver started")
}
