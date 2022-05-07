package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//	"encoding/json"
	"github.com/fastrodev/fastrex"
)

func main() {

	app := fastrex.New()
	app.Template("web/modules/root/build/index.html")
	app.Static("web/modules/root/build")
	app.Get("/", func(req fastrex.Request, res fastrex.Response) {
		err := res.Render()
		if err != nil {
			panic(err)
		}
	})
	err := app.Listen(8080)
	if err != nil {
		panic(err)
	}
	resp, err := http.Get("https://api.hgbrasil.com/weather?woeid=455884")
	if err != nil {
		fmt.Println("No response from request")

	}

	type respostasJson struct {
		By            string `json:"by"`
		ExecutionTime int64  `json:"execution_time"`
		FromCache     bool   `json:"from_cache"`
		Results       struct {
			Cid           string `json:"cid"`
			City          string `json:"city"`
			CityName      string `json:"city_name"`
			ConditionCode string `json:"condition_code"`
			ConditionSlug string `json:"condition_slug"`
			Currently     string `json:"currently"`
			Date          string `json:"date"`
			Description   string `json:"description"`
			Forecast      []struct {
				Condition   string `json:"condition"`
				Date        string `json:"date"`
				Description string `json:"description"`
				Max         int64  `json:"max"`
				Min         int64  `json:"min"`
				Weekday     string `json:"weekday"`
			} `json:"forecast"`
			Humidity   int64  `json:"humidity"`
			ImgID      string `json:"img_id"`
			Sunrise    string `json:"sunrise"`
			Sunset     string `json:"sunset"`
			Temp       int64  `json:"temp"`
			Time       string `json:"time"`
			WindSpeedy string `json:"wind_speedy"`
		} `json:"results"`
		ValidKey bool `json:"valid_key"`
	}

	defer resp.Body.Close()

	//fmt.Printf("\n\n%+v", resp)
	var respostas respostasJson

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &respostas)

	// response body is []byte
	//fmt.Println(string(body))

	fmt.Printf("\n\n%+v", respostas.Results.CityName)
	fmt.Printf("\n\n%+v", respostas.Results.Forecast[1].Max)

}
