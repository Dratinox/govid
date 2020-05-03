package govid

import (
	"testing"
)

func TestGetAllData(t *testing.T) {
	data, err := GetAllData()
	if err != nil {
		t.Errorf("Got unexpected error %s", err)
	}

	if data.Countries == nil {
		t.Errorf("Expected list of counries with number of cases, got nil")
	}
}
