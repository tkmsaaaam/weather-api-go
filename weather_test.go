package weather_test

import (
	"fmt"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
	"github.com/tkmsaaaam/weather-api-go"
	"net/http"
	"testing"
)

func TestGetOk(t *testing.T) {
	recorderClient, _ := recorder.New("./fixtures/ok")
	defer func(recorderClient *recorder.Recorder) {
		err := recorderClient.Stop()
		if err != nil {
			fmt.Println(err)
		}
	}(recorderClient)
	weatherClient := weather.Client{Client: &http.Client{
		Transport: recorderClient,
	}}
	t.Run("Get", func(t *testing.T) {
		response, _ := weatherClient.Get("130010")
		assert.Equal(t, "東京都 東京 の天気", response.Title)
	})
}

func TestGetErr(t *testing.T) {
	recorderClient, _ := recorder.New("./fixtures/err")
	defer func(recorderClient *recorder.Recorder) {
		err := recorderClient.Stop()
		if err != nil {
			fmt.Println(err)
		}
	}(recorderClient)
	weatherClient := weather.Client{Client: &http.Client{
		Transport: recorderClient,
	}}
	t.Run("Get", func(t *testing.T) {
		response, _ := weatherClient.Get("400000")
		assert.Equal(t, "", response.Title)
	})
}
