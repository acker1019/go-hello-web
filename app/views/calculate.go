package views

import (
	"ack/go-hello-web/app/services/calculate"
	"ack/go-hello-web/app/validator"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Calculate(c *gin.Context) {
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		return
	}

	operator := strings.ToLower(data["operator"].(string))
	operands := data["operands"].([]interface{})

	var output interface{}
	switch operator {

	case "add":
		if !validator.IsBinary(operands) {
			c.String(http.StatusUnprocessableEntity, "Err: \"add\" is a binary operation!")
			return
		}
		output = calculate.Add(operands[0], operands[1])

	default:
		c.String(http.StatusNotFound, "Unsuppored operator!")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("%v", output))
}
