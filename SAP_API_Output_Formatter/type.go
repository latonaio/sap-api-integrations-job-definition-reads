package sap_api_output_formatter

type JobDefinitionCollectionReads struct {
	ConnectionKey               string `json:"connection_key"`
	Result                      bool   `json:"result"`
	RedisKey                    string `json:"redis_key"`
	Filepath                    string `json:"filepath"`
	APISchema                   string `json:"api_schema"`
	JobDefinitionCollectionCode string `json:"jobdefinitioncollection_code"`
	Deleted                     bool   `json:"deleted"`
}

type JobDefinitionCollection struct {
	ObjectID       string `json:"ObjectID"`
	JobID          string `json:"JobID"`
	JobName        string `json:"JobName"`
	ExpirationDate string `json:"ExpirationDate"`
}
