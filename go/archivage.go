package main

import (
	"os"
	"strconv"
	"time"
)

func createFileName() string {
	currentTime := time.Now().Format("2006-01-02-15_04_05")
	fileName := currentTime + ".csv"
	return fileName
}

func createFile(fileName string) {
	fields := "Altname,Wsname,Aclass_base,Base,Aclass_quote,Quote,Pair_decimals,Cost_decimals,Lot_decimals,Lot_multiplier,Fee_volume_currency,Margin_call,Margin_stop,Ordermin,Costmin,Tick_size,Status"
	kAssetPairs := getKrakenAssetPairs()
	for _, v := range kAssetPairs.Result {
		newString := "\n"
		newString += v.Altname + ","
		newString += v.Wsname + ","
		newString += v.Aclass_base + ","
		newString += v.Base + ","
		newString += v.Aclass_quote + ","
		newString += v.Quote + ","
		newString += strconv.Itoa(v.Pair_decimals) + ","
		newString += strconv.Itoa(v.Cost_decimals) + ","
		newString += strconv.Itoa(v.Lot_decimals) + ","
		newString += strconv.Itoa(v.Lot_decimals) + ","
		newString += strconv.Itoa(v.Lot_multiplier) + ","
		newString += strconv.Itoa(v.Lot_multiplier) + ","
		newString += v.Fee_volume_currency + ","
		newString += strconv.Itoa(v.Margin_call) + ","
		newString += strconv.Itoa(v.Margin_stop) + ","
		newString += v.Ordermin + ","
		newString += v.Tick_size + ","
		newString += v.Status
		fields += newString
	}
	err := os.WriteFile("Archive/"+fileName, []byte(fields), 0644)
	if err != nil {
		panic(err)
	}

}
