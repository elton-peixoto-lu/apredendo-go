package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"encoding/json"
)

func main() {

	resp2, err := http.Get("http://symbolbase.kube.prod.easynvest.io/SymbolCollection/25/detail/true")
	if err != nil {
		fmt.Println("No response from request")

	}

	fmt.Printf("\n\n%+v", resp2)

	resp, err := http.Get("https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json")
	if err != nil {
		fmt.Println("No response from request")
		//struc mercado
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

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var respostas2 MyJsonName
	json.Unmarshal(body, &respostas2)
	//	fmt.Printf("\n\n%+v", respostas2.Response.TrsrBdTradgList[1].TrsrBd.MinInvstmtAmt)
	///sctruc symbol
	type Resposta2 struct {
		Date   string `json:"date"`
		ID     string `json:"id"`
		Result struct {
			Category struct {
				BackgroundColor string `json:"backgroundColor"`
				CategoryID      int64  `json:"categoryId"`
				CategoryName    string `json:"categoryName"`
				CategoryOrder   int64  `json:"categoryOrder"`
				CreationDate    string `json:"creationDate"`
			} `json:"category"`
			ChannelID       interface{}   `json:"channelId"`
			CollectionID    int64         `json:"collectionId"`
			CollectionOrder int64         `json:"collectionOrder"`
			Description     string        `json:"description"`
			Errors          []interface{} `json:"errors"`
			ExternalContent interface{}   `json:"externalContent"`
			Filters         []struct {
				IDFilter string `json:"idFilter"`
			} `json:"filters"`
			HasError    bool `json:"hasError"`
			IsAutomatic bool `json:"isAutomatic"`
			Medias      []struct {
				CreationDate string `json:"creationDate"`
				Description  string `json:"description"`
				MediaID      int64  `json:"mediaId"`
				MediaOrder   int64  `json:"mediaOrder"`
				Source       string `json:"source"`
				Title        string `json:"title"`
				Type         struct {
					MediaTypeID int64  `json:"mediaTypeId"`
					Name        string `json:"name"`
				} `json:"type"`
			} `json:"medias"`
			Name                        string      `json:"name"`
			ProcessHourLimit            interface{} `json:"processHourLimit"`
			ProcessHourLimitDescription interface{} `json:"processHourLimitDescription"`
			Products                    []struct {
				CreationDate string `json:"creationDate"`
				Name         string `json:"name"`
				ProductID    int64  `json:"productId"`
			} `json:"products"`
			SubTitle              interface{} `json:"subTitle"`
			SymbolCollectionItems []struct {
				AskOffer struct {
					CreationDate string  `json:"creationDate"`
					IsOfferValid bool    `json:"isOfferValid"`
					MinTick      float64 `json:"minTick"`
					OfferType    int64   `json:"offerType"`
					PercentIndex float64 `json:"percentIndex"`
					Price        float64 `json:"price"`
					SymbolID     string  `json:"symbolId"`
					Yield        float64 `json:"yield"`
				} `json:"askOffer"`

				BidOffer                 interface{} `json:"bidOffer"`
				ContractID               int64       `json:"contractId"`
				CouponPayment            string      `json:"couponPayment"`
				CreationDate             string      `json:"creationDate"`
				DaysToExpire             int64       `json:"daysToExpire"`
				ExchangeSymbolID         interface{} `json:"exchangeSymbolId"`
				ExternalID               interface{} `json:"externalId"`
				ExternalName             interface{} `json:"externalName"`
				FixedIncomeRentability   interface{} `json:"fixedIncomeRentability"`
				HasSimulationCalculation bool        `json:"hasSimulationCalculation"`
				IncomeTaxDescription     string      `json:"incomeTaxDescription"`
				IncomeTaxFree            bool        `json:"incomeTaxFree"`
				IncomeTaxRate            float64     `json:"incomeTaxRate"`
				Index                    struct {
					CreationDate string `json:"creationDate"`
					IndexID      string `json:"indexId"`
					IndexType    int64  `json:"indexType"`
					Nickname     string `json:"nickname"`
				} `json:"index"`
				IsGuaranteedByFGC bool `json:"isGuaranteedByFGC"`
				IsOddLot          bool `json:"isOddLot"`
				IsPromotion       bool `json:"isPromotion"`
				Issuer            struct {
					CreationDate string        `json:"creationDate"`
					IssuerID     int64         `json:"issuerId"`
					IssuerName   string        `json:"issuerName"`
					IssuerRating []interface{} `json:"issuerRating"`
				} `json:"issuer"`
				Liquidity           int64  `json:"liquidity"`
				LiquidityName       string `json:"liquidityName"`
				MarketType          int64  `json:"marketType"`
				MaturityDate        string `json:"maturityDate"`
				MaturityDescription string `json:"maturityDescription"`
				Name                string `json:"name"`
				Product             struct {
					CreationDate string `json:"creationDate"`
					Name         string `json:"name"`
					ProductID    int64  `json:"productId"`
				} `json:"product"`
				Qualified              bool    `json:"qualified"`
				RelativeRentability    float64 `json:"relativeRentability"`
				RentabilityDescription string  `json:"rentabilityDescription"`
				RentabilityType        int64   `json:"rentabilityType"`
				SecurityTypeID         string  `json:"securityTypeId"`
				SecurityTypeName       string  `json:"securityTypeName"`
				SimulatedRentability   struct {
					AnnualGrossRateProfit  float64 `json:"annualGrossRateProfit"`
					AnnualNetRateProfit    float64 `json:"annualNetRateProfit"`
					DailyGrossRateProfit   float64 `json:"dailyGrossRateProfit"`
					GrossAmount            float64 `json:"grossAmount"`
					GrossAmountProfit      float64 `json:"grossAmountProfit"`
					MonthlyGrossRateProfit float64 `json:"monthlyGrossRateProfit"`
					NetAmount              float64 `json:"netAmount"`
					NetAmountProfit        float64 `json:"netAmountProfit"`
					RateProfit             float64 `json:"rateProfit"`
					TaxesAmount            float64 `json:"taxesAmount"`
					TaxesRate              float64 `json:"taxesRate"`
				} `json:"simulatedRentability"`
				StockBalance         interface{}   `json:"stockBalance"`
				SuitabilityProfileID int64         `json:"suitabilityProfileId"`
				SymbolID             string        `json:"symbolId"`
				SymbolRating         []interface{} `json:"symbolRating"`
				Valid                bool          `json:"valid"`
			} `json:"symbolCollectionItems"`
			URLImage interface{} `json:"urlImage"`
		} `json:"result"`
		Status struct {
			Messages []string `json:"messages"`
			Success  bool     `json:"success"`
		} `json:"status"`
	}

	defer resp2.Body.Close()

	body, err = ioutil.ReadAll(resp2.Body) // response body is []byte

	var respostas Resposta2

	//var respostas2 MyJsonName

	json.Unmarshal(body, &respostas)
	//json.Unmarshal(body, &respostas2)

	//erro = json.Unmarshal(body, &MyJsonName)
	fmt.Printf("\n\n\n")
	fmt.Printf("\n\n%+v", respostas.Status)

	if respostas2.ResponseStatus == 200 {
		fmt.Printf("\n\n\n Mercado Disponivel")
	}

	for i := 0; i < len(respostas.Result.SymbolCollectionItems); i++ {
		for y := 0; y < len(respostas2.Response.TrsrBdTradgList); y++ {
			fmt.Printf("\n\n\n")
			//fmt.Printf("\n\n%+v", respostas2.Response.TrsrBdTradgList[y].TrsrBd.Nm)
			//fmt.Printf("\n\n Valor Da nossa PLATAFORMA")

			fmt.Printf("\n\n\n")
			fmt.Printf("\n\n\n Verificando valores")
			fmt.Printf("\n\n%+v", respostas.Result.SymbolCollectionItems[i].Name)
			fmt.Printf("\n\n%+v", respostas.Result.SymbolCollectionItems[i].AskOffer.MinTick)
			fmt.Printf("\n\n%+v", respostas2.Response.TrsrBdTradgList[y].TrsrBd.Nm)
			fmt.Printf("\n\n%+v", respostas2.Response.TrsrBdTradgList[y].TrsrBd.MinInvstmtAmt)
			fmt.Printf("\n\n\n")
			if respostas2.Response.TrsrBdTradgList[y].TrsrBd.MinInvstmtAmt == respostas.Result.SymbolCollectionItems[i].AskOffer.MinTick {
				fmt.Printf("\n\n\n Valores compativeis")

			} else {
				fmt.Printf("Valores incompativeis")
			}

			if respostas2.Response.TrsrBdTradgList[i].TrsrBd.MinInvstmtAmt == 0 {
				fmt.Printf("\n\n\n")
				fmt.Printf("Titulo esta indiponivel no momento")

			}
		}

	}

}
