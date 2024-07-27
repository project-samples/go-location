package app

import (
	"github.com/core-go/core/cors"
	"github.com/core-go/core/server"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Config struct {
	Server     server.ServerConf `mapstructure:"server"`
	Allow      cors.AllowConfig  `mapstructure:"allow"`
	Mongo      MongoConfig       `mapstructure:"mongo"`
	Log        log.Config        `mapstructure:"log"`
	MiddleWare mid.LogConfig     `mapstructure:"middleware"`
}

type MongoConfig struct {
	Uri      string `yaml:"uri" mapstructure:"uri" json:"uri,omitempty" gorm:"column:uri" bson:"uri,omitempty" dynamodbav:"uri,omitempty" firestore:"uri,omitempty"`
	Database string `yaml:"database" mapstructure:"database" json:"database,omitempty" gorm:"column:database" bson:"database,omitempty" dynamodbav:"database,omitempty" firestore:"database,omitempty"`
}
