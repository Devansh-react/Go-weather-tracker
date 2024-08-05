## Golang Weather Tracker





![Logo](https://img.freepik.com/premium-vector/sun-cloud-with-rain-drops-weather-concept-3d-vector-icon-cartoon-minimal-style_365941-815.jpg?w=740)




## Documentation

This Golang application fetches weather data for a specified city using the OpenWeatherMap API. It provides a simple HTTP server with two endpoints:

/start: Returns a welcome message.
/weather/{city}: Returns weather data for the specified city in JSON format. 

### Prerequisites:

Go programming language installed.
An OpenWeatherMap API key. Create a {.apiconfig.json file } in the project directory with the following structure:
### JSON

```http
{
  "Apikey": "YOUR_API_KEY"
}

```
### TO Run

```http
 go run main.go
```




## Installation

Install dependencies using 
```bash
go mod tidy.
```
    
Install dependencies using 
## Deployment
Access the endpoints using a web browser or HTTP client:
(replace {city} with the desired city name)

```bash
http://localhost:3000/start
http://localhost:3000/weather/{city} 
```


## API response

To run tests, run the following command

```bash
{
  "Name": "City Name",
  "Main": {
    "Kelvin": temperature in Kelvin,
    "Celsius": temperature in Celsius
  },
  "Wind": {
    "Speed": wind speed
  },
  "Weather": [
    {
      "Description": weather description
    }
  ]
}

```


## Features

1. The code assumes a .apiconfig.json file exists in the project directory with the API key.

2. Error handling can be improved for API requests and JSON parsing.
Consider implementing unit tests for the query function.

3. The code doesn't explicitly convert Kelvin to Celsius. This conversion can be added to the weatherData struct or the query function.

4. Implement caching for API responses.
5. Add support for different temperature units (e.g., Fahrenheit).
6. Enhance error handling and logging.Improve API response formatting.Deploy the application to a production environment.

