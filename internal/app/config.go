package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/mongo"
	"github.com/core-go/service"
	"github.com/core-go/service/cors"
)

type Config struct {
	Server     service.ServerConf `mapstructure:"server"`
	Allow      cors.AllowConfig   `mapstructure:"allow"`
	Mongo      mongo.MongoConfig  `mapstructure:"mongo"`
	Log        log.Config         `mapstructure:"log"`
	MiddleWare mid.LogConfig      `mapstructure:"middleware"`
}
