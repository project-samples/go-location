package location

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
