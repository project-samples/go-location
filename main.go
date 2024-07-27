package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	"github.com/core-go/core/cors"
	svr "github.com/core-go/core/server"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/gorilla/mux"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/config")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	log.Initialize(cfg.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(cfg.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	err = app.Route(r, context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(svr.ServerInfo(cfg.Server))
	c := cors.New(cfg.Allow)
	handler := c.Handler(r)
	svr.StartServer(cfg.Server, handler)
}
