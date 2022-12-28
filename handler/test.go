package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/FiiLabs/OpenAPIService/errors"
	"github.com/FiiLabs/OpenAPIService/models/do"
	"github.com/FiiLabs/OpenAPIService/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context)  {
	oriTextHashBytes := sha256.Sum256([]byte("123"))
	xxx := hex.EncodeToString(oriTextHashBytes[:])
	var apikk do.APIKey
	abc,err := apikk.GetSecret("123")
	if err != nil {
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	data := map[string]interface{}{
		"data": xxx,
		"dddd": abc,
	}
	c.JSONP(http.StatusOK, data)
}
