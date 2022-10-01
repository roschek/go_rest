package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/roschek/go_rest.git/internal/config"
	"github.com/roschek/go_rest.git/internal/user"
	"github.com/roschek/go_rest.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("CREATE ROUTER")
	router := httprouter.New()

	cfg := config.GetConfig()
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket at %s", socketPath)

		logger.Info("listen unit socket")
		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Info("listen unit tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Info(fmt.Sprintf("server is listaning on port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
