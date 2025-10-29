package opensea_api

import (
	"fmt"
	"io"
	"net/http"
)

func AccountAPI() {

	url := "https://api.opensea.io/api/v2/accounts/0x3E0bDb54f94D735dDCf8D2074c852a8C22914aA7"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", "6098d14358814514a0cae4d5369bb77f")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

// {
//  "address" : "0x3e0bdb54f94d735ddcf8d2074c852a8c22914aa7",
//  "username" : "balabala_big",
//  "profile_image_url" : "",
//  "banner_image_url" : "",
//  "website" : "",
//  "social_media_accounts" : [ ],
//  "bio" : "",
//  "joined_date" : "1970-01-01"
//}
