package http

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

//checking for valid graphql command
func CheckRequest(c *gin.Context, operationName string) bool {
	if c.Request.Method == "GET" {
		pram := c.Request.URL.RequestURI()
		splitedString := strings.Split(pram, "{")
		if len(splitedString) == 1 {
			return false
		}
		output := strings.Contains(splitedString[1], operationName)
		return output

	} else {
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := string(buf[0:num])

		splitedString := strings.Split(reqBody, "{")
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))
		if len(splitedString) == 1 {
			return false
		}
		return strings.Contains(splitedString[1], operationName)
	}

}
