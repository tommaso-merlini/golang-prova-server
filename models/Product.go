package models

type Product struct {
	Name string `json:"name" bson:"name"`
	//    Pages            int    `json:"pages" bson:"pages"`
}