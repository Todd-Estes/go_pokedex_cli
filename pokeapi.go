package main

import (
	"fmt"
	"net/http"
	"errors"
	"io"
	"encoding/json"
	"github.com/Todd-Estes/internal/pokecache"
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



func getPokeLocations(url string, loc *Locations, c *pokecache.Cache) error {
	_,ok := c.Cache[url]
	if ok {
	results := c.Cache[url].Val
		outputCachedResults(loc, results)
		return nil
	}


	// Get response from API
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return errors.New("empty name")
	} 
	defer res.Body.Close()
	body, erruh := io.ReadAll(res.Body)
	if erruh != nil {
		fmt.Println(err)
		return errors.New("erruh error")
	}

	// Add URL and response to cache
	c.Add(url, body)
	// Maybe we can use outputCachedResults
	outputCachedResults(loc, c.Cache[url].Val)

	//Original code...lets keep this, but we're gonna comment it out for now



	// errah := json.Unmarshal(body, &loc)
	// if errah != nil {
	// 	fmt.Println(errah)
	// }
	// for _, result := range loc.Results {
	// 	fmt.Printf("%s\n", result.Name)
	// }
	// return "Success!!", nil
	return nil
}

func getPrevPokeLocations(loc *Locations, c *pokecache.Cache) error {
		url := loc.Previous
		_, ok := c.Cache[url]
	if ok {
	results := c.Cache[url].Val
		outputCachedResults(loc, results)
		return nil
	} else {
		getPokeLocations(url, loc, c)
		return nil
	}



	// res, err := http.Get(loc.Previous)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", errors.New("empty name")
	// } 
	// defer res.Body.Close()
	// body, erruh := io.ReadAll(res.Body)
	// if erruh != nil {
	// 	fmt.Println(err)
	// 	return "", errors.New("erruh error")
	// }
	// errah := json.Unmarshal(body, &loc)
	// if errah != nil {
	// 	fmt.Println(errah)
	// }
	// for _, result := range loc.Results {
	// 	fmt.Printf("%s\n", result.Name)
	// }
	// return "Success!!", nil
}

func outputCachedResults(loc *Locations, results []byte) (string, error) {
	error := json.Unmarshal(results, loc)
	if error != nil {
		fmt.Println(error)
		return "", error
	}
	for _, result := range loc.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return "Success!!", nil
}