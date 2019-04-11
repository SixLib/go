package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	router "github.com/sixlib/go/Router"
)

const (
	port = "8080" //配置端口
)

func main() {
	Router := router.NewRouter()
	server := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Handler:      Router,
	}
	fmt.Println(`🚀 http://localhost:` + port)
	log.Fatal(server.ListenAndServe())
}
