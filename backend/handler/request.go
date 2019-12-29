package handler

type Request struct {
	FullUrl      string `json:"url"`
	MinutesValid int    `json:"minutes_valid"`
}
