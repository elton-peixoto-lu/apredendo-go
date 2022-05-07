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
	type MyJsonName struct {
		Response struct {
			BdTxTp struct {
				Cd int64 `json:"cd"`
			} `json:"BdTxTp"`
			BizSts struct {
				Cd   string `json:"cd"`
				DtTm string `json:"dtTm"`
			} `json:"BizSts"`
			TrsrBdTradgList []struct {
				TrsrBd struct {
					FinIndxs struct {
						Cd int64  `json:"cd"`
						Nm string `json:"nm"`
					} `json:"FinIndxs"`
					AnulInvstmtRate   float64 `json:"anulInvstmtRate"`
					AnulRedRate       float64 `json:"anulRedRate"`
					Cd                int64   `json:"cd"`
					Featrs            string  `json:"featrs"`
					InvstmtStbl       string  `json:"invstmtStbl"`
					IsinCd            string  `json:"isinCd"`
					MinInvstmtAmt     float64 `json:"minInvstmtAmt"`
					MinRedQty         float64 `json:"minRedQty"`
					MinRedVal         float64 `json:"minRedVal"`
					MtrtyDt           string  `json:"mtrtyDt"`
					Nm                string  `json:"nm"`
					RcvgIncm          string  `json:"rcvgIncm"`
					SemiAnulIntrstInd bool    `json:"semiAnulIntrstInd"`
					UntrInvstmtVal    float64 `json:"untrInvstmtVal"`
					UntrRedVal        float64 `json:"untrRedVal"`
					WdwlDt            string  `json:"wdwlDt"`
				} `json:"TrsrBd"`
			} `json:"TrsrBdTradgList"`
			TrsrBondMkt struct {
				ClsgDtTm string `json:"clsgDtTm"`
				OpngDtTm string `json:"opngDtTm"`
				QtnDtTm  string `json:"qtnDtTm"`
				Sts      string `json:"sts"`
				StsCd    int64  `json:"stsCd"`
			} `json:"TrsrBondMkt"`
		} `json:"response"`
		ResponseStatus     int64  `json:"responseStatus"`
		ResponseStatusText string `json:"responseStatusText"`
		StatusInfo         string `json:"statusInfo"`
	}

	defer resp.Body.Close()

	fmt.Printf("\n\n%+v", resp)

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//fmt.Println(string(body))

	var respostas MyJsonName

	erro := json.Unmarshal(body, &respostas)
	fmt.Printf("\n\n\n")

	fmt.Printf("\n\n%+v", respostas.ResponseStatus)
	fmt.Printf("\n\n%+v", respostas.Response.TrsrBdTradgList[2])
	fmt.Printf("\n\n%+v", respostas.Response.TrsrBdTradgList[2].TrsrBd.Cd)
	//fmt.Printf("\n\n%+v", respostas.Securities[0].ProtocolNumber)
	fmt.Printf("\n\n%+v", respostas.Response.TrsrBdTradgList[1].TrsrBd.MinInvstmtAmt)
	fmt.Printf("\n\n%+v", respostas.Response.TrsrBdTradgList[1].TrsrBd.MinRedVal)

	if erro != nil {
		fmt.Println("error:", erro)
	}

}
