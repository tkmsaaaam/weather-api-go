package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type Description struct {
	PublicTime          time.Time `json:"publicTime"`
	PublicTimeFormatted string    `json:"publicTimeFormatted"`
	HeadlineText        string    `json:"headlineText"`
	BodyText            string    `json:"bodyText"`
	Text                string    `json:"text"`
}

type Detail struct {
	Weather string `json:"weather"`
	Wind    string `json:"wind"`
	Wave    string `json:"wave"`
}

type DetailTemperature struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

type Temperature struct {
	Min DetailTemperature `json:"min"`
	Max DetailTemperature `json:"max"`
}

type ChanceOfRain struct {
	T0006 string `json:"T00_06"`
	T0612 string `json:"T06_12"`
	T1218 string `json:"T12_18"`
	T1824 string `json:"T18_24"`
}

type Image struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Forecast struct {
	Date         string       `json:"date"`
	DateLabel    string       `json:"dateLabel"`
	Telop        string       `json:"telop"`
	Detail       Detail       `json:"detail"`
	Temperature  Temperature  `json:"temperature"`
	ChanceOfRain ChanceOfRain `json:"chanceOfRain"`
	Image        Image        `json:"image"`
}

type Location struct {
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	District   string `json:"district"`
	City       string `json:"city"`
}

type CopyrightImage struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Provider struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Note string `json:"note"`
}

type Copyright struct {
	Title    string         `json:"title"`
	Link     string         `json:"link"`
	Image    CopyrightImage `json:"image"`
	Provider []Provider     `json:"provider"`
}

type NormalResponse struct {
	PublicTime          time.Time   `json:"publicTime"`
	PublicTimeFormatted string      `json:"publicTimeFormatted"`
	PublishingOffice    string      `json:"publishingOffice"`
	Title               string      `json:"title"`
	Link                string      `json:"link"`
	Description         Description `json:"description"`
	Forecasts           []Forecast  `json:"forecasts"`
	Location            Location    `json:"location"`
	Copyright           Copyright   `json:"copyright"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Client struct {
	*http.Client
}

func New() Client {
	return Client{http.DefaultClient}
}

func (client Client) Get(city string) (*NormalResponse, error) {
	re, _ := regexp.Compile(`[0-9]{6}`)
	if !re.Match([]byte(city)) {
		return nil, fmt.Errorf("weather-api-go: %s", "CITY ID is invalid.")
	}
	const baseUrl = "https://weather.tsukumijima.net/api/forecast/city/"
	url := baseUrl + city
	req, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		return nil, fmt.Errorf("weather-api-go: can not make request. %w", requestErr)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather-api-go: request is failed. %v", err)
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("weather-api-go: can not read boady. %v", readErr)
	}

	var errorResponse ErrorResponse
	json.Unmarshal(body, &errorResponse)
	if errorResponse.Error != "" {
		return nil, fmt.Errorf("weather-api-go: can not get normal result. %s", errorResponse.Error)
	}

	var response NormalResponse
	normalResParseErr := json.Unmarshal(body, &response)
	if normalResParseErr != nil {
		return nil, fmt.Errorf("weather-api-go: can not parse result. %v", normalResParseErr)
	}
	return &response, nil
}
