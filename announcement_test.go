package itest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Announcement struct {
	UserId     string   `json:"user_id"`
	UserDevice Device   `json:"device_id"`
	SeenDevice Device   `json:"seen_device"`
	Location   Location `json:"location"`
	Timestamp  int64    `json:"timestamp"`
}

type Location struct {
	Latitute  float64 `json:"latitute"`
	Longitude float64 `json:"longitude"`
}

type Device struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func TestAnnouncement(t *testing.T) {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(map[string][]*Announcement{
		"announcements": {{
			UserId:     "test",
			UserDevice: Device{Id: "test-1", Type: "iphone 14"},
			SeenDevice: Device{Id: "test-2", Type: "iphone 15"},
			Location:   Location{Latitute: 32.05766501361105, Longitude: 34.76640727232065},
			Timestamp:  time.Now().Unix(),
		}}}))

	response, err := GetHttpClient().Post(GetAnnouncementsUrl(), &buf)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, response.StatusCode, "announcements")
}
