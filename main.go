package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	apiURL = "https://maps.googleapis.com/maps/api/distancematrix/json"
)

type APP struct {
	errorLog *log.Logger
	infoLog  *log.Logger
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

func main() {

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &APP{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Parse latitude and longitude from command line arguments
	lat1 := flag.String("lat_o", "", "Latitude for the origin")
	lng1 := flag.String("lng_o", "", "Longitude for the origin")

	lat2 := flag.String("lat_d", "", "Latitude for the destination")
	lng2 := flag.String("lng_d", "", "Longitude for the destination")
	flag.Parse()

	if *lat1 == "" || *lng1 == "" || *lat2 == "" || *lng2 == "" {
		app.infoLog.Println("Please provide latitude and longitude for the origin and destination")
		return
	}

	// Loading Google API key from environment variable
	err := godotenv.Load(".env")
	if err != nil {
		app.errorLog.Fatal("Error loading environment_variables...")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		app.infoLog.Fatal("Please set the GOOGLE_API_KEY environment variable")
	}

	// Building API request URL
	url := fmt.Sprintf("%s?departure_time=now&origins=%s,%s&destinations=%s,%s&key=%s", apiURL, *lat1, *lng1, *lat2, *lng2, apiKey)

	// invoking HTTP GET request to Google Distance Matrix API and populate TrafficData
	var data TrafficData
	data.getData(url)

	// Handling errors related to response and response elements
	data.handleErrors()

	// Extracting and printing traffic condition
	data.printTrafficData()
}
