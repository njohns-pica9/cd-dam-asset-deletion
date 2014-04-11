package main

import (
	"fmt"
	"os"
	"encoding/json"
)

var config struct {
	UploadPath string `json:"upload_path"`
	DbConnection string `json:"db_connection"`
	ElasticHost string `json:"elastic_host"`
	ElasticIndex string `json:"elastic_index"`
}

func initCredentials() {
	configFile, err := os.Open("config.json")

	if err != nil {
		fmt.Printf("[Error]: %s \n", err.Error())
		os.Exit(1)
	}

	err = json.NewDecoder(configFile).Decode(&config)

	if err != nil {
		fmt.Printf("[Error]: %s \n", err.Error())
		os.Exit(1)
	}
}

func init() {
	initCredentials();
}

func main() {
	fmt.Println("Connecting to DAM")
	db := makeConnection()

	defer db.Close()

	asset_id := os.Args[1]
	
	fmt.Printf("Deleteing Asset Files [%s]\n", asset_id)

	deleteFiles(asset_id)
	deleteAsset(db, asset_id)
	removeFromElastic(asset_id)
}