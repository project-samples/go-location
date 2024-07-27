package rate

import (
	"github.com/core-go/search"
	"time"
)

type RateFilter struct {
	*search.Filter
	Id         string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Review     string     `mapstructure:"review" json:"review" gorm:"column:review" bson:"review" dynamodbav:"review,omitempty" firestore:"review,omitempty" `
	LocationId string     `mapstructure:"locationId" json:"locationId,omitempty" gorm:"column:locationid" bson:"locationId,omitempty" dynamodbav:"locationId,omitempty" firestore:"locationId,omitempty" validate:"required,max=255"`
	UserId     string     `mapstructure:"userId" json:"userId,omitempty" gorm:"column:userId" bson:"userId,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	RateTime   *time.Time `mapstructure:"rate_time" json:"rateTime,omitempty" gorm:"column:ratetime" bson:"ratetime,omitempty" dynamodbav:"rateTime,omitempty" firestore:"-"`
}
