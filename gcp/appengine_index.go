package gcp

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
	"legislation_analysis/analyze"
	"net/http"
)

func PersistLegislativeItemsToAppEngineIndex(connection ConnectionMetadata, legislativeItems []analyze.LegislationItem) error {
	index, err := search.Open("users")
	if err != nil {
		return err
	}
	for i := range legislativeItems {
		_, err = index.Put(ctx, legislativeItems[i].Text, legislativeItems[i])
		if err != nil {
			return err
		}
	}

	return nil
}
