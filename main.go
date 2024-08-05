package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// weather data structure
type wetherData struct {
	Name string `json:"Name"`
	Main struct {
		Kelvin  float64 `json:"Kelvin"`
		Celsius float64 `json:"Celsius"`
	} `json:"Main"`
	Wind struct {
		Speed float64 `json:"Speed"`
	} `json:"Wind"`
	Weather []struct {
		Description string `json:"Description"`
	} `json:"Weather"`
}

// apikeyt structure

type apiconfigdata struct {
	Apikey string `json:"Apikey"`
}

// function to get our api key / error
func getapikey(fileName string) (apiconfigdata, error) {
	// reading the input file
	byte, err := os.ReadFile(fileName)
	if err != nil {
		return apiconfigdata{}, err
	}
	var ch apiconfigdata
	err = json.Unmarshal(byte, &ch)

	if err != nil {
		return apiconfigdata{}, err
	}

	return ch, nil

}
func Goproject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte("welcome to check weather app"))
	if err != nil {
		fmt.Printf("error%v\n", err)
	}

}

func query(city string) (wetherData, error) {
	apiconfigdata, err := getapikey(".apiconfig.json")

	if err != nil {
		return wetherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/3.0/weather?&appid=" + apiconfigdata.Apikey + "&q=" + city)
	if err != nil {
		return wetherData{}, err
	}
	defer resp.Body.Close()

	var d wetherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return wetherData{}, err
	}
	return d, nil

}

func report(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	City := params["city"]
	data, err := query(City)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
	r.HandleFunc("/start", Goproject)
	r.HandleFunc("/weather/{city}", report)

}
