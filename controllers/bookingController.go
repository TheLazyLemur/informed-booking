package controllers

import (
	"encoding/json"
	"informed-booking-api/data"
	"net/http"
)

type bookingHandlers struct {
	store map[int]data.Booking
}

func (h *bookingHandlers) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	m := make([]data.Booking, 0, len(h.store))
	for _, val := range h.store {
		m = append(m, val)
	}

	u, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(u))
}

func NewBookingHandlers() *bookingHandlers {
	return &bookingHandlers{
		store: map[int]data.Booking{
			1: {
				BookingID: 1,
				Name:      "John Doe",
			},
			2: {
				BookingID: 2,
				Name:      "Jane Doe",
			},
            3: {
                BookingID: 3,
                Name:      "Dan Rousseasu",
            },
		},
	}
}
