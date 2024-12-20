package structs

type MappingModel struct {
	Id             string           `json:"_id,omitempty" bson:"_id,omitempty"`
	SourceId       []int            `json:"source_id" bson:"source_id"`
	ModelId        int              `json:"model_id" bson:"model_id"`
	Mapping        []mapping        `json:"mapping" bson:"mapping"`
	Transformation []transformation `json:"transformation" bson:"transformation"`
	Filters        []filter         `json:"filters" bson:"filters"`
}

type mapping struct {
	SourcePath string `json:"source_path" bson:"source_path"`
	ModelPath  string `json:"model_path" bson:"model_path"`
}

type transformation struct {
	FieldPath string `json:"field_path" bson:"field_path"`
	Func      string `json:"func" bson:"func"`
	Msg       string `json:"msg,omitempty" bson:"msg,omitempty"`
}

type filter struct {
	FieldPath string `json:"field_path" bson:"field_path"`
	Func      string `json:"func" bson:"func"`
	Arg       any    `json:"arg" bson:"arg"`
}
