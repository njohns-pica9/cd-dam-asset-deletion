package main

import (
    "os"
    "fmt"
)

func deleteFiles(asset_id string) {
    directories := []string{"assets", "pops", "thumbs", "usage_rights"};

    dep := []byte(asset_id)
    lastIndex := string(dep[len(dep)-1:])

    for _, v := range directories {
        var file = fmt.Sprintf("%s/%s/%s/%s", config.UploadPath, v, lastIndex, asset_id)
        os.Remove(file)
    }
}
