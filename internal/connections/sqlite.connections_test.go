package connections

import "testing"

func TestGetConnetion(t *testing.T) {
	db := GetConnetion()

	if db == nil {
		t.Fatal("db no puede ser nil")
	}
}
