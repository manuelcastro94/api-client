package f1_api_client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestAlonsoSixSeasonsWithRenaultURL(t *testing.T) {
	httpClient := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	f1Client := Create(&httpClient)
	result := f1Client.Drivers("alonso").Constructors("renault").Seasons("").Query("json")

	data := result["MRData"]
	if data == nil {
		t.Failed()
	}
	seasonTable := data.(map[string]interface{})["SeasonTable"]
	if seasonTable == nil {
		t.Failed()
	}
	seasons := seasonTable.(map[string]interface{})["Seasons"]
	if seasons == nil {
		t.Failed()
	}
	assert.Equal(t, 6, len(seasons.([]interface{})))
}
