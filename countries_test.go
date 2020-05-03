package govid

import (
	"fmt"
	"testing"
)

func TestGetAlLCountries(t *testing.T) {
	data, err := GetAllCountries()
	if err != nil {
		fmt.Println(err)
	}
	if len(data) < 1 {
		t.Errorf("Expected length of data is %d; want more than 1", len(data))
	}
}
