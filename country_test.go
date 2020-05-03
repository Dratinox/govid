package govid

import (
	"fmt"
	"testing"
)

func TestGetDayOneCountryAll(t *testing.T) {
	data, err := GetDayOneCountryAll("south-africa")
	if err != nil {
		fmt.Println(err)
	}
	if len(data) < 1 {
		t.Errorf("Expected length of data is %d; want more than 1", len(data))
	}
	fmt.Println(string(data))
}
