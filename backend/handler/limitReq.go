package handler

type LimitReq struct {
	FullUrl      string `json:"full_url"`
	MinutesValid int    `json:"minutes_valid"`
}
