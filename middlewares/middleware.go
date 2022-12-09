package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
type OpenApiKey struct {
	ApiKey    string `bson:"api_key" json:"api_key"`
	ApiSecret string `bson:"api_secret" json:"api_secret"`
}
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
const (
	HeaderTimestamp = "X-Timestamp"
	HeaderSignature = "X-Signature"
	HeaderApiKey    = "X-Api-Key"
	NetworkDelay    = 1500

	signStrFmt                = "X-Timestamp: %d\nURI: %s\nBody: %s"
)
const (
	ErrBadRequest  = 40000 // 错误的请求
	ErrTxNotFound  = 40004
	ErrTxNotUnique = 40005

	ErrSystemError = 50000 // 系统异常
)

func FailBadRequest(msg string) BaseResponse {
	return BaseResponse{
		Code:    ErrBadRequest,
		Message: msg,
	}
}
func SignatureVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		{
			//TODO:should delete in produce env
			c.Request = SignRequest(c.Request, "xxx", "123")
			apiKey := c.Request.Header.Get(HeaderApiKey)
			ok, apiSecret := verifyApiKey(apiKey)
			if !ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, FailBadRequest(fmt.Sprintf("Invalid %s", HeaderApiKey)))
				return
			}

			timestamp := c.Request.Header.Get(HeaderTimestamp)
			timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil || !checkTimeliness(timestampInt) {
				c.AbortWithStatusJSON(http.StatusBadRequest, FailBadRequest(fmt.Sprintf("Timeliness error. Please check the request header %s", HeaderTimestamp)))
				return
			}

			signature := c.Request.Header.Get(HeaderSignature)

			if calculateSignature(c.Request,apiSecret, timestampInt) == signature {
				c.Next()
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, FailBadRequest(fmt.Sprintf("Invalid %s", HeaderSignature)))
			}
		}
	}
}

func calculateSignature(r *http.Request,apiSecret string, timestamp int64) string {
	timestamps := strconv.FormatInt(timestamp, 10)
	params := map[string]interface{}{}
	params["path_url"] = r.URL.Path

	for k, v := range r.URL.Query() {
	k = "query_" + k
	params[k] = v[0]
	}
	var bodyBytes []byte
	if r.Body != nil {
	bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	if bodyBytes != nil {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	paramsBody := map[string]interface{}{}
	_ = json.Unmarshal(bodyBytes, &paramsBody)
	hexHash := hash(timestamps + apiSecret)
	for k, v := range paramsBody {
	k = "body_" + k
	params[k] = v
	}
	sortParams := params
	if sortParams != nil {
		sortParamsBytes, _ := json.Marshal(sortParams)
		hexHash = hash(string(sortParamsBytes) + timestamps + apiSecret)
	}
	return hexHash
}

func checkTimeliness(timestamp int64) bool {
	now := time.Now().Unix()
	if now >= timestamp - NetworkDelay && now <= timestamp + NetworkDelay {
		return true
	}
	return false
}

func verifyApiKey(apiKey string) (bool, string) {
	if apiKey == "" {
		return false, ""
	}

	//res, err := apiKeyRepo.FindByApiKey(apiKey)
	//if err != nil {
	//	return false, ""
	//}
	var res OpenApiKey
	res.ApiSecret = "123"
	return true, res.ApiSecret
}