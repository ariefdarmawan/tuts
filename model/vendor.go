package tutsobj

type Vendor struct {
	ID    string `bson:"_id" json:"_id"`
	Title string
}
