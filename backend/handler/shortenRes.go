package handler

import "time"

type ShortenRes struct {
	ShortUrl   string    `json:"short_url"`
	CreatedAt  time.Time `json:"created_at"`
	ValidUntil time.Time `json:"valid_until"`
}
