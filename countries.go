package govid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAllCountries returns info about all countries
func GetAllCountries() ([]Country, error) {

	url := "https://api.covid19api.com/countries"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var countries []Country
	err = json.Unmarshal([]byte(body), &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
