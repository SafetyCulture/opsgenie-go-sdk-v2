package deployment

import (
	"net/http"

	"github.com/SafetyCulture/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type GetRequestStatusRequest struct {
	client.BaseRequest
	RequestId string `json:"requestId,omitempty"`
}

func (r *GetRequestStatusRequest) Validate() error {
	if r.RequestId == "" {
		return errors.New("RequestId can not be empty")
	}

	return nil
}

func (r *GetRequestStatusRequest) ResourcePath() string {
	return "/v2/deployments/requests/" + r.RequestId
}

func (r *GetRequestStatusRequest) Method() string {
	return http.MethodGet
}
