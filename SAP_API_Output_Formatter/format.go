package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-job-definition-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToJobDefinitionCollection(raw []byte, l *logger.Logger) ([]JobDefinitionCollection, error) {
	pm := &responses.JobDefinitionCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to JobDefinitionCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	jobDefinitionCollection := make([]JobDefinitionCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		jobDefinitionCollection = append(jobDefinitionCollection, JobDefinitionCollection{
			ObjectID:       data.ObjectID,
			JobID:          data.JobID,
			JobName:        data.JobName,
			ExpirationDate: data.ExpirationDate,
		})
	}

	return jobDefinitionCollection, nil
}
