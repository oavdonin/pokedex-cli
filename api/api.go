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
		fmt.Println("locations initialized!")
		locations = &ResponseLocations{Next: apiEndpoint + "location/"}
	}
	return locations
}

func GetLocations(direction string) error {
	var pagination string
	if direction == "back" && locations.Previous == "" || locations.Previous == "pokedex" {
		return fmt.Errorf("you're at the 0 index go ahead first")
	} else if direction == "back" && locations.Previous != "" {
		pagination = locations.Previous
	} else if direction == "next" {
		pagination = locations.Next
	} else {
		return fmt.Errorf("error: something is fishy with GetLocations argument")
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
