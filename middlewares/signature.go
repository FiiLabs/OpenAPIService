package middlewares

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func SignRequest(r *http.Request, apiKey, apiSecret string) *http.Request {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

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
	hexHash := hash(timestamp + apiSecret)
	for k, v := range paramsBody {
		k = "body_" + k
		params[k] = v
	}
	sortParams := params
	if sortParams != nil {
		sortParamsBytes, _ := json.Marshal(sortParams)
		hexHash = hash(string(sortParamsBytes) + timestamp + apiSecret)
	}
	r.Header.Set("X-Api-Key", apiKey)
	r.Header.Set("X-Signature", hexHash)
	r.Header.Set("X-Timestamp", timestamp)
	return r
}
func hash(oriText string) string {
	oriTextHashBytes := sha256.Sum256([]byte(oriText))
	return hex.EncodeToString(oriTextHashBytes[:])
}
