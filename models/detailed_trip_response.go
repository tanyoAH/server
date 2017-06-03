package models

type DetailedTripResponse struct {
	Trip
	User       BasicUserResponse       `json:"user"`
	Activities []BasicActivityResponse `json:"activities"`
}
