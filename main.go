package main

import (
	"io/ioutil"
	"legislation_analysis/analyze"
	"legislation_analysis/gcp"
	"os"
	"strconv"
)

func main() {
	legislativeText, err := ioutil.ReadFile("/Users/timfulmer/go/src/legislation_analysis/txt/legislation.txt")
	if err != nil {
		panic(err)
	}
	legislativeItems, err := analyze.LegislativeText(legislativeText)
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}
	var connectionMetadata = gcp.ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"), Port: port,
		User: os.Getenv("POSTGRES_USER"), Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE")}
	err = gcp.PersistLegislativeItems(connectionMetadata, legislativeItems)
	if err != nil {
		panic(err)
	}
}
