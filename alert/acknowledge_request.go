package alert

import (
	"net/http"

	"github.com/SafetyCulture/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type AcknowledgeAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User            string `json:"user,omitempty"`
	Source          string `json:"source,omitempty"`
	Note            string `json:"note,omitempty"`
}

func (r *AcknowledgeAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *AcknowledgeAlertRequest) ResourcePath() string {

	return "/v2/alerts/" + r.IdentifierValue + "/acknowledge"

}

func (r *AcknowledgeAlertRequest) Method() string {
	return http.MethodPost
}

func (r *AcknowledgeAlertRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == ALIAS {
		params["identifierType"] = "alias"

	} else if r.IdentifierType == TINYID {
		params["identifierType"] = "tiny"

	} else {
		params["identifierType"] = "id"

	}
	return params
}
