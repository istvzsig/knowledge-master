package types

type HttpResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type HttpResponseSuccess struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
