package main

import (
    "fmt"
    "net/http"
    "log"
)

func removeFromElastic(asset_id string) {
    var asset = fmt.Sprintf("%s/%s/asset/%s", config.ElasticHost, config.ElasticIndex, asset_id)

    req, err := http.NewRequest("DELETE", asset, nil)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Response was: ", resp.Status)
}
