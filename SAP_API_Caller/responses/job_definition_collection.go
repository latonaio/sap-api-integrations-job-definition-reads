package responses

type JobDefinitionCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ObjectID       string `json:"ObjectID"`
			JobID          string `json:"JobID"`
			JobName        string `json:"JobName"`
			ExpirationDate string `json:"ExpirationDate"`
		} `json:"results"`
	} `json:"d"`
}
