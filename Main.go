package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	router "github.com/sixlib/go/Router"
)

const (
	baseAddr = ":3000" //配置端口
)

/**
 * Addr 端口
 */
func serve(Addr string) {
	Router := router.NewRouter()

	server := http.Server{
		Addr:         Addr,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Handler:      Router,
	}
	fmt.Println(`🚀 http://localhost` + Addr)
	log.Fatal(server.ListenAndServe())
}
func main() {
	serve(`:8080`)
}
