package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPublicContent(page string) []byte {
	mylink := public_url + page
	fmt.Println(mylink)
	response, err := http.Get(mylink)
	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	}

	return body
}

func getKrakenTime() *KTime {
	body := getPublicContent("Time")

	var result KTime
	json.Unmarshal(body, &result)
	return &result
}

func getKrakenStatus() *KStatus {
	body := getPublicContent("SystemStatus")

	var result KStatus
	json.Unmarshal(body, &result)
	return &result
}

func getKrakenAssets() *KAssets {
	body := getPublicContent("Assets")

	var result KAssets
	json.Unmarshal(body, &result)

	return &result
}

func getKrakenAssetPairs() *KAssetPairs {
	body := getPublicContent("AssetPairs")

	var result KAssetPairs
	json.Unmarshal(body, &result)

	return &result
}
