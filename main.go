package main

import (
	"informed-booking-api/controllers"
	"net/http"
)

func main() {
	bh := controllers.NewBookingHandlers()

	http.HandleFunc("/", bh.Get)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
