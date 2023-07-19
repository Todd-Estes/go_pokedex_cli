package main

import (
	"fmt"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
)

type Locations struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}



func getPokeLocations(loc *Locations) (string, error) {
	res, err := http.Get(loc.Next)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("empty name")
	} 
	defer res.Body.Close()
	body, erruh := ioutil.ReadAll(res.Body)
	if erruh != nil {
		fmt.Println(err)
		return "", errors.New("erruh error")
	}
	errah := json.Unmarshal(body, &loc)
	if errah != nil {
		fmt.Println(errah)
	}
	for _, result := range loc.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return "Success!!", nil
}

func getPrevPokeLocations(loc *Locations) (string, error) {
	res, err := http.Get(loc.Previous)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("empty name")
	} 
	defer res.Body.Close()
	body, erruh := ioutil.ReadAll(res.Body)
	if erruh != nil {
		fmt.Println(err)
		return "", errors.New("erruh error")
	}
	errah := json.Unmarshal(body, &loc)
	if errah != nil {
		fmt.Println(errah)
	}
	for _, result := range loc.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return "Success!!", nil
}