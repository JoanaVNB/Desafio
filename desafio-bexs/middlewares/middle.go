package middle

import (
	"desafio/controllers"
	"github.com/gin-gonic/gin"
)

func Autenticar() gin.HandlerFunc {
	return func (c *gin.Context) {
		controllers.Login(c)
		c.Next()
	}
}