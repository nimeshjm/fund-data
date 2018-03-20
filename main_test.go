package main

import (
	"testing"
	"fmt"
)

func TestGetIdByISIN(t *testing.T) {
	id := GetIdByISIN("GB00B3Q8YX99")
	if id != "F00000O4Y5" {
		t.Errorf("Id was incorrect, got: %s, want: %s.", id, "F00000O4Y5")
	}
}

func TestGetNavPrice(t *testing.T) {
	price, currency, date := GetNavPrice("F00000OXIA")
	if price == 0 || currency == "" || date == "" {
		t.Errorf("Id was incorrect, got: %d, want: %s.", 0, "!= 0")
	}

	fmt.Println(price, currency, date)
	}

