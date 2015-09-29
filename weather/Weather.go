package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://weather.livedoor.com/forecast/webservice/json/v1?city=110010"
	res, _ := http.Get(url)
	defer res.Body.Close()
	byteArr, _ := ioutil.ReadAll(res.Body)
	var weatherData WeatherData

	err := json.Unmarshal(byteArr, &weatherData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weatherData.Description.Text)
}

type WeatherData struct {
	Location         Location
	Title            string
	Link             string
	PublicTime       string
	Description      Description
	Forecasts        []Forecasts
	PinpointLocation []PinpointLocation
	Copyright        Copyright
}

type Location struct {
	Area string
	Pref string
	City string
}

type Description struct {
	Text       string
	PublicTime string
}

type Forecasts struct {
	Date        string
	DateLabel   string
	Telop       string
	Image       Image
	Temperature Temperature
}

type Image struct {
	Title  string
	Link   string
	Url    string
	Width  int
	Height int
}

type Temperature struct {
	Celsius    MaxMin
	Fahrenheit MaxMin
}

type MaxMin struct {
	Max string
	Min string
}

type PinpointLocation struct {
	Name string
	Link string
}

type Copyright struct {
	Title    string
	Link     string
	Image    Image
	Provider string
}
