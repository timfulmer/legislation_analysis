package analyze

import (
	"strings"
	"testing"
)

func TestLegislativeText(t *testing.T) {
	legislationItems, err := LegislativeText([]byte("Whereas following the purchases of the Lordship of Schel­len­berg and the County of Vaduz in 1699 and 1712, respectively, on January 23, 1719, the Holy Roman Emperor Charles VI united the two jurisdictions, elevating them to the rank of Reichsfürstentum (Imperial Principality), establishing the borders and territory that still exist today"))
	if err != nil {
		t.Error("Received unexpected error")
	}
	if len(legislationItems) != 5 {
		t.Error("Received incorrect number of legislation items")
	}
}

func TestLegislativeText_multiple(t *testing.T) {
	legislationItems, err := LegislativeText([]byte("Whereas following the purchases of the Lordship of Schel­len­berg and the County of Vaduz in 1699 and 1712, respectively, on January 23, 1719, the Holy Roman Emperor Charles VI united the two jurisdictions, elevating them to the rank of Reichsfürstentum (Imperial Principality), establishing the borders and territory that still exist today Whereas following the purchases of the Lordship of Schel­len­berg and the County of Vaduz in 1699 and 1712, respectively, on January 23, 1719, the Holy Roman Emperor Charles VI united the two jurisdictions, elevating them to the rank of Reichsfürstentum (Imperial Principality), establishing the borders and territory that still exist today"))
	if err != nil {
		t.Error("Received unexpected error")
	}
	if len(legislationItems) != 5 {
		t.Error("Received incorrect number of legislation items")
	}
	if legislationItems[0].Count != 2 {
		t.Error("Received incorrect count of legislative item")
	}
}

func TestLegislativeText_stopWords(t *testing.T) {
	legislationItems, err := LegislativeText([]byte("That a collection of the rules of the committees of the Senate, together with related materials, be printed as a Senate document, and that there be printed 250 additional copies of such document for the use of the Committee on Rules and Administration."))
	if err != nil {
		t.Error("Received unexpected error")
	}
	if len(legislationItems) != 2 {
		t.Error("Received incorrect number of legislation items")
	}
	for i := range legislationItems {
		if legislationItems[i].Text == "Committee" {
			t.Error("Received stop word in results")
		}
	}
}

func TestLegislativeText_stopCharacters(t *testing.T) {
	legislationItems, err := LegislativeText([]byte("the net earning from self-employment reduced by the excess (if any) of subparagraph (A)(i) over subparagraph (A)(ii), over"))
	if err != nil {
		t.Error("Received unexpected error")
	}
	if len(legislationItems) != 0 {
		t.Error("Received incorrect number of legislation items")
	}
	for i := range legislationItems {
		if strings.Index(legislationItems[i].Text, "(") > -1 {
			t.Error("Received stop character in results")
		}
	}
}
