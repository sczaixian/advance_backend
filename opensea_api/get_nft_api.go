package opensea_api

import (
	"fmt"
	"io"
	"net/http"
)

func GetNFTAPI() {

	url := "https://api.opensea.io/api/v2/chain/abstract/contract/address/nfts/identifier"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", "6098d14358814514a0cae4d5369bb77f")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
