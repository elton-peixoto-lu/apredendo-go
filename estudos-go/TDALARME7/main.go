package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"encoding/json"
)

func main() {
	resp2, err := http.Get("https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json")
	if err != nil {
		fmt.Println("No response from request")
		fmt.Printf("\n\n%+v", resp2)
	}
	resp, err := http.Get("https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json")
	if err != nil {
		fmt.Println("No response from request")
		fmt.Printf("\n\n%+v", resp)

	}

	defer resp.Body.Close()

	fmt.Printf("\n\n%+v", resp2)

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))
	fmt.Printf("\n\n%+v" + "Aqui começa outro")
	fmt.Printf("\n\n%+v" + "Aqui começa outro")
	fmt.Printf("\n\n%+v" + "Aqui começa outro")
	body, err = ioutil.ReadAll(resp2.Body) // response body is []byte
	fmt.Println(string(body))

}
