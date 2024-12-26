package main

import (
	"log"
	"net/http"
	"time"

	"github.com/imhasandl/grpc-go-project/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (srv *httpServer) Run() {
	mux := http.NewServeMux()

	orderService := service.NewOrderService()

	// err := http.ListenAndServe(srv.addr, router)

	server := &http.Server{
		Addr:              srv.addr,
		Handler:           mux,
		ReadHeaderTimeout: 30 * time.Second,
	}

	log.Println("Starting server on port: ", srv.addr)

	log.Fatal(server.ListenAndServe())
}
