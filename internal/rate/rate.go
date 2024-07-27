package rate

import "time"

type Rate struct {
	Id         string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	LocationId string     `mapstructure:"locationId" json:"locationId,omitempty" gorm:"column:locationid" bson:"locationId,omitempty" dynamodbav:"locationId,omitempty" firestore:"locationId,omitempty" validate:"required,max=255"`
	RateTime   *time.Time `mapstructure:"rateTime" json:"rateTime,omitempty" gorm:"column:ratetime" bson:"rateTime,omitempty" dynamodbav:"rateTime,omitempty" firestore:"-"`
	UserId     string     `mapstructure:"userId" json:"userId,omitempty" gorm:"column:userId" bson:"updatedBy,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Rate       int32      `mapstructure:"rate" json:"rate" gorm:"column:rate" bson:"rate" dynamodbav:"rate" firestore:"rate"`
	Version    int        `mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
	Review     string     `mapstructure:"review" json:"review,omitempty" gorm:"column:review" bson:"review,omitempty" dynamodbav:"review,omitempty" firestore:"review,omitempty" validate:"required,max=255" match:"prefix" q:"prefix"`
}
