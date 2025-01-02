package response

import (
	"github.com/gin-gonic/gin"
)
type Response struct{
	Code 	int			`json:"code"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data,omitempty"`
}

func Success(c *gin.Context,data interface{}){
	

}
