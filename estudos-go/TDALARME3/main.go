package main
​
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
​
func main() {
	client := newAPIClient()
	resp, err := client.Get("https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json")
	if err != nil {
		fmt.Println("Get: ", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll: ", err.Error())
		return
	}
​
	// fmt.Println(string(body))
​
	response := &Responses{}
	err = json.Unmarshal(body, response)
	if err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return
	}
	trlist := response.Response.TrsrBdTradgList
	fmt.Println("trlist len:", len(trlist))
​
	for _, tr := range trlist {
		fmt.Printf("MinInvstmtAmt:%f | UntrInvstmtVal:%f \n", tr.TrsrBd.MinInvstmtAmt, tr.TrsrBd.UntrInvstmtVal)
	}
}
​
func newAPIClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
​
	return &http.Client{Transport: tr}
}
​
type Responses struct {
	ResponseStatus     int      `json:"responseStatus"`
	ResponseStatusText string   `json:"responseStatusText"`
	StatusInfo         string   `json:"statusInfo"`
	Response           Response `json:"response"`
}
​
type TrsrBd struct {
	MinInvstmtAmt  float64 `json:"minInvstmtAmt"`
	UntrInvstmtVal float64 `json:"untrInvstmtVal"`
}
​
type TrsrBdTradgList struct {
	TrsrBd TrsrBd `json:"TrsrBd,omitempty"`
}
​
type Response struct {
	TrsrBdTradgList []TrsrBdTradgList `json:"TrsrBdTradgList"`
}