package handler

import "time"

type Response struct {
	ShortUrl   string    `json:"url"`
	ValidUntil time.Time `json:"valid_until"`
}
