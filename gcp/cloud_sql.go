package gcp

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"legislation_analysis/analyze"
	"log"
	"time"
)

func openSqlDb(connection ConnectionMetadata) (*sql.DB, error) {

	// TODO: switch for local v gae

	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connection.Host, connection.Port, connection.User, connection.Password, connection.Database)
	db, err := sql.Open("cloudsqlpostgres", postgresInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//noinspection GoUnhandledErrorResult
	defer db.Close()

	return db, nil
}

func PersistLegislativeItemsToCloudSql(connection ConnectionMetadata, legislativeItems []analyze.LegislationItem) error {
	log.Println("Persisting legislative items")
	startTime := time.Now()

	db, err := openSqlDb(connection)
	if err != nil {
		return err
	}

	// truncate table to get a clean database
	_, err = db.Exec("TRUNCATE TABLE legislative_item")
	if err != nil {
		return err
	}

	var totalCount = 0
	for i := range legislativeItems {
		insert := "INSERT INTO legislative_item (legislative_text,incidence_count,total_incidence_count) VALUES ($1,$2,$3)"
		_, err = db.Exec(insert, legislativeItems[i].Text, legislativeItems[i].Count, totalCount)
		if err != nil {
			return err
		}
	}
	executionTime := time.Since(startTime)
	log.Printf("Persistence took %s\n", executionTime)

	return nil
}

func SearchLegislativeItemsInCloudSql(connection ConnectionMetadata, searchTerm string) ([]analyze.LegislationItem, error) {
	log.Println("Searching legislative items")
	startTime := time.Now()

	db, err := openSqlDb(connection)
	if err != nil {
		return nil, err
	}

	sel := "SELECT legislative_text,incidence_count,total_incidence_count FROM legislative_item WHERE legislative_text LIKE $1"
	rows, err := db.Query(sel, "%"+searchTerm+"%")
	if err != nil {
		return nil, err
	}
	legislationItems := make([]analyze.LegislationItem, 0, 0)
	for rows.Next() {
		var text string
		var count int
		var total int
		if err := rows.Scan(&text, &count, &total); err != nil {
			return nil, err
		}
		legislationItems = append(legislationItems, analyze.LegislationItem{Text: text, Count: count,
			Bills: []string{}, Sponsors: []string{}, TotalCount: total})
	}
	executionTime := time.Since(startTime)
	log.Printf("Searching took %s\n", executionTime)

	return legislationItems, nil
}
