package heartbeat

import (
	"testing"

	"github.com/SafetyCulture/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAddRequest_Validate(t *testing.T) {
	enabled := true
	request := &AddRequest{Description: "Descriptio2", Interval: 2, IntervalUnit: Minutes, Enabled: &enabled, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Name cannot be empty. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 0, IntervalUnit: Minutes, Enabled: &enabled, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Interval cannot be smaller than 1. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 10, Enabled: &enabled, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. IntervalUnit cannot be empty. ").Error())
}

func TestGetRequest_Validate(t *testing.T) {
	getRequest := &getRequest{}
	err := getRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("HeartbeatName cannot be empty").Error())
}

func TestDeleteRequest_Validate(t *testing.T) {
	deleteRequest := &deleteRequest{}
	err := deleteRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("HeartbeatName cannot be empty").Error())
}
