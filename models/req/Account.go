package req

type AccountCreateReq struct {
	Name string `json:"name"`
	OperationId string `json:"operation_id"`
}
