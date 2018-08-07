package main

import (
	"math"
	"os/exec"
	"time"
)

const (
	delayBetweenFetch = 10000
	blinksPerCycle    = 8
	blinkSpeed        = delayBetweenFetch / blinksPerCycle
	ledCount          = 31
	offset            = ledCount / 2
	script            = "leds.py"
	url               = "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=mont&key=MW9S-E7SL-26DU-VV8V&json=y"
)

func main() {
	var trains []Train
	var status string

	for {
		client := LiveGetWebRequest{}
		response, _ := client.FetchBytes(url)
		newTrains, err := ParseTrains(response)
		if err == nil {
			trains = newTrains
			status = "#000800"
		} else {
			status = "#080000"
		}
		colorArrays := trainColors(trains)
		blink(colorArrays, status)
	}
}

func blink(colorArrays [ledCount][]string, status string) {
	for i := 0; i < blinksPerCycle; i++ {
		lightLeds(colorsAt(colorArrays, status, i))
		time.Sleep(blinkSpeed * time.Millisecond)
	}
}

func lightLeds(colors []string) {
	args := append([]string{"leds.py"}, colors...)

	exec.Command("python", args...).Run()
}

func colorsAt(colorArrays [ledCount][]string, status string, n int) []string {
	var args []string
	for _, colors := range colorArrays {
		if len(colors) > 0 {
			args = append(args, colors[n%len(colors)])
		} else {
			args = append(args, "0")
		}
	}

	for i := 0; i < 21; i++ {
		args = append(args, "0")
	}
	args = append(args, []string{status, "0"}[n%2])

	return args
}

func trainColors(trains []Train) [ledCount][]string {
	var colors [ledCount][]string
	for _, train := range trains {
		if math.Abs(float64(train.Location())) <= offset {
			minutes := train.Location() + offset
			colors[minutes] = append(colors[minutes], train.Hexcolor)
		}
	}
	return colors
}
