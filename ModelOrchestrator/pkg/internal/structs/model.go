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
	FieldPath string `json:"path" bson:"path"`
	FieldType string `json:"type" bson:"type"`
}

type Resp struct {
	PostgresModel
	MongoModel
}
