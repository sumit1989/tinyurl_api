package model

type Response struct {
	ResponseCode   string `json:"responseCode"`
	ResponseMssage string `json:"responseMssage"`
}

type Error struct {
	ErrorCode   int `json:"errorCode"`
	ErrorMssage string `json:"errorMssage"`
}
