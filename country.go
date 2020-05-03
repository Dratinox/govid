package govid

import (
	"io/ioutil"
	"net/http"
)

/*
GetDayOneCountryAll returns info with cases for specific country
@param country string
Must be slug from GetAllData or GetAllCountries
*/
func GetDayOneCountryAll(country string) ([]byte, error) {

	url := "https://api.covid19api.com/dayone/country/" + country
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
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

	return body, nil
}
