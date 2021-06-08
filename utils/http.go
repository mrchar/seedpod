package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalResponse(response *http.Response, v interface{}) error {
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}

	return nil
}
