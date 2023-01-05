package main

import (
	"encoding/json"
	"fmt"
)

const (
	main_url    = "https://api.kraken.com/"
	public_path = "0/public/"
	public_url  = main_url + public_path
)

type Asset struct {
	Aclass          string      `json:"aclass"`
	Altname         string      `json:"altname"`
	Decimals        int         `json:"decimals"`
	DisplayDecimals int         `json:"display_decimals"`
	CollateralValue json.Number `json:"collateral_value"`
	Status          string      `json:"status"`
}

type KAssets struct {
	Result map[string]Asset `json:"result"`
}

type KTime struct {
	Result struct {
		Unixtime uint   `json:"unixtime"`
		Rfc1123  string `json:"rfc1123"`
	} `json:"result"`
}

type KStatus struct {
	Result struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
	} `json:"result"`
}

func main() {
	kstatus := getKrakenStatus()
	fmt.Println(kstatus.Result.Status)
	fmt.Println("")
	ktime := getKrakenTime()
	fmt.Println(ktime.Result.Unixtime)
	fmt.Println("")
	kassets := getKrakenAssets()
	fmt.Println(kassets.Result)
}
