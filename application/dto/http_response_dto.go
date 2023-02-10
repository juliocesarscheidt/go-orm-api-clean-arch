package dto

// HttpResponseDto - simple HTTP response DTO for data
type HttpResponseDto struct {
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata"`
}
