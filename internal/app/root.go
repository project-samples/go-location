package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/mongo"
	"github.com/core-go/service"
)

type Root struct {
	Server     service.ServerConfig `mapstructure:"server"`
	Mongo      mongo.MongoConfig    `mapstructure:"mongo"`
	Log        log.Config           `mapstructure:"log"`
	MiddleWare mid.LogConfig        `mapstructure:"middleware"`
}
