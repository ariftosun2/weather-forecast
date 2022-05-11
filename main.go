package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ApiResult struct {
	Date        string `json:"date"`
	Day         string `json:"day"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Degree      string `json:"degree"`
	Min         string `json:"min"`
	Max         string `json:"max"`
	Night       string `json:"night"`
	Humidity    string `json:"humidity"`
}

type wantCoutry struct {
	City  string `json:"city"`
	Token string `json:"token"`
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
	req, err := ruquestWeather(objRequest.City, objRequest.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	var std1 *ApiResult

	for _, element := range req.Result {
		std1 = &ApiResult{
			Day:         element.Day,
			Date:        element.Date,
			Description: element.Description,
			Degree:      element.Degree,
			Icon:        element.Icon,
			Status:      element.Status,
			Min:         element.Min,
			Max:         element.Max,
			Night:       element.Night,
			Humidity:    element.Humidity,
		}

	}
	for _, element := range req.Result {
		fmt.Println(element.Date)
		fmt.Println(element.Day)
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	// standard output to print merged data
	err = t.Execute(os.Stdout, std1)

	return c.JSON(http.StatusOK, err)

}
func ruquestWeather(city, token string) (*ApiResponse, error) {
	url := "https://api.collectapi.com/weather/getWeather?data.lang=tr&data.city=" + city
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", token)

	resq, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resq.Body.Close()

	bodys, _ := ioutil.ReadAll(resq.Body)
	var newBody ApiResponse
	err = json.Unmarshal(bodys, &newBody)
	if err != nil {
		log.Fatal(err)
	}
	return &newBody, nil
}
