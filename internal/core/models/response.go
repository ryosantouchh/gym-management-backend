package models

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	ResponseMessage
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}
