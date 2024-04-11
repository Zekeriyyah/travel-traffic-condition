# **Travel Traffic Condition**
*A Golang program that fetches live traffic data for a given latitude and longitude
using the Google Maps Platform Distance Matrix API.*

### About
The latitude and longitude are provided as a command-line flags
> -lat_o : latitude of the origin   <br>
> -lng_o : longitude of the origin  <br>
> -lat_d : latitude of the destination   <br>
> -lng_d : longitude of the destination  
To get a reasonable traffic condition of a place, the origin and the destination dimension must be well defined. Hence the need for latitude and longitude 
of the origin and destination. Use the above flags to define the dimensions.
### Run
Clone the repo and run   <br>
Create .env file and store your google map Api key with Key=GOOGLE_MAP_KEY and Value=YOUR_API_KEY    <br> <br>
`go get github.com/joho/godotenv`   <br>
`go build .`   <br>
`./<app_name>  -lat_o  ORIG_LAT  -lng_o  ORIG_LONG  -lat_d  DEST_LAT  -lng_d  DEST_LONG`   <br><br>
Replace <app_name> with the name of the generated file after build and the value of the flags with valid latitude and longitude of the origin and the destination  

## Example
### Command:
> './travel-traffic -lat_o 40.6655101 -lng_o -73.89188969999998 -lat_d 40.729029 -lng_d -73.7527626'    <br>
### Response:
> Traffic condition for location from 'P.O. Box 793, Brooklyn, NY 11207, USA' to '215-12 86th Ave, Jamaica, NY 11427, USA':   <br><br>
> Total Distance: 17.8 km    <br>
> Time it will take: 27 mins   <br>
> Time it will take considering Traffic: 30 mins   <br>


