package weather_test

import (
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tkmsaaaam/weather-api-go"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		client := weather.New()
		expected := weather.Client{Client: http.DefaultClient}
		if expected != client {
			t.Errorf("add() = %v, want %v", expected, client)
		}
	})
}

type localRoundTripper struct {
	handler http.Handler
}

func (localRoundTripper localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ressponseRecorder := httptest.NewRecorder()
	localRoundTripper.handler.ServeHTTP(ressponseRecorder, req)
	return ressponseRecorder.Result(), nil
}

//go:embed testdata/*
var testData embed.FS

func TestGetOk(t *testing.T) {
	type want struct {
		Title string
		err   error
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "ok",
			want: want{Title: "東京都 東京 の天気", err: nil},
		},
	}
	for _, tt := range tests {
		mux := http.NewServeMux()
		client := weather.Client{&http.Client{Transport: localRoundTripper{handler: mux}}}
		mux.HandleFunc("/api/forecast/city/130010", func(w http.ResponseWriter, req *http.Request) {
			res, _ := testData.ReadFile("testdata/130010.json")
			w.Write(res)
		})
		t.Run(tt.name, func(t *testing.T) {
			response, err := client.Get("130010")
			if response.Title != tt.want.Title {
				t.Errorf("add() = %v, want %v", response.Title, tt.want.Title)
			}
			if err != tt.want.err {
				t.Errorf("add() = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestGetErr(t *testing.T) {
	type want struct {
		response *weather.NormalResponse
		err      string
	}

	tests := []struct {
		name string
		id   string
		want want
	}{
		{
			name: "id is invalid",
			id:   "1",
			want: want{response: nil, err: "weather-api-go: CITY ID is invalid."},
		},
		{
			name: "not found",
			id:   "400000",
			want: want{response: nil, err: "weather-api-go: can not get normal result. The specified city ID is invalid."},
		},
	}
	for _, tt := range tests {
		mux := http.NewServeMux()
		client := weather.Client{&http.Client{Transport: localRoundTripper{handler: mux}}}
		mux.HandleFunc("/api/forecast/city/"+tt.id, func(w http.ResponseWriter, req *http.Request) {
			res, _ := testData.ReadFile("testdata/" + tt.id + ".json")
			w.Write(res)
		})
		t.Run(tt.name, func(t *testing.T) {
			response, err := client.Get(tt.id)
			if response != tt.want.response {
				t.Errorf("add() = %v, want %v", response, tt.want.response)
			}
			if err.Error() != tt.want.err {
				t.Errorf("add() = %v, want %v", err, tt.want.err)
			}
		})
	}
}
