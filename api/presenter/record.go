package presenter

type Record struct {
	Key        string `json:"title"`
	TotalCount int    `json:"totalCount"`
	CreatedAt  string `json:"createdAt"`
}

//type Record struct {
//	ID         primitive.ObjectID `bson:"_id"`
//	Key        string             `bson:"text"`
//	TotalCount int                `bson:"totalCount"`
//	CreatedAt  time.Time          `bson:"created_at"`
//}