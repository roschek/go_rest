package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/roschek/go_rest.git/internal/user"
	"github.com/roschek/go_rest.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("CREATE ROUTER")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listaning on port 0.0.0.0:1234")
	log.Fatalln(server.Serve(listener))
}
