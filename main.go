package main

import (
	"context"
	"fmt"
	"go-service/internal/app"
	"net/http"

	"github.com/core-go/config"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	sv "github.com/core-go/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()
	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println(sv.ServerInfo(conf.Server))
	// http.ListenAndServe(sv.Addr(conf.Server.Port), r)
	c := cors.New(cors.Options{
		AllowedHeaders:   conf.Allow.AllowHeaders,
		AllowedOrigins:   conf.Allow.AllowOrigins,
		AllowedMethods:   conf.Allow.AllowMethods,
		AllowCredentials: conf.Allow.Credentials})
	handler := c.Handler(r)
	if conf.Allow.Https {
		http.ListenAndServeTLS(conf.Allow.SecurePort, conf.Allow.Cert, conf.Allow.Key, handler)
	} else {
		http.ListenAndServe(sv.Addr(conf.Server.Port), handler)
	}
}
