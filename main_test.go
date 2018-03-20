package main

import "testing"

func TestGetIdByISIN(t *testing.T) {
	id := GetIdByISIN("GB00B3Q8YX99")
	if id != "F00000O4Y5" {
		t.Errorf("Id was incorrect, got: %s, want: %s.", id, "F00000O4Y5")
	}
}
