package structs

import "time"

type LogsModel struct {
	Id               int           `json:"id,omitempty"`
	SourceId         []int32       `json:"source_id"`
	ModelId          int           `json:"model_id"`
	Time             time.Duration `json:"time"`
	TimeWithResponse time.Duration `json:"time_with_response"`
	Status           string        `json:"status"`
	CreationDate     time.Time     `json:"creation_date,omitempty"`
}
