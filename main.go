package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	utility "github.com/Zekeriyyah/travel-traffic/utils"
	"github.com/joho/godotenv"
)

const (
	apiURL = "https://maps.googleapis.com/maps/api/distancematrix/json"
)

func main() {

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &utility.APP{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
	// Parse latitude and longitude from command line arguments
	lat1 := flag.String("lat_o", "", "Latitude for the origin")
	lng1 := flag.String("lng_o", "", "Longitude for the origin")

	lat2 := flag.String("lat_d", "", "Latitude for the destination")
	lng2 := flag.String("lng_d", "", "Longitude for the destination")
	flag.Parse()

	if *lat1 == "" || *lng1 == "" || *lat2 == "" || *lng2 == "" {
		app.InfoLog.Println("Please provide latitude and longitude for the origin and destination")
		return
	}

	// Loading Google API key from environment variable
	err := godotenv.Load(".env")
	if err != nil {
		app.ErrorLog.Fatal("Error loading environment_variables...")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		app.InfoLog.Fatal("Please set the GOOGLE_API_KEY environment variable")
	}

	// Building API request URL
	url := fmt.Sprintf("%s?departure_time=now&origins=%s,%s&destinations=%s,%s&key=%s", apiURL, *lat1, *lng1, *lat2, *lng2, apiKey)

	// invoking HTTP GET request to Google Distance Matrix API to populate TrafficData
	data := &utility.TrafficData{}
	data.GetData(url)

	// Handling errors related to response and response elements
	data.HandleErrors()

	// Extracting and printing traffic condition
	data.PrintTrafficData()
}
