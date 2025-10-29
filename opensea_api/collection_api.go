package opensea_api

import (
	"fmt"
	"io"
	"net/http"
)

func CollectionAPI() {
	//0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb  CryptoPunks
	url := "https://api.opensea.io/api/v2/collections/sigma-omegahub"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", "6098d14358814514a0cae4d5369bb77f")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

//{
//  "collection" : "sigma-omegahub",
//  "name" : "Sigma OmegaHub\uD83E\uDE99",
//  "description" : "Sigma is transcending crypto hubs",
//  "image_url" : "https://i2.seadn.io/bera_chain/baa6e4978c264512a19cc2df0de7bd28/213193d6d4ede87d2c354dae60df30/22213193d6d4ede87d2c354dae60df30.jpeg",
//  "banner_image_url" : "",
//  "owner" : "0x7eb0f203bae286e3a525424247b974a329baad4b",
//  "safelist_status" : "not_requested",
//  "category" : "",
//  "is_disabled" : false,
//  "is_nsfw" : false,
//  "trait_offers_enabled" : false,
//  "collection_offers_enabled" : true,
//  "opensea_url" : "https://opensea.io/collection/sigma-omegahub",
//  "project_url" : "",
//  "wiki_url" : "",
//  "discord_url" : "",
//  "telegram_url" : "",
//  "twitter_username" : null,
//  "instagram_username" : "",
//  "contracts" : [ {
//    "address" : "0xf4131a5a5607e3a4e89f4b4dbb6c831b396d48a1",
//    "chain" : "bera_chain"
//  } ],
//  "editors" : [ "0x7eb0f203bae286e3a525424247b974a329baad4b" ],
//  "fees" : [ {
//    "fee" : 1.0,
//    "recipient" : "0x0000a26b00c1f0df003000390027140000faa719",
//    "required" : true
//  } ],
//  "total_supply" : 1,
//  "created_date" : "2025-10-29",
//  "payment_tokens" : [ ]
//}
