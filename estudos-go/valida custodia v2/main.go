package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func qualquernome(conta string, respostas string, Responses string) {
	fmt.Println("Insira a conta")
	fmt.Scanln(&conta)
	fmt.Println("A conta analisada sera : ")
	fmt.Println("https://custody-treasurydirect-backgroundservice.kube.prod.easynvest.io/v1/custodies/" + conta)

	resp, err := http.Get("https://custody-treasurydirect-backgroundservice.kube.prod.easynvest.io/v1/custodies/" + conta)

	if err != nil {
		fmt.Println("No response from request")

		defer resp.Body.Close()

		//fmt.Printf("\n\n%+v", resp)

		body, err := ioutil.ReadAll(resp.Body) // response body is []byte

		var respostas Responses

		erro := json.Unmarshal(body, &respostas)
		fmt.Printf("\n\n\n")

		if erro != nil {
			fmt.Println("error:", erro)
		}
		for i := 0; i < len(respostas.Securities); i++ { //enquanto i for menor que a extensão da string
			fmt.Printf("\n\n\n")
			fmt.Printf("Título Nome")
			fmt.Printf("\n\n%+v", respostas.Securities[i].Nickname)
			fmt.Printf("\n\n\n")
			fmt.Printf("numero Protocolo ")
			fmt.Printf("\n\n%+v", respostas.Securities[i].ProtocolNumber)
			fmt.Printf("\n\n\n")
			fmt.Printf("Data do pedido ")
			fmt.Printf("\n\n%+v", respostas.Securities[i].OrderProcessedAt)

			if respostas.Securities[i].ProtocolNumber != 0 {
				fmt.Printf("\n\n%+v", "Encontrado Protocolo em aberto")

			} else {
				fmt.Println("\n\n", "Esta tudo bem por aqui")

			}
			return
		}

	}

}

func main() {

	qualquernome()
	//var conta string
	//fmt.Println("Insira a conta")
	//fmt.Scanln(&conta)

	type Responses struct {
		Securities []struct {
			InvestedCapital         float64 `json:"InvestedCapital"`
			NetValue                float64 `json:"NetValue"`
			GrossValue              float64 `json:"GrossValue"`
			Quantity                float64 `json:"Quantity"`
			Maturity                string  `json:"Maturity"`
			Ir                      float64 `json:"Ir"`
			Iof                     float64 `json:"Iof"`
			OtherFee                float64 `json:"OtherFee"`
			TotalTaxes              float64 `json:"TotalTaxes"`
			Index                   string  `json:"Index"`
			Settlement              bool    `json:"Settlement"`
			DailySettlement         bool    `json:"DailySettlement"`
			FixedIncomeSecurityType string  `json:"FixedIncomeSecurityType"`
			Nickname                string  `json:"Nickname"`
			SecurityTypeName        string  `json:"SecurityTypeName"`
			FinancialValueCurrent   float64 `json:"FinancialValueCurrent"`
			CustodyID               int     `json:"CustodyId"`
			UnitPrice               float64 `json:"UnitPrice"`
			SettlementType          int     `json:"SettlementType"`
			SymbolID                string  `json:"SymbolId"`
			OrderProcessedAt        string  `json:"OrderProcessedAt"`
			IsInternalOrder         bool    `json:"IsInternalOrder"`
			ProtocolNumber          int     `json:"ProtocolNumber"`
			CurrencyProfitAndLoss   float64 `json:"CurrencyProfitAndLoss"`
			PercentageProfitAndLoss float64 `json:"PercentageProfitAndLoss"`
			ShowProfitAndLoss       bool    `json:"ShowProfitAndLoss"`
		} `json:"Securities"`
		UpdateDate string `json:"UpdateDate"`
		Cache      bool   `json:"Cache"`
	}

}
