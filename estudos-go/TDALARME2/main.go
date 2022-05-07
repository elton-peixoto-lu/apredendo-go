package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json")
	if err != nil {
		fmt.Println("No response from request")
	}

	type Responses struct {
		ResponseStatus     int    `json:"responseStatus"`
		ResponseStatusText string `json:"responseStatusText"`
		StatusInfo         string `json:"statusInfo"`
	}

	type TrsrBd struct {
		MinInvstmtAmt  float64 `json:"minInvstmtAmt"`
		UntrInvstmtVal float64 `json:"untrInvstmtVal"`
	}

	type TrsrBondMkt struct {
		OpngDtTm string `json:"opngDtTm"`
		ClsgDtTm string `json:"clsgDtTm"`
		QtnDtTm  string `json:"qtnDtTm"`
		StsCd    int    `json:"stsCd"`
		sts      string `json:"sts"`
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))

	var respostas Responses

	var teste1 TrsrBd

	var statusmercado TrsrBondMkt

	erro := json.Unmarshal(body, &respostas)
	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("\n\n%+v", respostas.ResponseStatus)

	// <nil> Testando

	erro = json.Unmarshal(body, &teste1)

	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("\n\n%+v", teste1.UntrInvstmtVal)

	if respostas.ResponseStatus == 200 {
		fmt.Printf("\n\n%+v", "sucesso")
	}

	erro = json.Unmarshal(body, &statusmercado)
	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("\n\n%+v", statusmercado.StsCd)

}
