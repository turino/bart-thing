package main

import (
	"fmt"
	"testing"
)

const testJSON = `{"root":{"station":[{"name":"Montgomery St.","abbr":"MONT","etd":[{"destination":"SF Airport","abbreviation":"SFIA","limited":"0","estimate":[{"minutes":"5","platform":"1","direction":"South","length":"10","color":"YELLOW","hexcolor":"#ffff33","bikeflag":"1","delay":"0"}]},{"destination":"Warm Springs","abbreviation":"WARM","limited":"0","estimate":[{"minutes":"6","platform":"2","direction":"North","length":"10","color":"GREEN","hexcolor":"#339933","bikeflag":"1","delay":"0"},{"minutes":"6","platform":"2","direction":"North","length":"10","color":"RED","hexcolor":"#339933","bikeflag":"1","delay":"0"}]}]}]}}`

func TestParseTrains(t *testing.T) {
	trains, err := ParseTrains([]byte(testJSON))
	// trains, err := ParseTrains(nil)

	// colorArrays := trainColors(trains)
	// blink(colorArrays)

	fmt.Println("blah: ", trains, err)

	if trains[0].Hexcolor != "#003300" {
		t.Error("message: ", trains[0].Hexcolor)
	}
	if len(trains) != 3 {
		t.Error("message: ", "wrong number of trains")
	}
}

const (
// url      = "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=mont&key=MW9S-E7SL-26DU-VV8V&json=y"
// ledCount = 31
// offset   = ledCount / 2
)

// func TestGetTrains(t *testing.T) {
// 	client := LiveGetWebRequest{}
// 	response, err := client.FetchBytes(url)
// 	var trains []Train
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	} else {
// 		trains, err = ParseTrains(response)
// 		if err != nil {
// 			fmt.Println("error: ", err)
// 		} else {
// 			fmt.Println(trains)
// 		}
// 	}

// 	lightLeds(trains)

// 	if len(trains) != 3 {
// 		t.Error("message: ", len(trains))
// 	}
// }
