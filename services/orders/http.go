package main

import (
	"log"
	"net/http"

	"github.com/imhasandl/grpc-go-project/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (srv *httpServer) Run() {
	router := http.NewServeMux()

	orderService := service.NewOrderService()

	log.Println("Starting server on port: ", srv.addr)

	err := http.ListenAndServe(srv.addr, router)

	log.Fatal(err)
}