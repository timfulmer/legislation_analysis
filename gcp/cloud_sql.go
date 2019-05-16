package gcp

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"legislation_analysis/analyze"
	"log"
	"time"
)

type ConnectionMetadata struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func PersistLegislativeItems(connection ConnectionMetadata, legislativeItems []analyze.LegislationItem) error {
	log.Println("Persisting legislative items")
	startTime := time.Now()
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connection.Host, connection.Port, connection.User, connection.Password, connection.Database)
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
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
		totalCount += legislativeItems[i].Count
	}
	for i := range legislativeItems {
		insert := "INSERT INTO legislative_item (legislative_text,incidence_count,total_incidence_count) VALUES ($1,$2,$3)"
		_, err = db.Exec(insert, legislativeItems[i].Text, legislativeItems[i].Count, totalCount)
		if err != nil {
			return err
		}
	}

	err = db.Close()
	if err != nil {
		return err
	}
	executionTime := time.Since(startTime)
	log.Printf("Persistence took %s\n", executionTime)

	return nil
}

func SearchLegislativeItems(connection ConnectionMetadata, searchTerm string) ([]analyze.LegislationItem, error) {
	log.Println("Searching legislative items")
	startTime := time.Now()
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connection.Host, connection.Port, connection.User, connection.Password, connection.Database)
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
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
		legislationItems = append(legislationItems, analyze.LegislationItem{text, count, []string{}, []string{}, total})
	}
	executionTime := time.Since(startTime)
	log.Printf("Searching took %s\n", executionTime)

	return legislationItems, nil
}
