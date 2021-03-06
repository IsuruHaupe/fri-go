package main

import (
	"fmt"
	"net/http"

	"github.com/fri-go/internal/server"
)

func main() {
	fmt.Println("Server listening on port 8080")
	fmt.Println("2 routes availables :")
	fmt.Println("GET /{idAirport}/{sensorType}/filter/{date1}/{date2}")
	fmt.Println("GET /{idAirport}/{date}/results/average")

	router := server.InitializeRouter()
	fmt.Println(http.ListenAndServe(":8080", router))
}
