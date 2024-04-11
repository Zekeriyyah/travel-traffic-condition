package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	apiURL = "https://maps.googleapis.com/maps/api/distancematrix/json"
)

type TrafficData struct {
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

func main() {

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Parse latitude and longitude from command line arguments
	lat1 := flag.String("lat_o", "", "Latitude for the origin")
	lng1 := flag.String("lng_o", "", "Longitude for the origin")

	lat2 := flag.String("lat_d", "", "Latitude for the destination")
	lng2 := flag.String("lng_d", "", "Longitude for the destination")
	flag.Parse()

	if *lat1 == "" || *lng1 == "" || *lat2 == "" || *lng2 == "" {
		infoLog.Println("Please provide latitude and longitude for the origin and destination")
		return
	}

	// Loading Google API key from environment variable
	err := godotenv.Load(".env")
	if err != nil {
		errorLog.Fatal("Error loading environment_variables...")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		infoLog.Println("Please set the GOOGLE_API_KEY environment variable")
		return
	}

	// Building API request URL
	url := fmt.Sprintf("%s?departure_time=now&origins=%s,%s&destinations=%s,%s&key=%s", apiURL, *lat1, *lng1, *lat2, *lng2, apiKey)

	// invoking HTTP GET request to Google Distance Matrix API
	resp, err := http.Get(url)
	if err != nil {
		errorLog.Printf("HTTP request failed: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return
	}

	var data TrafficData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON response: %s\n", err.Error())
		return
	}

	// Checking for status of the response
	if data.Status != "OK" {
		msg := checkStatus(data)
		errorLog.Fatal(msg)
	}

	// Checking for status of the response element
	if data.Rows[0].Elements[0].Status != "OK" {
		msg := checkElementStatus(data)
		errorLog.Fatal(msg)
	}

	// Extracting and printing traffic condition
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
