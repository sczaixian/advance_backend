package opensea_api

import (
	"fmt"
	"io"
	"net/http"
)

func API_Call(url, apiKey string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
