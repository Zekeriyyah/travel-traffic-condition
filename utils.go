package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var app *APP

func (data *TrafficData) getData(url string) {
	resp, err := http.Get(url)
	if err != nil {
		app.errorLog.Fatalf("HTTP request failed: %s\n", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.errorLog.Fatalf("Error reading response body: %s\n", err.Error())
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		app.errorLog.Printf("Error decoding JSON response: %s\n", err.Error())
		return
	}
}

func (data *TrafficData) handleErrors() {
	// Checking for status of the response
	if data.Status != "OK" {
		msg := data.CheckStatus()
		app.errorLog.Fatal(msg)
	}

	// Checking for status of the response element
	if data.Rows[0].Elements[0].Status != "OK" {
		msg := data.CheckElementStatus()
		app.errorLog.Fatal(msg)
	}
}
func (data *TrafficData) printTrafficData() {
	if len(data.Rows) > 0 && len(data.Rows[0].Elements) > 0 {
		distance := data.Rows[0].Elements[0].Distance.Text
		duration := data.Rows[0].Elements[0].Duration.Text
		durationInTraffic := data.Rows[0].Elements[0].DurationInTraffic.Text

		fmt.Printf("Traffic condition for location from '%s' to '%s':\n", data.OriginAddress[0], data.DestinationAddress[0])
		fmt.Println()
		fmt.Printf("Total Distance: %s\t\n", distance)
		fmt.Printf("Time it will take: %s\t\n", duration)
		fmt.Printf("Time it will take considering Traffic: %s\t\n", durationInTraffic)
	} else {
		fmt.Println("Unable to retrieve traffic data for the location")
	}
}
