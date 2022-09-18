package weather_test

import (
	"fmt"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
	"github.com/tkmsaaaam/weather-api-go"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		client := weather.New()
		assert.Equal(t, weather.Client{Client: http.DefaultClient}, client)
	})
}

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
		response, err := weatherClient.Get("130010")
		assert.Equal(t, "東京都 東京 の天気", response.Title)
		assert.Equal(t, nil, err)
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
		response, err := weatherClient.Get("400000")
		assert.Equal(t, "", response.Title)
		assert.Equal(t, nil, err)
	})
}
