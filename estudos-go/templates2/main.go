package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

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

}

//request index page handle
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", "This is the index page!")
}

//request contact page handle
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", "This is the contact page!")
}

//request about page handle
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	//	fmt.Printf("\n\n%+v", respostas2.Response.TrsrBdTradgList[1].TrsrBd.MinInvstmtAmt)

	templates.ExecuteTemplate(w, "about.html", "this")

}
