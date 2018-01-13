package bart

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
)

type apiResponse struct {
	Root root `json:"root"`
}

type root struct {
	Stations []station `json:"station"`
}

type station struct {
	Destinations []destination `json:"etd"`
}

type destination struct {
	Abbreviation string  `json:"abbreviation"`
	Trains       []Train `json:"estimate"`
}

// Train is the thing we want
type Train struct {
	Color     string `json:"color"`
	Direction string `json:"direction"`
	Hexcolor  string `json:"hexcolor"`
	Minutes   string `json:"minutes"`
}

// Location is minutes
func (train Train) Location() int {
	i, _ := strconv.Atoi(train.Minutes)
	if train.Direction == "North" {
		i = i * -1
	}
	return i
}

// ParseTrains parses API json response for trains
func ParseTrains(bartJSON []byte) (trains []Train) {
	var response apiResponse

	err := json.Unmarshal(bartJSON, &response)
	if err != nil {
		log.Fatal(err)
	}

	for _, destination := range response.Root.Stations[0].Destinations {
		for _, train := range destination.Trains {
			trains = append(trains, train)
		}
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].Location() < trains[j].Location()
	})

	return trains
}
