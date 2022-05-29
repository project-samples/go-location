package location

import (
	"github.com/core-go/mongo/geo"
	"time"
)

type Location struct {
	Id          string        `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Type        string        `mapstructure:"type" json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required,max=40"`
	Name        string        `mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty" validate:"required,max=255"`
	Description string        `mapstructure:"description" json:"description,omitempty" gorm:"column:description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	Status      string        `mapstructure:"status" json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" validate:"required,max=255"`
	ImageURL    string        `mapstructure:"image_url" json:"imageURL,omitempty" gorm:"column:imageurl" bson:"imageURL,omitempty" dynamodbav:"imageURL,omitempty" firestore:"imageURL,omitempty"`
	CustomURL   string        `mapstructure:"custom_url" json:"customURL,omitempty" gorm:"column:customurl" bson:"customURL,omitempty" dynamodbav:"customURL,omitempty" firestore:"customURL,omitempty"`
	Latitude    *float64      `mapstructure:"latitude" json:"latitude,omitempty" gorm:"column:latitude" bson:"-" dynamodbav:"latitude,omitempty" firestore:"latitude,omitempty"`
	Longitude   *float64      `mapstructure:"longitude" json:"longitude,omitempty" gorm:"column:longitude" bson:"-" dynamodbav:"longitude,omitempty" firestore:"longitude,omitempty"`
	Geo         *geo.JSON     `mapstructure:"geo" json:"-" bson:"geo,omitempty" gorm:"-" dynamodbav:"-" firestore:"-"`
	Info        *LocationInfo `mapstructure:"info" json:"info,omitempty" bson:"-" gorm:"column:info" dynamodbav:"info,omitempty" firestore:"info,omitempty"`
	CreatedBy   string        `mapstructure:"created_by" json:"createdBy,omitempty" gorm:"column:createdby" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt   *time.Time    `mapstructure:"created_at" json:"createdAt,omitempty" gorm:"column:createdat" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"-"`
	UpdatedBy   string        `mapstructure:"updated_by" json:"updatedBy,omitempty" gorm:"column:updatedby" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt   *time.Time    `mapstructure:"updated_at" json:"updatedAt,omitempty" gorm:"column:updatedat" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"-"`
	Version     int           `mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
}
type LocationInfo struct {
	Id           string  `json:"-"  gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id"`
	Rate         float64 `json:"rate" gorm:"column:rate" bson:"rate" dynamodbav:"rate" firestore:"rate"`
	Rate1        int32   `json:"rate1" gorm:"column:rate1" bson:"rate1" dynamodbav:"rate1" firestore:"rate1"`
	Rate2        int32   `json:"rate2" gorm:"column:rate2" bson:"rate2" dynamodbav:"rate2" firestore:"rate2"`
	Rate3        int32   `json:"rate3" gorm:"column:rate3" bson:"rate3" dynamodbav:"rate3" firestore:"rate3"`
	Rate4        int32   `json:"rate4" gorm:"column:rate4" bson:"rate4" dynamodbav:"rate4" firestore:"rate4"`
	Rate5        int32   `json:"rate5" gorm:"column:rate5" bson:"rate5" dynamodbav:"rate5" firestore:"rate5"`
	RateLocation float64 `json:"rateLocation" gorm:"column:ratelocation" bson:"rateLocation" dynamodbav:"rateLocation" firestore:"rateLocation"`
}
type LocationRate struct {
	Id         string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	LocationId string     `mapstructure:"locationId" json:"locationId,omitempty" gorm:"column:locationid" bson:"locationId,omitempty" dynamodbav:"locationId,omitempty" firestore:"locationId,omitempty" validate:"required,max=255"`
	RateTime   *time.Time `mapstructure:"rateTime" json:"rateTime,omitempty" gorm:"column:ratetime" bson:"rateTime,omitempty" dynamodbav:"rateTime,omitempty" firestore:"-"`
	UserId     string     `mapstructure:"userId" json:"userId,omitempty" gorm:"column:userId" bson:"updatedBy,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Rate       int32      `mapstructure:"rate" json:"rate" gorm:"column:rate" bson:"rate" dynamodbav:"rate" firestore:"rate"`
	Version    int        `mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
	Review     string     `mapstructure:"review" json:"review,omitempty" gorm:"column:review" bson:"review,omitempty" dynamodbav:"review,omitempty" firestore:"review,omitempty" validate:"required,max=255" match:"prefix" q:"prefix"`
}
