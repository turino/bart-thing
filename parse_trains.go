package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strconv"
)

var config configuration

type (
	configuration struct {
		Hexcodes          map[string]string `json:"hexcodes"`
		DestinationColors map[string]string `json:"destinationColors"`
	}

	// Train is the thing we want
	Train struct {
		Color       string `json:"color"`
		Direction   string `json:"direction"`
		Hexcolor    string `json:"hexcolor"`
		Minutes     string `json:"minutes"`
		Destination string
	}

	destination struct {
		Abbreviation string  `json:"abbreviation"`
		Trains       []Train `json:"estimate"`
	}

	station struct {
		Destinations []destination `json:"etd"`
	}

	apiResponse struct {
		Root struct {
			Stations []station `json:"station"`
		} `json:"root"`
	}
)

func init() {
	loadConfig("config.json", &config)
}

func loadConfig(path string, config *configuration) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(config)
	if err != nil {
		log.Fatal(err)
	}
}

// Location is minutes
func (train *Train) Location() int {
	i, _ := strconv.Atoi(train.Minutes)
	if train.Direction == "North" {
		i = i * -1
	}
	return i
}

func (train *Train) updateHexColor() {
	color := config.DestinationColors[train.Destination]
	if len(color) == 0 {
		color = train.Color
	}
	newColor := config.Hexcodes[color]

	if len(newColor) > 0 {
		train.Hexcolor = newColor
	}
}

// ParseTrains parses API json response for trains
func ParseTrains(bartJSON []byte) ([]Train, error) {
	var response apiResponse
	var trains []Train

	err := json.Unmarshal(bartJSON, &response)
	if err != nil {
		return nil, err
	}

	for _, destination := range response.Root.Stations[0].Destinations {
		for _, train := range destination.Trains {
			train.Destination = destination.Abbreviation
			train.updateHexColor()
			trains = append(trains, train)
		}
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].Location() < trains[j].Location()
	})

	return trains, nil
}
