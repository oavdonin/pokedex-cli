package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiEndpoint string = "https://pokeapi.co/api/v2/"

var locations *ResponseLocations

func InitializeLocations() *ResponseLocations {
	if locations == nil {
		locations = &ResponseLocations{Next: apiEndpoint + "location/"}
	}
	return locations
}

func GetLocations(direction ...string) error {
	var pagination string
	if len(direction) == 0 {
		pagination = locations.Next
	} else if direction[0] == "down" && locations.Previous != "" {
		pagination = locations.Previous
	} else {
		return fmt.Errorf("GetLocations was called with a wrong")
	}
	req, err := http.NewRequest("GET", pagination, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return err
	}
	for _, item := range locations.Results {
		fmt.Println(item.Name)
	}
	return nil
}
