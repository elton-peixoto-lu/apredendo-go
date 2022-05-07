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

	//	type Responses struct {
	//		ResponseStatus     int    `json:"responseStatus"`
	//		ResponseStatusText string `json:"responseStatusText"`
	//		StatusInfo         string `json:"statusInfo"`
	//	}

	/*	type TrsrBd struct {
		MinInvstmtAmt  float64 `json:"minInvstmtAmt"`
		UntrInvstmtVal float64 `json:"untrInvstmtVal"`
	}*/

	type TrsrBondMkt []struct {
		OpngDtTm string `json:"opngDtTm"`
		ClsgDtTm string `json:"clsgDtTm"`
		QtnDtTm  string `json:"qtnDtTm"`
		StsCd    int    `json:"stsCd"`
		Sts      int    `json:"sts"`
	}

	//	type TrsrBdTradgList struct {
	//		Cd                int     `json:"cd"`
	//		Nm                string  `json:"nm"`
	//		Featrs            string  `json:"featrs"`
	//		MtrtyDt           string  `json:"mtrtyDt"`
	/*		MinInvstmtAmt     float64 `json:"minInvstmtAmt"`
			UntrInvstmtVal    float64 `json:"untrInvstmtVal"`
			InvstmtStbl       string  `json:"invstmtStbl"`
			SemiAnulIntrstInd bool    `json:"semiAnulIntrstInd"`
			RcvgIncm          string  `json:"rcvgIncm"`
			AnulInvstmtRate   float64 `json:"anulInvstmtRate"`
			AnulRedRate       float64 `json:"anulRedRate"`
			MinRedQty         float64 `json:"minRedQty"`
			UntrRedVal        float64 `json:"untrRedVal"`
			MinRedVal         float64 `json:"minRedVal"`
			IsinCd            string  `json:"isinCd"`
		} */

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))

	//var respostas Responses

	//var teste1 TrsrBd

	var statusmercado TrsrBondMkt

	//var teste2 TrsrBdTradgList

	/*	erro := json.Unmarshal(body, &respostas)
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
		}*/

	erro := json.Unmarshal(body, &statusmercado)
	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("\n\n%+v", statusmercado)

	//	erro = json.Unmarshal(body, &teste2)
	//	if erro != nil {
	//		fmt.Println("error:", erro)
	//	}
	//	fmt.Printf("\n\n%+v", teste2)

}
