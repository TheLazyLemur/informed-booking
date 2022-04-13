package controllers

import (
	"encoding/json"
	"informed-booking-api/data"
	"net/http"
)

type bookingHandlers struct {
	store map[int]data.Booking
}

func errHandler(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
}

func (h *bookingHandlers) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
	case "POST":
		h.post(w, r)
	case "PUT":
		h.put(w, r)
	case "DELETE":
		h.delete(w, r)
	}
}

func (h *bookingHandlers) delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
}

func (h *bookingHandlers) put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
    _, err := w.Write([]byte("Booking updated"))
    if err != nil {
        errHandler(w, err)
    }
}

func (h *bookingHandlers) post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
    _, err := w.Write([]byte("Booking created"))
    if err != nil {
        errHandler(w, err)
    }
}

func (h *bookingHandlers) get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	m := make([]data.Booking, 0, len(h.store))
	for _, val := range h.store {
		m = append(m, val)
	}

	u, err := json.Marshal(m)
	if err != nil {
        errHandler(w, err)
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
