package main

import (
	"fmt"
	"informed-booking-api/controllers"
	"net/http"
)

func main() {
	bh := controllers.NewBookingHandlers()

	http.HandleFunc("/", bh.Handle)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
