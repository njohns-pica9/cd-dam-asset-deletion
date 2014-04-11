package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "fmt"
)

func makeConnection() *sql.DB {
    db, err := sql.Open("postgres", config.DbConnection)

    if err != nil {
        log.Fatal(err)
    }

    return db
}

func deleteAsset(db *sql.DB, asset_id string) {
    deleteUsageRights(db, asset_id)
    deleteAssetKin(db, asset_id)
    deleteADV(db, asset_id)
    deleteCTA(db, asset_id)

    var rows string

    err := db.QueryRow("DELETE FROM assets WHERE asset_id = $1", asset_id).Scan(&rows)
    
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No asset with that ID.")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf("Deleted %s row \n", rows)
    }
}

func deleteUsageRights(db *sql.DB, asset_id string) {
    var rows string

    err := db.QueryRow("DELETE FROM usage_rights WHERE asset_id = $1", asset_id).Scan(&rows)
    
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No usage_rights with that Asset ID.")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf("Deleted %s row \n", rows)
    }
}

func deleteADV(db *sql.DB, asset_id string) {
    var rows string

    err := db.QueryRow("DELETE FROM assets_datapoints_values WHERE asset_id = $1", asset_id).Scan(&rows)
    
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No assets_datapoints_values with that Asset ID.")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf("Deleted %s rows \n", rows)
    }
}

func deleteAssetKin(db *sql.DB, asset_id string) {
    var rows string

    err := db.QueryRow("DELETE FROM assets_kin WHERE asset_id = $1", asset_id).Scan(&rows)
    
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No assets_kin with that Asset ID.")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf("Deleted %s row \n", rows)
    }
}

func deleteCTA(db *sql.DB, asset_id string) {
    var rows string

    err := db.QueryRow("DELETE FROM container_tags_assets WHERE asset_id = $1", asset_id).Scan(&rows)
    
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No container_tags_assets with that Asset ID.")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf("Deleted %s rows \n", rows)
    }
}
