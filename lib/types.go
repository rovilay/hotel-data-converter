package utils

import "io"

// HotelDataConverter handles conversion of Hotel data from one format to another
type HotelDataConverter struct {
	File      io.Reader
	CreatedAt int64
}

// Hotel describes the structure of the hotel data to expect
type Hotel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	URI     string `json:"uri"`
}
