package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://youtube-video-downloader2.p.rapidapi.com/?url=https://www.youtube.com/watch?v=imNtSPM3-r4&ab_channel=AltShiftX"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "youtube-video-downloader2.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "6193ed6f16mshbdf7bea7d9ddf11p19b669jsncdc35fcfa66c")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
