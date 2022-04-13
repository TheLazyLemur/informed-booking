package data

type Booking struct {
	BookingID int    `json:"booking_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
}
