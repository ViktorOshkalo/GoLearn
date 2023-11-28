package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Location struct {
		Name      string
		Country   string
		Localtime string
		TimeZone  string `json:"tz_id"`
	}
	Current struct {
		LastUpdateTime string  `json:"last_updated"`
		Temperature    float32 `json:"temp_c"`
		Feelslike      float32 `json:"feelslike_c"`
		Wind           float32 `json:"wind_kph"`
		Humidity       float32
		Condition      struct {
			Text string
		}
	}
}

type Weather struct {
	LocationName   string
	Country        string
	Localtime      string
	TimeZone       string
	LastUpdateTime string
	Temperature    float32
	Feelslike      float32
	Wind           float32
	Humidity       float32
	Condition      string
}

type WeatherClient struct {
	Configuration struct {
		apiKey string
	}
}

func (client *WeatherClient) SetApiKey(apiKey string) {
	client.Configuration.apiKey = apiKey
}

func (client WeatherClient) GetWeather(city string) (*Weather, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?&q=%s&key=%s", city, client.Configuration.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status code %d", resp.StatusCode)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	wheatherPresentation := getWhetherPresentation(weatherResp)
	return &wheatherPresentation, nil
}

func getWhetherPresentation(wr WeatherResponse) Weather {
	return Weather{
		LocationName:   wr.Location.Name,
		Country:        wr.Location.Country,
		Localtime:      wr.Location.Localtime,
		TimeZone:       wr.Location.TimeZone,
		LastUpdateTime: wr.Current.LastUpdateTime,
		Temperature:    wr.Current.Temperature,
		Feelslike:      wr.Current.Feelslike,
		Wind:           wr.Current.Wind,
		Humidity:       wr.Current.Humidity,
		Condition:      wr.Current.Condition.Text,
	}
}
