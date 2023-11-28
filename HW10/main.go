package main

import (
	"encoding/json"
	"fmt"
	"io"
	"main/translator"
	"main/weather"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var weatherClient weather.WeatherClient
var translateClient translator.TranslatorClient

func init() {
	// weather
	apiKey, err := readApiKey()
	if err != nil {
		panic(err)
	}

	weatherClient = weather.WeatherClient{}
	weatherClient.SetApiKey(apiKey)

	// translator
	translateClient = translator.TranslatorClient{}
}

func readApiKey() (string, error) {
	content, err := os.ReadFile("weather_apikey.txt")
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")

	weather, err := weatherClient.GetWeather(city)
	if err != nil {
		http.Error(w, "Uneble to get weather", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

type TranslateRequestParams struct {
	FormLang string `json:"from_lang"`
	ToLang   string `json:"to_lang"`
	Text     string `json:"text"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestParams TranslateRequestParams
	err = json.Unmarshal(body, &requestParams)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	translation, err := translateClient.Translate(requestParams.FormLang, requestParams.ToLang, requestParams.Text)
	if err != nil {
		http.Error(w, "Uneble to get translation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(translation)
}

func main() {
	fmt.Println("GO!")

	r := mux.NewRouter()
	r.HandleFunc("/weather", WeatherHandler).Methods("GET")
	r.HandleFunc("/translate", TranslateHandler).Methods("POST")
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error: ", err)
	}
}
