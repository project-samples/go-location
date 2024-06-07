package app

import (
	"github.com/core-go/core"
	"github.com/core-go/core/cors"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/mongo"
)

type Config struct {
	Server     core.ServerConf   `mapstructure:"server"`
	Allow      cors.AllowConfig  `mapstructure:"allow"`
	Mongo      mongo.MongoConfig `mapstructure:"mongo"`
	Log        log.Config        `mapstructure:"log"`
	MiddleWare mid.LogConfig     `mapstructure:"middleware"`
}
