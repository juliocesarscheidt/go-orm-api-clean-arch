package dto

// HttpResponseDto - simple HTTP response DTO
type HttpResponseDto struct {
	Data interface{} `json:"data"`
}

type HttpResponseMessageDto struct {
	Message interface{} `json:"message"`
}
