package main

import (
	"fmt"
)

const (
	main_url    = "https://api.kraken.com/"
	public_path = "0/public/"
	public_url  = main_url + public_path
)

func main() {
	kstatus := getKrakenStatus()
	fmt.Println(kstatus.Result.Status)
	fmt.Println("")
	ktime := getKrakenTime()
	fmt.Println(ktime.Result.Unixtime)
	fmt.Println("")
	kassets := getKrakenAssets()
	fmt.Println(kassets.Result)

	kassetPairs := getKrakenAssetPairs()
	fmt.Println(kassetPairs.Result)

	fileName := createFileName()
	createFile(fileName)

	startServer()
}
