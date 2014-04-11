CD Dam Asset Deleter.
===============

I couldn't think of a better name.

## Usage Requirements
  * Nothing, static binary!

## Build Requirements
 * [Golang](http://golang.org)
 * Patience and/or balls of steel

## Build
```bash
git clone git@github.com:njohns-pica9/cd-dam-asset-deletion.git deleter
cd deleter
go build
```

## Config
```json
{
    "db_connection": "postgres://username:password@hostname/dbname?sslmode=disable",
    "upload_path": "/var/www/data",
    "elastic_host": "http://localhost:9200",
    "elastic_index": "dam3"
}

```