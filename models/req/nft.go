package req

type NFTReq struct {
	UsrName           string  `json:"usr_name"`
	OpId              string  `json:"operation_id"`
	NFTName           string  `json:"name"`
	NFTId             string  `json:"nft_id"`
	Denom             string  `json:"class_id"`
	Uri               string  `json:"uri"`
	UriHash           string  `json:"uri_hash"`
	Data              string  `json:"data"`
	Recipient		  string  `json:"recipient"`
}

type NFTTrfReq struct {
	UsrName           string  `json:"usr_name"`
	OpId              string  `json:"operation_id"`
	ClsId             string  `json:"class_id"`
	NFTId             string  `json:"nft_id"`
	Recipient		  string  `json:"recipient"`
}

type NFTEditReq struct {
	UsrName           string  `json:"usr_name"`
	OpId              string  `json:"operation_id"`
	NFTName           string  `json:"name"`
	NFTId             string  `json:"nft_id"`
	Denom             string  `json:"class_id"`
	Uri               string  `json:"uri"`
	UriHash           string  `json:"uri_hash"`
	Data              string  `json:"data"`
}