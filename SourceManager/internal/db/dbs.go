package db

import "time"

type MongoData struct {
	Fields   []Field `json:"fields" bson:"fields"`
	SourceID int     `json:"source_id" bson:"source_id"`
}

type Field struct {
	FieldName string `json:"name" bson:"name"`
	FieldType string `json:"type" bson:"type"`
	Parent    string `json:"parent,omitempty" bson:"parent,omitempty"`
}

type PostgresData struct {
	PostgresID int       `json:"id,omitempty"`
	Name       string    `json:"name"`
	URL        string    `json:"connection_string"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type Source struct {
	MongoData
	PostgresData
}
