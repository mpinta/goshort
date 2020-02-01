package handler

type ShortenReq struct {
	FullUrl      string `json:"full_url"`
	MinutesValid int    `json:"minutes_valid"`
}
