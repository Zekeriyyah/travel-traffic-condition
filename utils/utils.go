package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APP struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

type TrafficData struct {
	APP
	Status             string   `json:"status"`
	DestinationAddress []string `json:"destination_addresses"`
	OriginAddress      []string `json:"origin_addresses"`
	Rows               []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`

			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`

			DurationInTraffic struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration_in_traffic"`

			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
}

var app *APP

func (data *TrafficData) GetData(url string) {
	resp, err := http.Get(url)
	if err != nil {
		app.ErrorLog.Fatalf("HTTP request failed: %s\n", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.ErrorLog.Fatalf("Error reading response body: %s\n", err.Error())
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		app.ErrorLog.Printf("Error decoding JSON response: %s\n", err.Error())
		return
	}
}

func (data *TrafficData) HandleErrors() {
	// Checking for status of the response
	if data.Status != "OK" {
		msg := data.CheckStatus()
		app.ErrorLog.Fatal(msg)
	}

	// Checking for status of the response element
	if data.Rows[0].Elements[0].Status != "OK" {
		msg := data.CheckElementStatus()
		app.ErrorLog.Fatal(msg)
	}
}
func (data *TrafficData) PrintTrafficData() {
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
