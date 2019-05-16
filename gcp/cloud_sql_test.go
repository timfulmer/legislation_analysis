package gcp

import (
	"legislation_analysis/analyze"
	"os"
	"strconv"
	"testing"
)

func TestPersistLegislativeItems(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		t.Error(err)
	}
	err = PersistLegislativeItems(ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"),
		Port: port, User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DATABASE")},
		[]analyze.LegislationItem{{"test", 1, []string{}, []string{}, 1}})
	if err != nil {
		t.Error(err)
	}
}

func TestPersistLegislativeItems_BadCredentials(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		t.Error(err)
	}
	err = PersistLegislativeItems(ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"),
		Port: port, User: "dummy-user",
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DATABASE")},
		[]analyze.LegislationItem{{"test", 1, []string{}, []string{}, 1}})
	if err == nil {
		t.Error("Could connect with bad credentials")
	}
}

func TestSearchLegislativeItems(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		t.Error(err)
	}
	var connection = ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"),
		Port: port, User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DATABASE")}
	err = PersistLegislativeItems(connection, []analyze.LegislationItem{{"test", 1, []string{},
		[]string{}, 1}})
	if err != nil {
		t.Error(err)
	}
	legislativeItems, err := SearchLegislativeItems(connection, "test")
	if err != nil {
		t.Error(err)
	}
	if len(legislativeItems) != 1 {
		t.Error("Received incorrect number of legislative items")
	}
}

func TestSearchLegislativeItems_noMatches(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		t.Error(err)
	}
	var connection = ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"),
		Port: port, User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DATABASE")}
	err = PersistLegislativeItems(connection, []analyze.LegislationItem{{"dummy", 1, []string{},
		[]string{}, 1}})
	if err != nil {
		t.Error(err)
	}
	legislativeItems, err := SearchLegislativeItems(connection, "test")
	if err != nil {
		t.Error(err)
	}
	if len(legislativeItems) != 0 {
		t.Error("Received incorrect number of legislative items")
	}
}

func TestSearchLegislativeItems_multipleMatches(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		t.Error(err)
	}
	var connection = ConnectionMetadata{Host: os.Getenv("POSTGRES_HOST"),
		Port: port, User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DATABASE")}
	err = PersistLegislativeItems(connection, []analyze.LegislationItem{{"test1", 1, []string{},
		[]string{}, 1}, {"test2", 1, []string{},
		[]string{}, 1}})
	if err != nil {
		t.Error(err)
	}
	legislativeItems, err := SearchLegislativeItems(connection, "test")
	if err != nil {
		t.Error(err)
	}
	if len(legislativeItems) != 2 {
		t.Error("Received incorrect number of legislative items")
	}
}
