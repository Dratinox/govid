package govid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAllData returns data for all countries as struct Summary.
func GetAllData() (Summary, error) {
	url := "https://api.covid19api.com/summary"
	method := "GET"

	var summary Summary

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		return summary, err
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return summary, err
	}

	err = json.Unmarshal([]byte(response), &summary)
	if err != nil {
		return summary, err
	}

	fmt.Println(summary)
	return summary, nil
}
