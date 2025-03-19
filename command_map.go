package main

import (
	"errors"
	"fmt"
	"time"
)

func commandMapf(cfg *config) error {
	start := time.Now()
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	println("Your request took ", time.Since(start).Milliseconds(), " ms")
	return nil
}

func commandMapb(cfg *config) error {
	start := time.Now()
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	println("Your request took ", time.Since(start).Milliseconds(), " ms")
	return nil
}
