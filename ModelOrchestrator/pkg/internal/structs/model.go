package structs

import "time"

type PostgresModel struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Url         string    `json:"connection_string"`
	DateCreated time.Time `json:"created_at,omitempty"`
}

type MongoModel struct {
	Fields  []field `json:"fields" bson:"fields"`
	ModelID int     `json:"model_id,omitempty" bson:"model_id"`
}

type field struct {
	FieldName string `json:"name" bson:"name"`
	FieldType string `json:"type" bson:"type"`
	Parent    string `json:"parent,omitempty" bson:"parent,omitempty"`
}

type Resp struct {
	PostgresModel
	MongoModel
}
