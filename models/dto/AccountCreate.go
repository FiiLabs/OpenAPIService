package dto

type accountCreate struct {
	Account     string `json:"account"`
	Name        string `json:"name"`
	OperationId string `json:"operation_id"`
}

type AccountCreateDTO struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Data       accountCreate `json:"data"`
}
