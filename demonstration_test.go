package itest

import (
	"encoding/json"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestDemonstrations(t *testing.T) {

	CreateAnnouncement(t)

	response, err := GetHttpClient().Get("demonstrations")

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode, "demonstrations")

	var body map[string]struct {
		Count int64 `json:"count"`
	}
	json.NewDecoder(response.Body).Decode(&body)
	count := body["demonstrations"].Count
	log.Infof("number of demonstrators '%d'", count)
	require.True(t, count > 0)
}
