package alchemy_api

import (
	"fmt"
	"io"
	"net/http"
)

//https://www.alchemy.com/docs/reference/nft-api-quickstart

func AlchemyAPI() {
	//${baseURL}/getNFTMetadata?contractAddress=${contractAddress}&tokenId=${tokenId}
	contractAddress := "0x0000000000cf80E7Cf8Fa4480907f692177f8e06"
	tokenId := "73906452355594127029039375271145516945927406532858726769026903911185640775143"
	url := "https://eth-mainnet.g.alchemy.com/nft/v3/RmsPYhly5O6-XH8UdmqCQ/getNFTMetadata?contractAddress=%s&tokenId=%s"
	req, _ := http.NewRequest("GET", fmt.Sprintf(url, contractAddress, tokenId), nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

//{
//	"contract": {
//		"address": "0x0000000000cf80E7Cf8Fa4480907f692177f8e06",
//		"name": "NamefiNFT",
//		"symbol": "NFNFT",
//		"totalSupply": null,
//		"tokenType": "ERC721",
//		"contractDeployer": "0x1b0f291c8fFebE891886351CDfF8A304a840C8Ad",
//		"deployedBlockNumber": 19059948,
//		"openSeaMetadata": {
//			"floorPrice": 0.003,
//			"collectionName": "NamefiNFT",
//			"collectionSlug": "namefinft",
//			"safelistRequestStatus": "not_requested",
//			"imageUrl": "https://raw2.seadn.io/ethereum/0x0000000000cf80e7cf8fa4480907f692177f8e06/af1a8977e12ec6b05cd5759671bc6396.svg",
//			"description": "Namefi.io tokenize DNS domain name ownership on Ethereum.",
//			"externalUrl": null,
//			"twitterUsername": null,
//			"discordUrl": null,
//			"bannerImageUrl": null,
//			"lastIngestedAt": "2025-10-25T05:09:28.000Z"
//		},
//		"isSpam": true,
//		"spamClassifications": ["SuspiciousMetadata", "SpammyMetadata"]
//	},
//	"tokenId": "73906452355594127029039375271145516945927406532858726769026903911185640775143",
//	"tokenType": "ERC721",
//	"name": "vitalik.cloud",
//	"description": "vitalik.cloud - Namefi™️ NFT representing the beneficiary-ship of vitalik.cloud domain. vitalik.cloud is valuable because \"Vitalik\" resonates with Vitalik Buterin, co-founder of Ethereum, linking to technology and innovation. The \".cloud\" TLD emphasizes modern tech, cloud computing, and online solutions, relevant to tech startups or services. This combination could attract businesses in the tech sector or enthusiasts of blockchain technology, increasing its potential for branding and marketability as a domain associated with leading-edge technological solutions and thought leadership in the cloud computing space.",
//	"tokenUri": "https://md.namefi.io/vitalik.cloud",
//	"image": {
//		"cachedUrl": "https://nft-cdn.alchemy.com/eth-mainnet/a91d6d9cffbe5426ba8b50de8ced5868",
//		"thumbnailUrl": "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/a91d6d9cffbe5426ba8b50de8ced5868",
//		"pngUrl": "https://res.cloudinary.com/alchemyapi/image/upload/convert-png/eth-mainnet/a91d6d9cffbe5426ba8b50de8ced5868",
//		"contentType": "image/svg+xml; charset=utf-8",
//		"size": 15116,
//		"originalUrl": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg"
//	},
//	"animation": {
//		"cachedUrl": null,
//		"contentType": null,
//		"size": null,
//		"originalUrl": null
//	},
//	"raw": {
//		"tokenUri": "https://md.namefi.io/vitalik.cloud",
//		"metadata": {
//			"image": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg",
//			"external_url": "https://vitalik.cloud",
//			"is_normalized": true,
//			"image_url": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg",
//			"name": "vitalik.cloud",
//			"description": "vitalik.cloud - Namefi™️ NFT representing the beneficiary-ship of vitalik.cloud domain. vitalik.cloud is valuable because \"Vitalik\" resonates with Vitalik Buterin, co-founder of Ethereum, linking to technology and innovation. The \".cloud\" TLD emphasizes modern tech, cloud computing, and online solutions, relevant to tech startups or services. This combination could attract businesses in the tech sector or enthusiasts of blockchain technology, increasing its potential for branding and marketability as a domain associated with leading-edge technological solutions and thought leadership in the cloud computing space.",
//			"attributes": [{
//				"value": "🔓 Unlocked",
//				"trait_type": "Is Locked"
//			}, {
//				"value": false,
//				"trait_type": "Is Frozen"
//			}, {
//				"value": "cloud",
//				"trait_type": "Top Level Domain (TLD)"
//			}, {
//				"display_type": "number",
//				"value": 5,
//				"trait_type": "TLD Length"
//			}, {
//				"value": "vitalik",
//				"trait_type": "Second Level Domain (TLD)"
//			}, {
//				"display_type": "number",
//				"value": 7,
//				"trait_type": "SLD Length"
//			}, {
//				"display_type": "date",
//				"value": "2026-10-11",
//				"trait_type": "Expiration Date"
//			}],
//			"version": 0,
//			"url": "https://md.namefi.io/ethereum/vitalik.cloud",
//			"_extension": {
//				"currentOwnerLastUpdated": "2025-10-08T21:02:50.760Z",
//				"currentOwner": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
//			}
//		},
//		"error": null
//	},
//	"collection": {
//		"name": "NamefiNFT",
//		"slug": "namefinft",
//		"externalUrl": null,
//		"bannerImageUrl": null
//	},
//	"mint": {
//		"mintAddress": null,
//		"blockNumber": null,
//		"timestamp": null,
//		"transactionHash": null
//	},
//	"owners": null,
//	"timeLastUpdated": "2025-10-29T06:42:03.451Z"
//}

func API_Call() {
	//owner := "0x7ae58cd55a3466cd8785dfc2ea3c870fab8a625c"
	//baseURL := "https://eth-mainnet.g.alchemy.com/nft/v3/%s"
	owner := "vitalik.eth"
	apiKey := "RmsPYhly5O6-XH8UdmqCQ"
	url := "https://eth-mainnet.g.alchemy.com/nft/v3/%s/getNFTsForOwner?owner=%s&pageSize=5"
	req, _ := http.NewRequest("GET", fmt.Sprintf(url, apiKey, owner), nil)
	req.Header.Add("accept", "application/json")
	//req.Header.Add("x-api-key", apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

//url := "https://eth-mainnet.g.alchemy.com/nft/v3/%s/getNFTsForOwner?owner=%s&pageSize=5"
//{
//	"ownedNfts":[],
//	"totalCount":0,
//	"validAt":{
//		"blockNumber":23680915,
//		"blockHash":"0xaa93eec73f71875c44832fe9ba1bbef2136ce46da0524dbe54c8fb28b510014b",
//		"blockTimestamp":"2025-10-29T05:17:59Z"
//	},
//	"pageKey":null
//}

//{
//	"ownedNfts": [{
//		"contract": {
//			"address": "0x0000000000cf80E7Cf8Fa4480907f692177f8e06",
//			"name": "NamefiNFT",
//			"symbol": "NFNFT",
//			"totalSupply": null,
//			"tokenType": "ERC721",
//			"contractDeployer": "0x1b0f291c8fFebE891886351CDfF8A304a840C8Ad",
//			"deployedBlockNumber": 19059948,
//			"openSeaMetadata": {
//				"floorPrice": 0.003,
//				"collectionName": "NamefiNFT",
//				"collectionSlug": "namefinft",
//				"safelistRequestStatus": "not_requested",
//				"imageUrl": "https://raw2.seadn.io/ethereum/0x0000000000cf80e7cf8fa4480907f692177f8e06/af1a8977e12ec6b05cd5759671bc6396.svg",
//				"description": "Namefi.io tokenize DNS domain name ownership on Ethereum.",
//				"externalUrl": null,
//				"twitterUsername": null,
//				"discordUrl": null,
//				"bannerImageUrl": null,
//				"lastIngestedAt": "2025-10-25T05:09:28.000Z"
//			},
//			"isSpam": true,
//			"spamClassifications": ["SuspiciousMetadata", "SpammyMetadata"]
//		},
//		"tokenId": "73906452355594127029039375271145516945927406532858726769026903911185640775143",
//		"tokenType": "ERC721",
//		"name": "vitalik.cloud",
//		"description": "vitalik.cloud - Namefi™️ NFT representing the beneficiary-ship of vitalik.cloud domain. vitalik.cloud is valuable because \"Vitalik\" resonates with Vitalik Buterin, co-founder of Ethereum, linking to technology and innovation. The \".cloud\" TLD emphasizes modern tech, cloud computing, and online solutions, relevant to tech startups or services. This combination could attract businesses in the tech sector or enthusiasts of blockchain technology, increasing its potential for branding and marketability as a domain associated with leading-edge technological solutions and thought leadership in the cloud computing space.",
//		"tokenUri": "https://md.namefi.io/vitalik.cloud",
//		"image": {
//			"cachedUrl": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg",
//			"thumbnailUrl": null,
//			"pngUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg"
//		},
//		"animation": {
//			"cachedUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"raw": {
//			"tokenUri": "https://md.namefi.io/vitalik.cloud",
//			"metadata": {
//				"image": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg",
//				"external_url": "https://vitalik.cloud",
//				"is_normalized": true,
//				"image_url": "https://md.namefi.io/ethereum/svg/vitalik.cloud/image.svg",
//				"name": "vitalik.cloud",
//				"description": "vitalik.cloud - Namefi™️ NFT representing the beneficiary-ship of vitalik.cloud domain. vitalik.cloud is valuable because \"Vitalik\" resonates with Vitalik Buterin, co-founder of Ethereum, linking to technology and innovation. The \".cloud\" TLD emphasizes modern tech, cloud computing, and online solutions, relevant to tech startups or services. This combination could attract businesses in the tech sector or enthusiasts of blockchain technology, increasing its potential for branding and marketability as a domain associated with leading-edge technological solutions and thought leadership in the cloud computing space.",
//				"attributes": [{
//					"value": "🔓 Unlocked",
//					"trait_type": "Is Locked"
//				}, {
//					"value": false,
//					"trait_type": "Is Frozen"
//				}, {
//					"value": "cloud",
//					"trait_type": "Top Level Domain (TLD)"
//				}, {
//					"display_type": "number",
//					"value": 5,
//					"trait_type": "TLD Length"
//				}, {
//					"value": "vitalik",
//					"trait_type": "Second Level Domain (TLD)"
//				}, {
//					"display_type": "number",
//					"value": 7,
//					"trait_type": "SLD Length"
//				}, {
//					"display_type": "date",
//					"value": "2026-10-11",
//					"trait_type": "Expiration Date"
//				}],
//				"version": 0,
//				"url": "https://md.namefi.io/ethereum/vitalik.cloud",
//				"_extension": {
//					"currentOwnerLastUpdated": "2025-10-08T21:02:50.760Z",
//					"currentOwner": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
//				}
//			},
//			"error": null
//		},
//		"collection": {
//			"name": "NamefiNFT",
//			"slug": "namefinft",
//			"externalUrl": null,
//			"bannerImageUrl": null
//		},
//		"mint": {
//			"mintAddress": null,
//			"blockNumber": null,
//			"timestamp": null,
//			"transactionHash": null
//		},
//		"owners": null,
//		"timeLastUpdated": "2025-10-29T06:19:13.969Z",
//		"balance": "1",
//		"acquiredAt": {
//			"blockTimestamp": null,
//			"blockNumber": null
//		}
//	}, {
//		"contract": {
//			"address": "0x0000420538CD5AbfBC7Db219B6A1d125f5892Ab0",
//			"name": "Buttpluggy",
//			"symbol": "UwU",
//			"totalSupply": "1024",
//			"tokenType": "ERC721",
//			"contractDeployer": "0xC0FFEc688113B2C5f503dFEAF43548E73C7eCCB3",
//			"deployedBlockNumber": 19242688,
//			"openSeaMetadata": {
//				"floorPrice": 0.0,
//				"collectionName": "Buttpluggy",
//				"collectionSlug": "buttpluggy",
//				"safelistRequestStatus": "not_requested",
//				"imageUrl": "https://i2.seadn.io/ethereum/fc2d4f50929c4f73a9e5d0e784bb4693/46e3fd6dca8132b10f52a23d9babfe/f946e3fd6dca8132b10f52a23d9babfe.png",
//				"description": "**Discover the Future of Gas Efficiency with the First #Huff-Created Collection**\n\nExplore 1024 unique oscilloscope visuals, a breakthrough in CryptoArt by the WebtrES community's Buttplug project. Each piece, a marvel of innovation, is anchored in the Ethereum blockchain, ensuring authentic ownership. This pioneering collection not only showcases artistic finesse but also champions gas efficiency, marking a new era in digital collectibles.\n\n**Join the Evolution**\n\nEmbrace the artistry and technology fusion. Mine your exclusive collectible at [Buttplug Project Homepage](https://www.buttpluggy.com/). To mint one you have to submit a proof of work.\n",
//				"externalUrl": null,
//				"twitterUsername": null,
//				"discordUrl": null,
//				"bannerImageUrl": "https://i2.seadn.io/ethereum/fc2d4f50929c4f73a9e5d0e784bb4693/acfad1b38f190961c3938f1edf7f0a/8aacfad1b38f190961c3938f1edf7f0a.png?fit=inside",
//				"lastIngestedAt": "2025-10-24T16:42:21.000Z"
//			},
//			"isSpam": false,
//			"spamClassifications": []
//		},
//		"tokenId": "1001",
//		"tokenType": "ERC721",
//		"name": "Sergeant WhiskerBlast",
//		"description": "Once a decorated soldier in the Meowtary Forces, Sergeant WhiskerBlast was known for his bravery and indomitable spirit. After a fierce battle that left his screen cracked, he was honorably discharged. Unwilling to retire quietly, he equipped himself with rocket arms and legs, ensuring he could still zoom into action whenever needed. His signature shades, a token from his last mission, hide the scars of war but not the fierce determination in his eyes. He now leads a vigilante group known as the WebtrES Club, defending the digital frontier against cyber threats with his cat-like reflexes and explosive agility.",
//		"tokenUri": "https://buttpluggy.com/final/1001",
//		"image": {
//			"cachedUrl": "https://nft-cdn.alchemy.com/eth-mainnet/3d0532c4cfd14fa15987a5dcb226d522",
//			"thumbnailUrl": "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/3d0532c4cfd14fa15987a5dcb226d522",
//			"pngUrl": "https://res.cloudinary.com/alchemyapi/image/upload/convert-png/eth-mainnet/3d0532c4cfd14fa15987a5dcb226d522",
//			"contentType": "image/webp",
//			"size": 43008,
//			"originalUrl": "https://i2.seadn.io/ethereum/0x0000420538cd5abfbc7db219b6a1d125f5892ab0/cd1be97c3c72c99452bbadb2d58d93/53cd1be97c3c72c99452bbadb2d58d93.gif"
//		},
//		"animation": {
//			"cachedUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"raw": {
//			"tokenUri": "https://buttpluggy.com/final/1001",
//			"metadata": {
//				"name": "Sergeant WhiskerBlast",
//				"description": "Once a decorated soldier in the Meowtary Forces, Sergeant WhiskerBlast was known for his bravery and indomitable spirit. After a fierce battle that left his screen cracked, he was honorably discharged. Unwilling to retire quietly, he equipped himself with rocket arms and legs, ensuring he could still zoom into action whenever needed. His signature shades, a token from his last mission, hide the scars of war but not the fierce determination in his eyes. He now leads a vigilante group known as the WebtrES Club, defending the digital frontier against cyber threats with his cat-like reflexes and explosive agility.",
//				"image": "https://i2.seadn.io/ethereum/0x0000420538cd5abfbc7db219b6a1d125f5892ab0/cd1be97c3c72c99452bbadb2d58d93/53cd1be97c3c72c99452bbadb2d58d93.gif",
//				"attributes": [{
//					"value": "Soldier",
//					"trait_type": "Box"
//				}, {
//					"value": "Shades",
//					"trait_type": "Addon"
//				}, {
//					"value": "Broken",
//					"trait_type": "Screen"
//				}, {
//					"value": "Cat",
//					"trait_type": "Buttons"
//				}, {
//					"value": "Rocket",
//					"trait_type": "Arms and legs"
//				}]
//			},
//			"error": null
//		},
//		"collection": {
//			"name": "Buttpluggy",
//			"slug": "buttpluggy",
//			"externalUrl": null,
//			"bannerImageUrl": "https://i2.seadn.io/ethereum/fc2d4f50929c4f73a9e5d0e784bb4693/acfad1b38f190961c3938f1edf7f0a/8aacfad1b38f190961c3938f1edf7f0a.png?fit=inside"
//		},
//		"mint": {
//			"mintAddress": null,
//			"blockNumber": null,
//			"timestamp": null,
//			"transactionHash": null
//		},
//		"owners": null,
//		"timeLastUpdated": "2025-10-29T06:12:50.698Z",
//		"balance": "1",
//		"acquiredAt": {
//			"blockTimestamp": null,
//			"blockNumber": null
//		}
//	}, {
//		"contract": {
//			"address": "0x000386E3F7559d9B6a2F5c46B4aD1A9587D59Dc3",
//			"name": "Bored Ape Nike Club",
//			"symbol": "BANC",
//			"totalSupply": null,
//			"tokenType": "ERC721",
//			"contractDeployer": "0x51D7D428041E23ef51422e110dfEfF906e821CFe",
//			"deployedBlockNumber": 14276343,
//			"openSeaMetadata": {
//				"floorPrice": 0.0,
//				"collectionName": "BoredApeNikeClub",
//				"collectionSlug": "bored-ape-nike-club-v2",
//				"safelistRequestStatus": "not_requested",
//				"imageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/089a27ca53d629213c2af6d81e74af/87089a27ca53d629213c2af6d81e74af.gif",
//				"description": "COUNTDOWN OVER. MINTING LIVE.\n\n[Mint on the website.](https://nikemetaverse.xyz)\n",
//				"externalUrl": null,
//				"twitterUsername": null,
//				"discordUrl": null,
//				"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside",
//				"lastIngestedAt": "2025-10-23T18:48:14.000Z"
//			},
//			"isSpam": true,
//			"spamClassifications": ["EmptyMetadata", "SpammyMetadata"]
//		},
//		"tokenId": "1",
//		"tokenType": "ERC721",
//		"name": null,
//		"description": null,
//		"tokenUri": "http://api.nikeapenft.xyz/ipfs/1",
//		"image": {
//			"cachedUrl": null,
//			"thumbnailUrl": null,
//			"pngUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"animation": {
//			"cachedUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"raw": {
//			"tokenUri": "http://api.nikeapenft.xyz/ipfs/1",
//			"metadata": {},
//			"error": "Contract returned a broken token uri"
//		},
//		"collection": {
//			"name": "BoredApeNikeClub",
//			"slug": "bored-ape-nike-club-v2",
//			"externalUrl": null,
//			"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside"
//		},
//		"mint": {
//			"mintAddress": null,
//			"blockNumber": null,
//			"timestamp": null,
//			"transactionHash": null
//		},
//		"owners": null,
//		"timeLastUpdated": "2025-10-29T06:40:06.666Z",
//		"balance": "26",
//		"acquiredAt": {
//			"blockTimestamp": null,
//			"blockNumber": null
//		}
//	}, {
//		"contract": {
//			"address": "0x000386E3F7559d9B6a2F5c46B4aD1A9587D59Dc3",
//			"name": "Bored Ape Nike Club",
//			"symbol": "BANC",
//			"totalSupply": null,
//			"tokenType": "ERC721",
//			"contractDeployer": "0x51D7D428041E23ef51422e110dfEfF906e821CFe",
//			"deployedBlockNumber": 14276343,
//			"openSeaMetadata": {
//				"floorPrice": 0.0,
//				"collectionName": "BoredApeNikeClub",
//				"collectionSlug": "bored-ape-nike-club-v2",
//				"safelistRequestStatus": "not_requested",
//				"imageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/089a27ca53d629213c2af6d81e74af/87089a27ca53d629213c2af6d81e74af.gif",
//				"description": "COUNTDOWN OVER. MINTING LIVE.\n\n[Mint on the website.](https://nikemetaverse.xyz)\n",
//				"externalUrl": null,
//				"twitterUsername": null,
//				"discordUrl": null,
//				"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside",
//				"lastIngestedAt": "2025-10-23T18:48:14.000Z"
//			},
//			"isSpam": true,
//			"spamClassifications": ["EmptyMetadata", "SpammyMetadata"]
//		},
//		"tokenId": "2",
//		"tokenType": "ERC721",
//		"name": null,
//		"description": null,
//		"tokenUri": "http://api.nikeapenft.xyz/ipfs/2",
//		"image": {
//			"cachedUrl": null,
//			"thumbnailUrl": null,
//			"pngUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"animation": {
//			"cachedUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"raw": {
//			"tokenUri": "http://api.nikeapenft.xyz/ipfs/2",
//			"metadata": {},
//			"error": "Contract returned a broken token uri"
//		},
//		"collection": {
//			"name": "BoredApeNikeClub",
//			"slug": "bored-ape-nike-club-v2",
//			"externalUrl": null,
//			"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside"
//		},
//		"mint": {
//			"mintAddress": null,
//			"blockNumber": null,
//			"timestamp": null,
//			"transactionHash": null
//		},
//		"owners": null,
//		"timeLastUpdated": "2025-10-29T06:40:02.085Z",
//		"balance": "31",
//		"acquiredAt": {
//			"blockTimestamp": null,
//			"blockNumber": null
//		}
//	}, {
//		"contract": {
//			"address": "0x000386E3F7559d9B6a2F5c46B4aD1A9587D59Dc3",
//			"name": "Bored Ape Nike Club",
//			"symbol": "BANC",
//			"totalSupply": null,
//			"tokenType": "ERC721",
//			"contractDeployer": "0x51D7D428041E23ef51422e110dfEfF906e821CFe",
//			"deployedBlockNumber": 14276343,
//			"openSeaMetadata": {
//				"floorPrice": 0.0,
//				"collectionName": "BoredApeNikeClub",
//				"collectionSlug": "bored-ape-nike-club-v2",
//				"safelistRequestStatus": "not_requested",
//				"imageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/089a27ca53d629213c2af6d81e74af/87089a27ca53d629213c2af6d81e74af.gif",
//				"description": "COUNTDOWN OVER. MINTING LIVE.\n\n[Mint on the website.](https://nikemetaverse.xyz)\n",
//				"externalUrl": null,
//				"twitterUsername": null,
//				"discordUrl": null,
//				"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside",
//				"lastIngestedAt": "2025-10-23T18:48:14.000Z"
//			},
//			"isSpam": true,
//			"spamClassifications": ["EmptyMetadata", "SpammyMetadata"]
//		},
//		"tokenId": "3",
//		"tokenType": "ERC721",
//		"name": null,
//		"description": null,
//		"tokenUri": "http://api.nikeapenft.xyz/ipfs/3",
//		"image": {
//			"cachedUrl": null,
//			"thumbnailUrl": null,
//			"pngUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"animation": {
//			"cachedUrl": null,
//			"contentType": null,
//			"size": null,
//			"originalUrl": null
//		},
//		"raw": {
//			"tokenUri": "http://api.nikeapenft.xyz/ipfs/3",
//			"metadata": {},
//			"error": "Contract returned a broken token uri"
//		},
//		"collection": {
//			"name": "BoredApeNikeClub",
//			"slug": "bored-ape-nike-club-v2",
//			"externalUrl": null,
//			"bannerImageUrl": "https://i2.seadn.io/ethereum/83bf78b80c814fddb778b94dc2357d35/6d235e8782809e7cd9146f53a80c11/b66d235e8782809e7cd9146f53a80c11.png?fit=inside"
//		},
//		"mint": {
//			"mintAddress": null,
//			"blockNumber": null,
//			"timestamp": null,
//			"transactionHash": null
//		},
//		"owners": null,
//		"timeLastUpdated": "2025-10-29T06:39:54.390Z",
//		"balance": "18",
//		"acquiredAt": {
//			"blockTimestamp": null,
//			"blockNumber": null
//		}
//	}],
//	"totalCount": 27599,
//	"validAt": {
//		"blockNumber": 23681326,
//		"blockHash": "0xb973ede18e8dcb89c274e9373a751b5dbb4831e0f26057dda3c601c30b64ed00",
//		"blockTimestamp": "2025-10-29T06:40:23Z"
//	},
//	"pageKey": "MHgwMDAzODZlM2Y3NTU5ZDliNmEyZjVjNDZiNGFkMWE5NTg3ZDU5ZGMzOjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMzpmYWxzZQ=="
//}
