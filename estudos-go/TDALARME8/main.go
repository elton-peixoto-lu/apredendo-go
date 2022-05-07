package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"encoding/json"
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
		Response           struct {
			BdTxTp struct {
				Cd int `json:"cd"`
			} `json:"BdTxTp"`
			TrsrBondMkt struct {
				OpngDtTm string `json:"opngDtTm"`
				ClsgDtTm string `json:"clsgDtTm"`
				QtnDtTm  string `json:"qtnDtTm"`
				StsCd    int    `json:"stsCd"`
				Sts      string `json:"sts"`
			} `json:"TrsrBondMkt"`
			TrsrBdTradgList []struct {
				TrsrBd struct {
					Cd                int     `json:"cd"`
					Nm                string  `json:"nm"`
					Featrs            string  `json:"featrs"`
					MtrtyDt           string  `json:"mtrtyDt"`
					MinInvstmtAmt     float64 `json:"minInvstmtAmt"`
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
					FinIndxs          struct {
						Cd int    `json:"cd"`
						Nm string `json:"nm"`
					} `json:"FinIndxs"`
					WdwlDt string `json:"wdwlDt"`
				} `json:"TrsrBd,omitempty"`

				BizSts struct {
					Cd   string `json:"cd"`
					DtTm string `json:"dtTm"`
				} `json:"BizSts"`
			} `json:"response"`
		}
	}
	defer resp.Body.Close()

	fmt.Printf("\n\n%+v", resp)

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//fmt.Println(string(body))

	var respostas Responses
	var VARIAVEL Responses

	erro := json.Unmarshal(body, &respostas)
	fmt.Printf("\n\n\n")

	fmt.Printf("\n\n%+v", respostas.ResponseStatus)
	fmt.Printf("\n\n%+v", TrsrBdTradgList[1])
	//fmt.Printf("\n\n%+v", respostas.Securities[0].ProtocolNumber)

	if erro != nil {
		fmt.Println("error:", erro)
	}

}
