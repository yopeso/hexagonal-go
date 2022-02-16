package shortener

type LinkRedirect struct {
	Code      string `json:"code" msgpack:"code" bson:"code"`
	URL       string `json:"url" msgpack:"url" bson:"url"`
	CreatedAt int64  `json:"created_at" msgpack:"created_at" bson:"created_at"`
}
