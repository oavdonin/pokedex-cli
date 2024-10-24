package api

import (
	"encoding/json"
	"net/http"
)

const apiEndpoint string = "https://pokeapi.co/api/v2/"

func GetLocations() (ResponseLocations, error) {
	functionalEndpoint := apiEndpoint + "location/"
	req, err := http.NewRequest("GET", functionalEndpoint, nil)
	if err != nil {
		return ResponseLocations{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer resp.Body.Close()
	locationsResponse := ResponseLocations{}
	err = json.NewDecoder(resp.Body).Decode(&locationsResponse)
	if err != nil {
		return ResponseLocations{}, err
	}

	return locationsResponse, nil
}
