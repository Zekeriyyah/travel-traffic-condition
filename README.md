# **Travel Traffic Condition**
*A Golang program that fetches live traffic data for a given latitude and longitude
using the Google Maps Platform Distance Matrix API.*

### About
The latitude and longitude are provided as a command-line flags
> -lat_o : latitude of the origin
> -lng_o : longitude of the origin
> -lat_d : latitude of the destination
> -lng_d : longitude of the destination
### Run
Clone the repo and run
`go get github.com/joho/godotenv`
`go build .`
`./<app_name> -lat_o ORIG_LAT -lng_o ORIG_LONG -lat_d DEST_LAT -lng_d DEST_LONG`
Replace <app_name> with the name of the generated file after build and the value of the flags with valid latitude and longitude of the origin and the destination

## Example
### Command:
> './travel-traffic-condition -lat_o 40.6655101 -lng_o -73.89188969999998 -lat_d 40.729029 -lng_d -73.7527626'
> ### Response:
> Traffic condition for location from 'P.O. Box 793, Brooklyn, NY 11207, USA' to '215-12 86th Ave, Jamaica, NY 11427, USA':
> Total Distance: 17.8 km
> Time it will take: 27 mins
> Time it will take considering Traffic: 30 mins


