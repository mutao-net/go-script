package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const API = "https://api.openweathermap.org/data/2.5/weather"
const Icon = "https://openweathermap.org/img/wn/%d@2x.png"

func SetParams(appId string, lat string, lon string) string {
	//default Tokyo
	if len(lat) == 0 {
		lat = "35.681236"
	}
	if len(lon) == 0 {
		lon = "139.767125"
	}

	qs := []string{}
	params := map[string]string{
		"lat":   lat,
		"lon":   lon,
		"units": "metric",
		"lang":  "en",
		"appid": appId}
	for k, v := range params {
		q := k + "=" + url.QueryEscape(v)
		qs = append(qs, q)
	}
	return API + "?" + strings.Join(qs, "&")
}

func CallAPI(url string) ResponseResult {
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var results ResponseResult
	json.NewDecoder(r.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func main() {
	endpoint := SetParams("", "0", "0")
	fmt.Println(endpoint)
	r := CallAPI(endpoint)
	fmt.Println(r)
}

type ResponseResult struct {
	Coord struct {
		Lon int `json:"lon"`
		Lat int `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
