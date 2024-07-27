package location

import "github.com/core-go/search"

type LocationFilter struct {
	*search.Filter
	Id          string   `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Type        string   `mapstructure:"type" json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required,max=40" q:"equal"`
	Name        string   `mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty" validate:"required,max=255" match:"prefix" q:"prefix"`
	Description string   `mapstructure:"description" json:"description,omitempty" gorm:"column:description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty" q:"true"`
	Status      []string `mapstructure:"status" json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" validate:"required,max=255"`
	CustomURL   string   `mapstructure:"custom_url" json:"customURL,omitempty" gorm:"column:customurl" bson:"customURL,omitempty" dynamodbav:"customURL,omitempty" firestore:"customURL,omitempty"`
	Latitude    *float64 `mapstructure:"latitude" json:"latitude,omitempty" gorm:"column:latitude" bson:"-" dynamodbav:"latitude,omitempty" firestore:"latitude,omitempty"`
	Longitude   *float64 `mapstructure:"longitude" json:"longitude,omitempty" gorm:"column:longitude" bson:"-" dynamodbav:"longitude,omitempty" firestore:"longitude,omitempty"`
}
