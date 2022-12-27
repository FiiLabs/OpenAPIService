package req

type NFTClsReq struct {
	UsrName           string  `json:"usr_name"`
	OpId              string  `json:"operation_id"`
	ClsName           string  `json:"name"`
	ClsId             string  `json:"class_id"`
	Schema            string  `json:"schema"`
	Symbol            string  `json:"symbol"`
	Description       string  `json:"description"`
	Uri               string  `json:"uri"`
	UriHash           string  `json:"uri_hash"`
	Data              string  `json:"data"`
}

type NFTClsTrfReq struct {
	UsrName           string  `json:"usr_name"`
	OpId              string  `json:"operation_id"`
	ClsId             string  `json:"class_id"`
	Recipient		  string  `json:"recipient"`
}