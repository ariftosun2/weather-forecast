package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResult struct {
	Date        string
	Day         string
	Icon        string
	Description string
	Status      string
	Degree      string
	Min         string
	Max         string
	Night       string
	Humidity    string
}

type wantCoutry struct {
	City   string `json:"city"`
	Coutry string `json:"coutry"`
	Token  string `json:"token"`
}
type ApiResponse struct {
	Success bool        `[]byte:"success" json:"success,omitempty"`
	City    string      `[]byte:"city"    json:"city,omitempty"`
	Result  []ApiResult `[]byte:"result"  json:"result,omitempty"`
}

func main() {
	fmt.Println("havadurumu router")
	router := echo.New()

	router.GET("/weatherGet", weatherGet)
	err := router.Start(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

func weatherGet(c echo.Context) error {
	var objRequest wantCoutry
	if err := c.Bind(&objRequest); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	req, err := ruquestWeather(objRequest.City, objRequest.Coutry, objRequest.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, element := range req.Result {
		fmt.Println("Gün       : " + element.Day)
		fmt.Println("Tarih     : " + element.Date)
		fmt.Println("Açıklama  : " + element.Description)
		fmt.Println("Derece    : " + element.Degree)
		fmt.Println("")
	}
	return c.JSON(http.StatusOK, req)

}
func ruquestWeather(city, country, token string) (ApiResponse, error) {
	url := "https://api.collectapi.com/weather/getWeather?data.lang=" + "&data.city=" + country + city
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", token)

	resq, _ := http.DefaultClient.Do(req)

	defer resq.Body.Close()

	body, _ := ioutil.ReadAll(resq.Body)
	var newBody ApiResponse
	json.Unmarshal(body, &newBody)
	return newBody, nil
}
