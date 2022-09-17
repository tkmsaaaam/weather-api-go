package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Response struct {
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

func Get(city string) (Response, error) {
	const baseUrl = "https://weather.tsukumijima.net/api/forecast/city/"
	url := baseUrl + city
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	var response Response
	if err != nil {
		fmt.Println("Error Request:", err)
		return response, err
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("ioutil.ReadAll err:", readErr)
		return response, readErr
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		fmt.Println("json.Unmarshal err:", jsonErr)
		return response, jsonErr
	}
	return response, nil
}
